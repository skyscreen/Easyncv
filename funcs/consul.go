package funcs

import (
	log "github.com/Sirupsen/logrus"
	consulApi "github.com/hashicorp/consul/api"

	"fmt"
	"os"
	"encoding/json"
)

var consulClient *ConsulClient = new(ConsulClient)

type ConsulClient struct {
	client *consulApi.Client
}

type ParamsConsul struct{
	Framework string `json:"framework"`
	Version string `json:"version"`
	Consulurl string `json:"url"`

}


func LoadParamsConfConsul(file string) ParamsConsul {
	var paramsConsul ParamsConsul
	paramsFile, err := os.Open(file)
	defer paramsFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonPoint := json.NewDecoder(paramsFile)
	jsonPoint.Decode(&paramsConsul)
	return paramsConsul
}

func GetConsulClient(consulUrl string) (*ConsulClient, error) {
	/* reuse consul client */
	e := consulClient.connect(consulUrl)
	if e != nil {
		log.Errorf("Error occured %v.", e.Error())
		return nil, e
	}
	return consulClient, nil

}


func (c *ConsulClient) connect(rawUrl string) error {

	logPrefix := "ConsulAPI[api/client/consul.go]: Connect:"

	var alive bool
	if nil != c.client {
		/* validate if conection is successful by getting consul's leader */
		leader, e := c.client.Status().Leader()
		if e != nil || leader == "" {
			var msg string
			if e != nil {
				msg = fmt.Sprintf("%v cannot fetch consul leader, connection is bad. %v", logPrefix, e)
			} else {
				msg = fmt.Sprintf("%v fetched consul leader is empty, connection is bad.", logPrefix)
			}
			log.Warnf(msg)
		} else {
			alive = true
		}
	}
	if c.client == nil || !alive {
		cli, e := c.newConsulClient(rawUrl)
		if e != nil {
			log.Error(e)
			return e
		}
		c.client = cli
	}

	return nil
}

// newConsulClient uses http protocol
func (c *ConsulClient) newConsulClient(url string) (*consulApi.Client, error) {
	/* url: "127.0.0.1:8500",*/

	cli, err := consulApi.NewClient(&consulApi.Config{
		Address: url,
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return cli, nil
}


func (c *ConsulClient) PutKeyValue(key string, val string) (string, error) {
	kv := c.client.KV()
	Pair := &consulApi.KVPair{Key: key, Value: []byte(val)}
	if _, err := kv.Put(Pair, nil); err != nil {
		log.WithFields(log.Fields{
			"key":   key,
			"val":   val,
			"error": err,
		}).Error("Unable to put key/value into consul")
		return key, err
	}
	return "", nil
}

func (c *ConsulClient) DeletePrefix(key string) (bool, error) {
	kv := c.client.KV()
	if _, err := kv.DeleteTree(key, nil); err != nil {
		log.Debugf("Unable to delete key/value pair")
		return false, err
	}
	return true, nil
}