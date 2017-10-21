
job an-service-2017 {
    # Job should run in the US region
    region = "global"
    # Spread tasks between us-west-1 and us-east-1
    datacenters = ["dc1",]
    # run this job globally
    type = "service"
	priority = 50
    group "ubuntu0" {
        # We want x web servers
        count = 1
		ephemeral_disk {
			size = 200
		}
		restart {
    		attempts = 2
    		delay = "15s"
    		interval = "1m"
    		mode = "delay"
		}
		constraint {
			distinct_hosts = "true"
		}
        # Create a web front end using a docker image
        ## task name
        task "ubuntu0" {
            driver = "docker"
            config {
				image = "docker-dev-registry.mo.sap.corp/an/ubuntu:1-0-fe57305-514"
				#labels to be published as Datadog tags
				labels {
					cobalt_id = "an-service-1-0-fe57305-514_1502690250"
					cobalt_service_name = "an-service"
					cobalt_task = "ubuntu0"
					cobalt_podid = "1-0-fe57305-514"
				}
				# contiv-net
				network_mode = "contiv-pod-net"
				## for docker to map ports
				port_map {
							http_14840 = "14840"
							http_8080 = "8080"
				}
				# for registry login
				auth {
					username = "deploy"
					password = "deploy"
				}
				privileged = true
				userns_mode = "host"
            }
            env {
				# for containerbuddy to talk to consul
				NOMAD_IP = "10.173.76.57"
				CONSUL_IP = "global"
				A1 = "A2"
				AN_APP = "Authenticator"
				AN_BUILD = "AN.2017.07.mTrunk-1345"
				AN_DOMAIN = "User"
				AN_INSTANCEID = "101020043"
				AN_PORT = "14840"
				COBALT_CONTAINER_TYPE = "JMX"
				COBALT_DB_PERSIST = "false"
				COBALT_ID = "an-service-1-0-fe57305-514_1502690250"
				COBALT_LOG_PERSIST = "true"
				COBALT_MODE = "dev"
				COBALT_MODE_INTERFACE = "eth1"
				COBALT_PODID = "1-0-fe57305-514"
				COBALT_SERVICE_NAME = "an-service"
				COBALT_SERVICE_VERSION = "1.0"
				COBALT_WS = "cobalt/an-service/an-service-1-0-fe57305-514_1502690250/workspace/"
				ENTRY = "false"
				ORACLE_COBALT_ID = "e9fd7641-e2fe-4f79-5675-2ded07fee448-1_05012016"
				ORACLE_SERVICE_NAME = "Oracle-Linux-Server"
            }
			logs {
				max_files = 10
    			max_file_size = 10
			}
            resources {
				cpu = 500
				memory = 3072
				network {
					mbits = 10
					port "http_14840" {
					}
					port "http_8080" {
					}
					port "stats" {
					}
					port "cobalt"{
					}
				}
				iops = 0
			}
        }
		meta {
			is.entry = "false"
		}
	}
}