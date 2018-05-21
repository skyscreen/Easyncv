package controllers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/astaxie/beego"

	"testgo/nomad"

	"fmt"
	"easyncv/funcs"
)

type StopController struct {
	beego.Controller
}

type Stopfile struct {
	STOP_FILE string         `form:"stopfile"`
}

func (c *StopController) Get() {

	c.Data["Form"] = &Stopfile{}
	c.Data["Email"] = "sky.zhao@hydsoft.com"
	c.TplName = "stopcontroller/get.tpl"

}

func (c *StopController) Post() {
	//load parameters from json file
	//params :=LoadParamsConf("conf.json")
	//params :=nomad.LoadParamsConf("conf/hclstop.json")
	req := Stopfile{}
	if err := c.ParseForm(&req); err != nil {
		//handle error
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("empty body"))
		return
	}
	params := nomad.LoadParamsConf(req.STOP_FILE)

	cli, e := nomad.GetNomadClient(params.NomadUrl)
	if e != nil {
		log.Error(e)

	}

	//load parameters from json file for consul
	paramsConsul := funcs.LoadParamsConfConsul("conf/consul.json")
	cliConsul, e := funcs.GetConsulClient(paramsConsul.Consulurl)
	if e != nil {
		log.Error(e)

	}

	evalID := params.JobId

	fmt.Print(evalID)

	cli.KillJob(evalID)

	c.Data["jobid"] = evalID

	//consul exectute
	prefix := fmt.Sprintf("framework /%v/%v/state", paramsConsul.Framework, paramsConsul.Version)

	// delete all keys under deleteTree path
	_, err := cliConsul.DeletePrefix(prefix)
	if err != nil {
		log.Errorf("Consul Worker[worker/consul.go]: ERROR %v", err.Error())

	}

}