package funcs

import (
	log "github.com/Sirupsen/logrus"
	vaultApi "github.com/hashicorp/vault/api"
	"os"
	"fmt"
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
type role_security_id struct{
	role_id string
	security_id string

}


type secret_data struct{
	key string
	value string

}


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


func (c *VaultClient) SetPolicies(name, token , polices string) error {
	body := map[string]string{
		"polices": polices,
	}


	r := c.client.NewRequest("POST", fmt.Sprintf("/v1/auth/approle/role/%s", name))
	if err := r.SetJSONBody(body); err != nil {
		log.Error(err)
		return err
	}

	r.Headers.Set("Content-Type", "application/json")
	r.Headers.Set("X-Vault-Token", token)

	resp, err := c.client.RawRequest(r)
	if err != nil {
		log.Error(err)
		return err
	}
	defer resp.Body.Close()

	return nil
}


func (c *VaultClient) getCredentials(name, token string) role_security_id {



	r := c.client.NewRequest("GET", fmt.Sprintf("/v1/auth/approle/role/%s/role-id", name))

	r.Headers.Set("Content-Type", "application/json")
	r.Headers.Set("X-Vault-Token", token)


	resp, err := c.client.RawRequest(r)
	if err != nil {
		log.Error(err)
	}
	defer resp.Body.Close()

	var result role_security_id
	err = resp.DecodeJSON(&result)
	if err != nil {
		log.Error(err)
	}

	return result
}

func (c *VaultClient) createSecretId(name, token string) string {



	r := c.client.NewRequest("POST", fmt.Sprintf("/v1/auth/approle/role/%s/secret-id", name))

	r.Headers.Set("Content-Type", "application/json")
	r.Headers.Set("X-Vault-Token", token)

	resp, err := c.client.RawRequest(r)
	if err != nil {
		log.Error(err)
	}
	defer resp.Body.Close()

	var result map[string]map[string]string

	err = resp.DecodeJSON(&result)
	if err != nil {
		log.Error(err)
	}

	return result["data"]["secret_id"]
}


func (c *VaultClient) fetchNewToken(rsi *role_security_id) string {
	body := map[string]string{
		"role_id": rsi.role_id,
		"secret_id": rsi.security_id,
	}


	r := c.client.NewRequest("POST", fmt.Sprintf("/v1/auth/approle/login"))
	if err := r.SetJSONBody(body); err != nil {
		log.Error(err)

	}

	r.Headers.Set("Content-Type", "application/json")


	resp, err := c.client.RawRequest(r)
	if err != nil {
		log.Error(err)

	}
	defer resp.Body.Close()

	var result map[string]map[string]string

	err = resp.DecodeJSON(&result)
	if err != nil {
		log.Error(err)
	}

	return result["auth"]["client_token"]
}


func (c *VaultClient) createSecret(name, token string, sd secret_data) error {
	body := map[string]string{
		sd.key: sd.value,
	}


	r := c.client.NewRequest("POST", fmt.Sprintf("/v1/secret/%s", name))
	if err := r.SetJSONBody(body); err != nil {
		log.Error(err)
		return err
	}

	r.Headers.Set("Content-Type", "application/json")
	r.Headers.Set("X-Vault-Token", token)

	resp, err := c.client.RawRequest(r)
	if err != nil {
		log.Error(err)
		return err
	}
	defer resp.Body.Close()

	return nil
}


func (c *VaultClient) readData(name, token string) secret_data {



	r := c.client.NewRequest("GET", fmt.Sprintf("/v1/secret/%s", name))

	r.Headers.Set("Content-Type", "application/json")
	r.Headers.Set("X-Vault-Token", token)


	resp, err := c.client.RawRequest(r)
	if err != nil {
		log.Error(err)
	}
	defer resp.Body.Close()

	var result secret_data

	err = resp.DecodeJSON(&result)
	if err != nil {
		log.Error(err)
	}

	return result
}
