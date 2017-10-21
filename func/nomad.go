package nomad

import
(

	log "github.com/Sirupsen/logrus"
	nomadApi "github.com/hashicorp/nomad/api"
	"github.com/hashicorp/nomad/jobspec"
	_"github.com/GeertJohan/go.rice"


	structs "easyncv/structs"
	"strings"
	"errors"
	_"text/template"
	"bytes"
	"os"
	"fmt"
	"encoding/json"
)



var ncvConf *structs.NCVConfiguration


type Params struct{
	Run string `json:"run"`
	JobId string `json:"jobid"`
	NomadUrl string `json:"nomadurl"`
	HclFile string `json:"hclfile"`

}

const (
	NOMAD_AGENT_RETRY_TIME = 3
)



// NomadClient contains reference to hashicorp nomad api
type NomadClient struct {
	structs.Orch
	client *nomadApi.Client
}

func LoadParamsConf(file string) Params {
	var params Params
	paramsFile, err := os.Open(file)
	defer paramsFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonPoint := json.NewDecoder(paramsFile)
	jsonPoint.Decode(&params)
	return params
}

// SubmitJob ensures that it can successfully ask nomad to run a job
func (c *NomadClient) SubmitJob(nomadConfBuf interface{}) (string, error) {

	/* parse job */
	job, err := jobspec.Parse(nomadConfBuf.(*bytes.Buffer))
	if err != nil {
		log.Error(err)
		return "", err
	}
	job.Canonicalize()


	// Submit the job
	//TODO extract status code, if 500, fatal, 400, failure

	evalID, _, err := c.client.Jobs().Register(job, nil)
	if err != nil {
		log.Error(err)
		return "", err
	}

	return evalID.EvalID, nil
}



func GetNomadClient(rawUrl string) (*NomadClient, error) {
	logPrefix := "NomadAPI[api/client/nomad.go]: GetNomadClient:"
	log.Debugf("%v start", logPrefix)
	defer log.Debugf("%v done", logPrefix)

	var nomadCli *NomadClient
	for i := 0; i < NOMAD_AGENT_RETRY_TIME; i++ {
		storageSplit := strings.Split(rawUrl, ":")

		url :=storageSplit[0]
		if 3 == len(storageSplit) {
			switch storageSplit[0] {
			case structs.CONSUL:
				url = storageSplit[1]
			}
		}

		log.Infof("Creating nomad client, retry %v ", i+1)
		nomadCli, err := NewNomadClient(url)
		if err != nil {
			log.Errorf("Unable to get Nomad client at %v time, Info: %v", i+1, err.Error())
			if i == NOMAD_AGENT_RETRY_TIME-1 {
				return nil, err
			}
		} else {
			log.Infof("Get a nomad client")
			return nomadCli, nil
		}
	}
	return nomadCli, nil

}

func NewNomadClient(url string) (*NomadClient, error) {
	logPrefix := "NomadAPI[api/client/nomad.go]: NewNomadClient:"
	log.Debugf("%v start", logPrefix)
	defer log.Debugf("%v done", logPrefix)

	/* create nomad client */
	cli := new(NomadClient)

	e := cli.connect(url)
	if e != nil {
		log.Errorf("%v cannot connect to nomad: %v.", logPrefix, e.Error())
		return nil, e
	}

	return cli, nil
}


func NCVConfig() *structs.NCVConfiguration {
	return ncvConf
}

// Connect connect to Nomad service remotely
func (c *NomadClient) connect(url string) error {
	logPrefix := "NomadAPI[api/client/nomad.go]: Connect:"
	log.Debugf("%v start", logPrefix)
	defer log.Debugf("%v done", logPrefix)


	cli, err := nomadApi.NewClient(
		&nomadApi.Config{
			Address:   "http://" + url + ":4646",
			TLSConfig: &nomadApi.TLSConfig{},
		})
	if err != nil {
		return err
	}
	c.client = cli

	/* validate if conection is successful by getting nomad's leader */
	leader, e := c.client.Status().Leader()
	if e != nil {
		log.Errorf("%v cannot fetch nomad leader, connection is bad. %v", logPrefix, e)
		return e
	}
	if leader == "" {
		msg := fmt.Sprintf("%v fetched leaedr is empty, connection is bad.", logPrefix)
		log.Error(msg)
		return errors.New(msg)
	}

	return nil
}


// GetNomadVersion that returns nomad agent version
func (c *NomadClient) GetNomadVersion() (string, error) {

	self, _ := c.client.Agent().Self()
	if nomadVersion, ok := self.Config["Version"].(string); !ok {
		return "", errors.New("No 'Version' field in nomad-->agent-->self-->config, so we can't get nomad version")
	} else {
		return nomadVersion, nil
	}

	return "", nil
}




// KillJob stop a running job.
func (c *NomadClient) KillJob(jobID string) (string, error) {
	evalID, _, err := c.client.Jobs().Deregister(jobID,true, nil)
	if err != nil {
		log.Debug("KillJob Failed evalId: %v", evalID)
		return evalID, err
	}
	return evalID, nil
}