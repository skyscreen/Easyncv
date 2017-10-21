package main

import
(

	log "github.com/Sirupsen/logrus"

	"easyncv/funcs"
	"fmt"

)



func main() {
	fmt.Println("Deploy start")


	//load parameters from json file
	params :=funcs.LoadParamsConf("conf/hclstop.json")
	paramsConsul :=funcs.LoadParamsConfConsul("conf/consul.json")


	cli, e := funcs.GetNomadClient(params.NomadUrl)
	if e != nil {
		log.Error(e)

	}

	cliConsul, e := funcs.GetConsulClient(paramsConsul.Consulurl)
	if e != nil {
		log.Error(e)

	}

	arg := params.Run


	if (arg == "stop") {

		evalID := params.JobId

		cli.KillJob(evalID)

		//consul exectute
		// push framework state to consul
		prefix := fmt.Sprintf("framework /%v/%v/state", paramsConsul.Framework, paramsConsul.Version)

		// delete all keys under deleteTree path
		_, err := cliConsul.DeletePrefix(prefix)
		if err != nil {
			log.Errorf("Consul Worker[worker/consul.go]: ERROR %v", err.Error())

		}

	}




}
