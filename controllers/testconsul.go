package controllers

import (
	//"easyncv/funcs"
	"github.com/astaxie/beego"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"github.com/nanobox-io/golang-scribble"
)

type TestConsulController struct {
	beego.Controller
}

type TestConsulServer struct {
	TEST_SERVER string         `form:"testserver"`
}

type ConsulRec struct {
	Name       string
	Describtion string
}


var isConsulPeers = false
var consulPeers = ""
var db = ""


func (c *TestConsulController) Get() {

	c.Data["Form"] = &TestConsulServer{}
	c.Data["Email"] = "sky.zhao@hydsoft.com"
	c.TplName = "testconsulcontroller/get.tpl"

	db, _ := scribble.New("consulrecdb", nil)

	records, err := db.ReadAll("consulrecord")
	if err != nil {
		fmt.Println("Error", err)
	}



	var ConsulRecs = []ConsulRec{}

	for _, f := range records {
		consulRec := ConsulRec{}
		if err := json.Unmarshal([]byte(f), &consulRec); err != nil {
			fmt.Println("Error", err)
		}
		fmt.Println(consulRec.Name + ":" + consulRec.Describtion + "<br>")

		ConsulRecs = append(ConsulRecs,consulRec)
	}
	c.Data["ConsulRecs"] = ConsulRecs




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
	var f2 string
	if reqerr != nil {
		c.Data["desc"] = " From server " + req.TEST_SERVER + "  can not find consul peer servers"
		db, _ := scribble.New("consulrecdb", nil)
		consulrec_1 := ConsulRec{}
		consulrec_1.Name = req.TEST_SERVER
		consulrec_1.Describtion = c.Data["desc"].(string)
		db.Write("consulrecord", req.TEST_SERVER, consulrec_1)
		return

	}

	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))

		jsonErr := json.Unmarshal(body, &f1)
		if jsonErr != nil {
			db, _ := scribble.New("consulrecdb", nil)
			c.Data["desc"] = "  can not find  consul peer servers"
			consulrec_1 := ConsulRec{}
			consulrec_1.Name = req.TEST_SERVER
			consulrec_1.Describtion = c.Data["desc"].(string)
			db.Write("consulrecord", req.TEST_SERVER, consulrec_1)
			//log.Fatal(jsonErr)
			return
		}
		fmt.Println("==============================")

		for _, v := range f1 {

			fmt.Println(" value1=%v", v)
			consulPeers = v + "," + consulPeers
			isConsulPeers = true

		}

		if isConsulPeers == true {
			url  = "http://" + req.TEST_SERVER + ":8500/v1/status/leader"
			request, _ := http.NewRequest("GET", url, nil)
			request.Header.Set("Connection", "keep-alive")
			response, _ := client.Do(request)
			body, _ := ioutil.ReadAll(response.Body)
			json.Unmarshal(body, &f2)

			c.Data["desc"] = " From the server " + req.TEST_SERVER + " and find consul peer servers are :" + consulPeers + " and consul server is " + f2
		} else {
			c.Data["desc"] = "  can not find  consul peer servers"
		}

		db, _ := scribble.New("consulrecdb", nil)
		consulrec_1 := ConsulRec{}
		consulrec_1.Name = req.TEST_SERVER
		consulrec_1.Describtion = c.Data["desc"].(string)
		db.Write("consulrecord", req.TEST_SERVER, consulrec_1)




	}
}

