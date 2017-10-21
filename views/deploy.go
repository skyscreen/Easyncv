package main

import
(

	log "github.com/Sirupsen/logrus"


	"easyncv/func"
	"io/ioutil"
	"fmt"
	"bytes"
)



func main() {
	fmt.Println("Deploy start")


	//load parameters from json file
	params :=nomad.LoadParamsConf("conf/hcl.json")


	cli, e := nomad.GetNomadClient(params.NomadUrl)
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
	}




}
