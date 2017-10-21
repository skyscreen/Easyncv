package controllers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/astaxie/beego"

	"easyncv/func"
	"fmt"
	"bytes"
)

type StartController struct {
	beego.Controller
}

type Hcl struct {
	JOBID    string         `form:"jobid"`
	REGION  string `form:"region"`
	DATACENTERS  string `form:"datacenters"`
        STYPE  string `form:"type"`
	PRIORITY  int  `form:"priority"`
	GROUP  string `form:"group"`
	COUNT  int `form:"count"`
	EPHEMERAL_DISK_SIZE  int `form:"ephemeral_disk_size"`
	RESTART_ATTEMPTS  int `form:"restart_attempts"`
	RESTART_DELAY  string `form:"restart_delay"`
	RESTART_INTERVAL  string `form:"restart_interval"`
	RESTART_MODE  string `form:"restart_mode"`
	CONSTRAINT_DISTINCT_HOSTS  string `form:"constraint_distinct_hosts"`
	TASK  string `form:"task"`
	TASK_DRIVER  string `form:"task_driver"`
	TASK_CONFIG_IMAGE  string `form:"task_config_image"`
	TASK_CONFIG_LABELS_COBALT_ID  string `form:"task_config_labels_cobalt_id"`
	TASK_CONFIG_LABELS_COBALT_SERVICE_NAME  string `form:"task_config_labels_cobalt_service_name"`
	TASK_CONFIG_LABELS_COBALT_TASK  string `form:"task_config_labels_cobalt_task"`
	TASK_CONFIG_LABELS_COBALT_PODID  string `form:"task_config_labels_cobalt_podid"`
	TASK_CONFIG_NETWORK_MODE  string `form:"task_config_network_mode"`
	TASK_CONFIG_PORT_MAP_HTTP_14840  int `form:"task_config_port_map_http_14840"`
	TASK_CONFIG_PORT_MAP_HTTP_8080  int `form:"task_config_port_map_http_8080"`
	TASK_CONFIG_AUTH_USERNAME  string `form:"task_config_auth_username"`
	TASK_CONFIG_AUTH_PASSWORD  string `form:"task_config_auth_password"`
	TASK_CONFIG_PRIVILEGED  string `form:"task_config_privileged"`
	TASK_CONFIG_USERNS_MODE  string `form:"task_config_userns_mode"`
	TASK_ENV_NOMAD_IP string `form:"task_env_CONSUL_IP"`
	TASK_ENV_CONSUL_IP  string `form:"region"`
	TASK_ENV_A1  string `form:"task_env_A1"`
	TASK_ENV_AN_APP  string `form:"task_env_AN_APP"`
	TASK_ENV_AN_BUILD  string `form:"task_env_AN_BUILD"`
	TASK_ENV_AN_DOMAIN  string `form:"task_env_AN_DOMAIN"`
	TASK_ENV_AN_INSTANCEID  string `form:"task_env_AN_INSTANCEID"`
	TASK_ENV_AN_PORT  string `form:"task_env_AN_PORT"`
	TASK_ENV_COBALT_CONTAINER_TYPE  string `form:"task_env_COBALT_CONTAINER_TYPE"`
	TASK_ENV_COBALT_DB_PERSIST  string `form:"task_env_COBALT_DB_PERSIST"`
	TASK_ENV_COBALT_ID  string `form:"task_env_COBALT_ID"`
	TASK_ENV_COBALT_LOG_PERSIST  string `form:"task_env_COBALT_LOG_PERSIST"`
	TASK_ENV_COBALT_MODE string `form:"task_env_COBALT_MODE"`
	TASK_ENV_COBALT_MODE_INTERFACE  string `form:"task_env_COBALT_MODE_INTERFACE"`
	TASK_ENV_COBALT_PODID  string `form:"task_env_COBALT_PODID"`
	TASK_ENV_COBALT_SERVICE_NAME  string `form:"task_env_COBALT_SERVICE_NAME"`
	TASK_ENV_COBALT_SERVICE_VERSION  string `form:"task_env_COBALT_SERVICE_VERSION"`
	TASK_ENV_COBALT_WS  string `form:"task_env_COBALT_WS"`
	TASK_ENV_ENTRY  string `form:"ENTRY"`
	TASK_ENV_ORACLE_COBALT_ID  string `form:"task_env_ORACLE_COBALT_ID"`
	TASK_ENV_ORACLE_SERVICE_NAME  string `form:"task_env_ORACLE_SERVICE_NAME"`
	TASK_LOGS_MAX_FILES  string `form:"task_logs_max_files"`
	TASK_LOGS_MAX_FILES_SIZE  string `form:"task_logs_max_files_size"`
	TASK_RESOURCES_CPU  string `form:"task_resources_cpu"`
	TASK_RESOURCES_MEMORY  string `form:"task_resources_memory"`
	TASK_RESOURCES_NETWORK_MBITS  string `form:"task_resources_network_mbits"`
	TASK_RESOURCES_IOPS  int `form:"task_resources_iops"`
	META_IS_ENTRY  string `form:"meta_is_entry"`
}

