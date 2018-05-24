package controllers

import (
	//"easyncv/funcs"
	"github.com/astaxie/beego"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)




type TestNomadController struct {
	beego.Controller
}

type TestServer struct {
	TEST_SERVER string         `form:"testserver"`
}

type services struct {
	ID string `json:"ID"`
	Service string `json:"Service"`
	Address string `json:"Address"`
	Port string `json:"Port"`
	EnableTagOverride string `json:"EnableTagOverride"`
	CreateIndex string `json:"CreateIndex"`
	ModifyIndex string `json:"ModifyIndex"`

}


type servicesRes struct {
	name string
	jstring services

}

var isNomadServer = false
var NomadServerName = ""

func (c *TestNomadController) Get() {

	c.Data["Form"] = &TestServer{}
	c.Data["Email"] = "sky.zhao@hydsoft.com"
	c.TplName = "testnomadcontroller/get.tpl"






}

func (c *TestNomadController) Post() {
	req := TestServer{}
	if err := c.ParseForm(&req); err != nil {
		//handle error
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("empty body"))
		return
	}

	url := "http://" +req.TEST_SERVER + ":8500/v1/agent/services"

	fmt.Println(req.TEST_SERVER)
	log.Println(req.TEST_SERVER)

	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Connection", "keep-alive")
	response, reqerr := client.Do(request)
	var f interface{}
	if reqerr != nil {
		c.Data["desc"] = " From server " + req.TEST_SERVER + "  can not find nomad server"
		return

	}





	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))

		//jsonErr := json.Unmarshal(body, &servicesr)
		jsonErr := json.Unmarshal(body, &f)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		fmt.Println("==============================")
		m := f.(map[string]interface{})
		for k, v := range m {
			//fmt.Println("key=%v, value=%v", k, v)


			if (k == "nomad") {
				for k1, v1 := range v.(map[string]interface{}) {
					fmt.Println("key1=%v, value1=%v", k1, v1)
					if (k1 == "Address") {
						//fmt.Println("vvvvvvvvvvvv %v", v1)
						isNomadServer = true
						NomadServerName = v1.(string)
						break
					}
				}

			}

		}

		if isNomadServer == true {
			c.Data["desc"] = " From the server " + req.TEST_SERVER + " and find nomad server is " + NomadServerName
		} else {
			c.Data["desc"] = "  can not find nomad server"
		}

	}
}
