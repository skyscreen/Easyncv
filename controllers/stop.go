package controllers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/astaxie/beego"

	"testgo/nomad"



)

type StopController struct {
	beego.Controller
}


func (c *StopController) Get() {
	//load parameters from json file
	//params :=LoadParamsConf("conf.json")
	params :=nomad.LoadParamsConf("conf/hclstop.json")


	cli, e := nomad.GetNomadClient(params.NomadUrl)
	if e != nil {
		log.Error(e)

	}

	arg := params.Run

	if (arg == "stop") {

		evalID := params.JobId

		cli.KillJob(evalID)

		c.Data["jobid"] = evalID

	}

}