func (c *StartController) Get() {
	c.Data["Form"] = &Hcl{}
	c.Data["Email"] = "sky.zhao@hydsoft.com"
	c.TplName = "index.tpl"
}


func (c *StartController) Post() {


	req := Hcl{}

	if err := c.ParseForm(&req); err != nil {
		//handle error
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("empty body"))
		return
	}
        c.Data["JOBID"] = req.JOBID;
	c.Data["REGION"] = req.REGION;
	c.Data["DATACENTERS"] = req.DATACENTERS;
	c.Data["STYPE"] = req.STYPE;
	c.Data["PRIORITY"] = req.PRIORITY;
	c.Data["GROUP"] = req.GROUP;
	c.Data["COUNT"] = req.COUNT;
	c.Data["EPHEMERAL_DISK_SIZE"] = req.EPHEMERAL_DISK_SIZE;
	c.Data["RESTART_ATTEMPTS"] = req.RESTART_ATTEMPTS;
	c.Data["RESTART_DELAY"] = req.RESTART_DELAY;
	c.Data["RESTART_INTERVAL"] = req.RESTART_INTERVAL;
	c.Data["RESTART_MODE"] = req.RESTART_MODE;
	c.Data["CONSTRAINT_DISTINCT_HOSTS"] = req.CONSTRAINT_DISTINCT_HOSTS;
	c.Data["TASK"] = req.TASK;
	c.Data["TASK_DRIVER"] = req.TASK_DRIVER;
	c.Data["TASK_CONFIG_IMAGE"] = req.TASK_CONFIG_IMAGE;
	c.Data["TASK_CONFIG_LABELS_COBALT_ID"] = req.TASK_CONFIG_LABELS_COBALT_ID;
	c.Data["TASK_CONFIG_LABELS_COBALT_SERVICE_NAME"] = req.TASK_CONFIG_LABELS_COBALT_SERVICE_NAME;
	c.Data["TASK_CONFIG_LABELS_COBALT_TASK"] = req.TASK_CONFIG_LABELS_COBALT_TASK;
	c.Data["TASK_CONFIG_LABELS_COBALT_PODID"] = req.TASK_CONFIG_LABELS_COBALT_PODID;
	c.Data["TASK_CONFIG_NETWORK_MODE"] = req.TASK_CONFIG_NETWORK_MODE;
	c.Data["TASK_CONFIG_PORT_MAP_HTTP_14840"] = req.TASK_CONFIG_PORT_MAP_HTTP_14840;
	c.Data["TASK_CONFIG_PORT_MAP_HTTP_8080"] = req.TASK_CONFIG_PORT_MAP_HTTP_8080;
	c.Data["TASK_CONFIG_AUTH_USERNAME"] = req.TASK_CONFIG_AUTH_USERNAME;
	c.Data["TASK_CONFIG_AUTH_PASSWORD"] = req.TASK_CONFIG_AUTH_PASSWORD;
	c.Data["TASK_CONFIG_PRIVILEGED"] = req.TASK_CONFIG_PRIVILEGED;
	c.Data["TASK_CONFIG_USERNS_MODE"] = req.TASK_CONFIG_USERNS_MODE;
	c.Data["TASK_ENV_NOMAD_IP"] = req.TASK_ENV_NOMAD_IP;
	c.Data["TASK_ENV_CONSUL_IP"] = req.TASK_ENV_CONSUL_IP;
	c.Data["TASK_ENV_A1"] = req.TASK_ENV_A1;
	c.Data["TASK_ENV_AN_APP"] = req.TASK_ENV_AN_APP;
	c.Data["TASK_ENV_AN_BUILD"] = req.TASK_ENV_AN_BUILD;
	c.Data["TASK_ENV_AN_DOMAIN"] = req.TASK_ENV_AN_DOMAIN;
	c.Data["TASK_ENV_AN_INSTANCEID"] = req.TASK_ENV_AN_INSTANCEID;
	c.Data["TASK_ENV_AN_PORT"] = req.TASK_ENV_AN_PORT;
	c.Data["TASK_ENV_COBALT_CONTAINER_TYPE"] = req.TASK_ENV_COBALT_CONTAINER_TYPE;
	c.Data["TASK_ENV_COBALT_DB_PERSIST"] = req.TASK_ENV_COBALT_DB_PERSIST;
	c.Data["TASK_ENV_COBALT_ID"] = req.TASK_ENV_COBALT_ID;
	c.Data["TASK_ENV_COBALT_LOG_PERSIST"] = req.TASK_ENV_COBALT_LOG_PERSIST;
	c.Data["TASK_ENV_COBALT_MODE"] = req.TASK_ENV_COBALT_MODE;
	c.Data["TASK_ENV_COBALT_MODE_INTERFACE"] = req.TASK_ENV_COBALT_MODE_INTERFACE;
	c.Data["TASK_ENV_COBALT_PODID"] = req.TASK_ENV_COBALT_PODID;
	c.Data["TASK_ENV_COBALT_SERVICE_NAME"] = req.TASK_ENV_COBALT_SERVICE_NAME;
	c.Data["TASK_ENV_COBALT_SERVICE_VERSION"] = req.TASK_ENV_COBALT_SERVICE_VERSION;
	c.Data["TASK_ENV_COBALT_WS"] = req.TASK_ENV_COBALT_WS;
	c.Data["TASK_ENV_ENTRY"] = req.TASK_ENV_ENTRY;
	c.Data["TASK_ENV_ORACLE_COBALT_ID"] = req.TASK_ENV_ORACLE_COBALT_ID;
	c.Data["TASK_ENV_ORACLE_SERVICE_NAME"] = req.TASK_ENV_ORACLE_SERVICE_NAME;
	c.Data["TASK_LOGS_MAX_FILES"] = req.TASK_LOGS_MAX_FILES;
	c.Data["TASK_LOGS_MAX_FILES_SIZE"] = req.TASK_LOGS_MAX_FILES_SIZE;
	c.Data["TASK_RESOURCES_CPU"] = req.TASK_RESOURCES_CPU;
	c.Data["TASK_RESOURCES_MEMORY"] = req.TASK_RESOURCES_MEMORY;
	c.Data["TASK_RESOURCES_NETWORK_MBITS"] = req.TASK_RESOURCES_NETWORK_MBITS;
	c.Data["TASK_RESOURCES_IOPS"] = req.TASK_RESOURCES_IOPS;
	c.Data["META_IS_ENTRY"] = req.META_IS_ENTRY;


	content, err := c.RenderString()

	if err != nil {
		//handle error
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("empty body"))
		return
	}

	fmt.Println(content)

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


		evalID, e := cli.SubmitJob(bytes.NewBufferString(content))

		if e != nil {
			log.Error(e)

		}
		log.Info(evalID)
	}




}