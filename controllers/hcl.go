package controllers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/astaxie/beego"
	//"testgo/nomad"
	"fmt"
	"easyncv/funcs"
	"bytes"
	"io/ioutil"
)

type HclController struct {
	beego.Controller
}

type HclFile struct {
	HCL_FILEPATH string         `form:"hclfilename"`
}

func (c *HclController) Get() {

	c.Data["Form"] = &HclFile{}
	c.Data["Email"] = "sky.zhao@hydsoft.com"
	c.TplName = "hclcontroller/file.tpl"

}

func (c *HclController) Post() {

	req := HclFile{}

	if err := c.ParseForm(&req); err != nil {
		//handle error
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("empty body"))
		return
	}
	c.Data["HCL_FILEPATH"] = req.HCL_FILEPATH;


	content, err := c.RenderString()

	if err != nil {
		//handle error
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("empty body"))
		return
	}

	fmt.Println(content)

	//load parameters from json file for nomad
	params := funcs.LoadParamsConf(req.HCL_FILEPATH)

	cli, e := funcs.GetNomadClient(params.NomadUrl)
	if e != nil {
		log.Error(e)

	}


	//load parameters from json file for consul
	paramsConsul := funcs.LoadParamsConfConsul("/Users/c5257805/go-workspace/src/easyncv/conf/consul.json")

	cliConsul, e := funcs.GetConsulClient(paramsConsul.Consulurl)
	if e != nil {
		log.Error(e)

	}

	nomadVersion, e := cli.GetNomadVersion()
	if e != nil {
		log.Error(e)

	}
	log.Printf("Nomad Agent Version:%v", nomadVersion)

	dat, err := ioutil.ReadFile(params.HclFile)
	if err != nil {
		log.Error(err)

	}

	evalID, e := cli.SubmitJob(bytes.NewBufferString(string(dat)))

	if e != nil {
		log.Error(e)

	}
	log.Info(evalID)

	//consul exectute
	// push framework state to consul
	prefix := fmt.Sprintf("framework /%v/%v/state", paramsConsul.Framework, paramsConsul.Version)
	_, err = cliConsul.PutKeyValue(prefix, "Running")
	if err != nil {
		log.Errorf("Consul Worker[worker/consul.go]: ERROR %v", err.Error())

	}

}