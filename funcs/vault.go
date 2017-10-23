package funcs

import (
	log "github.com/Sirupsen/logrus"
	vaultApi "github.com/hashicorp/vault/api"


	"fmt"
	"os"
	"encoding/json"
)


var vaultClient *VaultClient = new(VaultClient)

type VaultClient struct {
	client *vaultApi.Client
	InitRequest vaultApi.InitRequest
	InitResponse vaultApi.InitResponse
	Sys vaultApi.Sys


}


type ParamsVault struct{
	Vaulturl string `json:"url"`

}

//process structs



func LoadParamsConfVault(file string) ParamsVault {
	var paramsVault ParamsVault
	paramsFile, err := os.Open(file)
	defer paramsFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonPoint := json.NewDecoder(paramsFile)
	jsonPoint.Decode(&paramsVault)
	return paramsVault
}

func GetVaultClient(vaultUrl string) (*VaultClient, error) {
	/* reuse vault client */
	e := vaultClient.connect(vaultUrl)
	if e != nil {
		log.Errorf("Error occured %v.", e.Error())
		return nil, e
	}
	return vaultClient, nil

}

func (c *VaultClient) connect(rawUrl string) error {

	logPrefix := "VaultAPI[api/client.go]: Connect:"

	var alive bool
	if nil != c.client {
		var msg string
		msg = fmt.Sprintf("%v fetched vault is not empty, connection still.", logPrefix)
		log.Warnf(msg)
		alive = true
		return nil
	}
	if c.client == nil || !alive {
		cli, e := c.newVaultClient(rawUrl)
		if e != nil {
			log.Error(e)
			return e
		}
		c.client = cli
	}

	return nil
}


// newVaultClient uses http protocol
func (c *VaultClient) newVaultClient(url string) (*vaultApi.Client, error) {


	cli, err := vaultApi.NewClient(&vaultApi.Config{
		Address: url,
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return cli, nil
}


func (c *VaultClient) Help(path string) (string, error) {
      help, e:=c.client.Help(path)
	if e != nil {
		log.Error(e)
		return "", e
	}
	return help.Help, nil

}


//below is vault process functions
//https://www.vaultproject.io/intro/getting-started/apis.html

func (c *VaultClient) initSys(opts *vaultApi.InitRequest)(*vaultApi.InitResponse, error){
	var inResponse *vaultApi.InitResponse
	inResponse, e :=c.Sys.Init(opts)
	if e != nil {
		log.Error(e)
		return nil, e
	}
        return inResponse, e

}

func (c *VaultClient) Unseal(shard string) (*vaultApi.SealStatusResponse, error){
	var sealStatusRes *vaultApi.SealStatusResponse
	sealStatusRes, e :=c.Sys.Unseal(shard)
	if e != nil {
		log.Error(e)
		return nil, e
	}
	return sealStatusRes, e
}

func (c *VaultClient) EnableAuthWithOptions(path string, options *vaultApi.EnableAuthOptions) error {

	e :=c.Sys.EnableAuthWithOptions(path, options)
	if e != nil {
		log.Error(e)
		return  e
	}
	return nil
}