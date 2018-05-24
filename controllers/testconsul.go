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

type TestConsulController struct {
	beego.Controller
}

type TestConsulServer struct {
	TEST_SERVER string         `form:"testserver"`
}

var isConsulPeers = false
var consulPeers = ""

func (c *TestConsulController) Get() {

	c.Data["Form"] = &TestConsulServer{}
	c.Data["Email"] = "sky.zhao@hydsoft.com"
	c.TplName = "testnomadcontroller/get.tpl"

}

func (c *TestConsulController) Post() {
	req := TestConsulServer{}
	if err := c.ParseForm(&req); err != nil {
		//handle error
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("empty body"))
		return
	}

	url := "http://" + req.TEST_SERVER + ":8500/v1/status/peers"

	fmt.Println(req.TEST_SERVER)
	log.Println(req.TEST_SERVER)

	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Connection", "keep-alive")
	response, reqerr := client.Do(request)

	var f1 []string
	if reqerr != nil {
		c.Data["desc"] = " From server " + req.TEST_SERVER + "  can not find consul peer servers"
		return

	}

	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))

		jsonErr := json.Unmarshal(body, &f1)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		fmt.Println("==============================")

		for _, v := range f1 {

			fmt.Println(" value1=%v", v)
			consulPeers = v + "," + consulPeers
			isConsulPeers = true

		}

		if isConsulPeers == true {
			c.Data["desc"] = " From the server " + req.TEST_SERVER + " and find consul peer servers are :" + consulPeers
		} else {
			c.Data["desc"] = "  can not find  consul peer servers"
		}

	}
}
