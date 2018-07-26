package controllers

import (

	"github.com/astaxie/beego"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"github.com/nanobox-io/golang-scribble"
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

type NomadRec struct {
	Name       string
	Describtion string
}

var isNomadServer = false
var NomadServerName = ""

func (c *TestNomadController) Get() {

	c.Data["Form"] = &TestServer{}
	c.Data["Email"] = "sky.zhao@hydsoft.com"
	c.TplName = "testnomadcontroller/get.tpl"

	db, _ := scribble.New("nomadrecdb", nil)

	records, err := db.ReadAll("nomadrecord")
	if err != nil {
		fmt.Println("Error", err)
	}

	var NomadRecs = []NomadRec{}

	for _, f := range records {
		nomadRec := NomadRec{}
		if err := json.Unmarshal([]byte(f), &nomadRec); err != nil {
			fmt.Println("Error", err)
		}
		fmt.Println(nomadRec.Name + ":" + nomadRec.Describtion + "<br>")

		NomadRecs = append(NomadRecs,nomadRec)
	}
	c.Data["NomadRecs"] = NomadRecs


}

func (c *TestNomadController) Post() {
	req := TestServer{}
	if err := c.ParseForm(&req); err != nil {
		//handle error
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("empty body"))
		c.Data["desc"] = "error"
		return
	}

	//url := "http://" +req.TEST_SERVER + ":4646/v1/status/peers"
	url := "http://" +req.TEST_SERVER + ":4646/v1/status/leader"

	fmt.Println(req.TEST_SERVER)
	log.Println(req.TEST_SERVER)

	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Connection", "keep-alive")
	response, reqerr := client.Do(request)

	var f1 string

	if reqerr != nil {
		c.Data["desc"] = " From server " + req.TEST_SERVER + "  can not find nomad server"
		db, _ := scribble.New("nomadrecdb", nil)
		nomadrec_1 := NomadRec{}
		nomadrec_1.Name = req.TEST_SERVER
		nomadrec_1.Describtion = "  can not find nomad server"
		db.Write("nomadrecord", req.TEST_SERVER, nomadrec_1)
		fmt.Println("error url:" + url)
		log.Println("error url:" + url)
		return

	}





	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)


		jsonErr := json.Unmarshal(body, &f1)
		if jsonErr != nil {
			db, _ := scribble.New("nomadrecdb", nil)
			nomadrec_1 := NomadRec{}
			nomadrec_1.Name = req.TEST_SERVER
			nomadrec_1.Describtion = "  can not find nomad server"
			db.Write("nomadrecord", req.TEST_SERVER, nomadrec_1)
			log.Fatal(jsonErr)
			return
		}
		fmt.Println("==============================" )




		fmt.Println(" value1=%v", f1)
		NomadServerName = f1
		isNomadServer = true







		if isNomadServer == true {
			c.Data["desc"] = " From the server " + req.TEST_SERVER + " and find nomad server is " + NomadServerName
		} else {
			c.Data["desc"] = "  can not find nomad server"
		}



		db, _ := scribble.New("nomadrecdb", nil)
		nomadrec_1 := NomadRec{}
		nomadrec_1.Name = req.TEST_SERVER
		nomadrec_1.Describtion = c.Data["desc"].(string)
		db.Write("nomadrecord", req.TEST_SERVER, nomadrec_1)




	}
}
