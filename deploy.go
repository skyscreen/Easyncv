package main

import
(

	log "github.com/Sirupsen/logrus"


	"easyncv/funcs"
	"io/ioutil"
	"fmt"
	"bytes"
)



func main() {
	fmt.Println("Deploy start")


	//load parameters from json file
	params :=funcs.LoadParamsConf("conf/hcl.json")
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

	if arg=="start" {



		/* get nomad agent version*/
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




}
