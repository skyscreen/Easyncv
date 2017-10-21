package main

import
(

	log "github.com/Sirupsen/logrus"

	"easyncv/func"
	"fmt"

)



func main() {
	fmt.Println("Deploy start")


	//load parameters from json file
	params :=nomad.LoadParamsConf("conf/hclstop.json")


	cli, e := nomad.GetNomadClient(params.NomadUrl)
	if e != nil {
		log.Error(e)

	}

	arg := params.Run


	if (arg == "stop") {

		evalID := params.JobId

		cli.KillJob(evalID)



	}




}
