job {{.JOBID}} {
    # Job should run in the US region
    region = "{{.REGION}}"
    # Spread tasks between us-west-1 and us-east-1
    datacenters = ["{{.DATACENTERS}}",]
    # run this job globally
    type = "{{.STYPE}}"
	priority = {{.PRIORITY}}
    group "{{.GROUP}}" {
        # We want x web servers
        count = {{.COUNT}}
		ephemeral_disk {
			size = {{.EPHEMERAL_DISK_SIZE}}
		}
		restart {
    		attempts = {{.RESTART_ATTEMPTS}}
    		delay = "{{.RESTART_DELAY}}"
    		interval = "{{.RESTART_INTERVAL}}"
    		mode = "{{.RESTART_MODE}}"
		}
		constraint {
			distinct_hosts = "{{.CONSTRAINT_DISTINCT_HOSTS}}"
		}
        # Create a web front end using a docker image
        ## task name
        task "{{.TASK}}" {
            driver = "{{.TASK_DRIVER}}"
            config {
				image = "{{.TASK_CONFIG_IMAGE}}"
				#labels to be published as Datadog tags
				labels {
					cobalt_id = "{{.TASK_CONFIG_LABELS_COBALT_ID}}"
					cobalt_service_name = "{{.TASK_CONFIG_LABELS_COBALT_SERVICE_NAME}}"
					cobalt_task = "{{.TASK_CONFIG_LABELS_COBALT_TASK}}"
					cobalt_podid = "{{.TASK_CONFIG_LABELS_COBALT_PODID}}"
				}
				# contiv-net
				network_mode = "{{.TASK_CONFIG_NETWORK_MODE}}"
				## for docker to map ports
				port_map {
							http_14840 = "{{.TASK_CONFIG_PORT_MAP_HTTP_14840}}"
							http_8080 = "{{.TASK_CONFIG_PORT_MAP_HTTP_8080}}"
				}
				# for registry login
				auth {
					username = "{{.TASK_CONFIG_AUTH_USERNAME}}"
					password = "{{.TASK_CONFIG_AUTH_PASSWORD}}"
				}
				privileged = {{.TASK_CONFIG_PRIVILEGED}}
				userns_mode = "{{.TASK_CONFIG_USERNS_MODE}}"
            }
            env {
				# for containerbuddy to talk to consul
				NOMAD_IP = "{{.TASK_ENV_NOMAD_IP}}"
				CONSUL_IP = "{{.TASK_ENV_CONSUL_IP}}"
				A1 = "{{.TASK_ENV_A1}}"
				AN_APP = "{{.TASK_ENV_AN_APP}}"
				AN_BUILD = "{{.TASK_ENV_AN_BUILD}}"
				AN_DOMAIN = "{{.TASK_ENV_AN_DOMAIN}}"
				AN_INSTANCEID = "{{.TASK_ENV_AN_INSTANCEID}}"
				AN_PORT = "{{.TASK_ENV_AN_PORT}}"
				COBALT_CONTAINER_TYPE = "{{.TASK_ENV_COBALT_CONTAINER_TYPE}}"
				COBALT_DB_PERSIST = "{{.TASK_ENV_COBALT_DB_PERSIST}}"
				COBALT_ID = "{{.TASK_ENV_COBALT_ID}}"
				COBALT_LOG_PERSIST = "{{.TASK_ENV_COBALT_LOG_PERSIST}}"
				COBALT_MODE = "{{.TASK_ENV_COBALT_MODE}}"
				COBALT_MODE_INTERFACE = "{{.TASK_ENV_COBALT_MODE_INTERFACE}}"
				COBALT_PODID = "{{.TASK_ENV_COBALT_PODID}}"
				COBALT_SERVICE_NAME = "{{.TASK_ENV_COBALT_SERVICE_NAME}}"
				COBALT_SERVICE_VERSION = "{{.TASK_ENV_COBALT_SERVICE_VERSION}}"
				COBALT_WS = "{{.TASK_ENV_COBALT_WS}}"
				ENTRY = "{{.TASK_ENV_ENTRY}}"
				ORACLE_COBALT_ID = "{{.TASK_ENV_ORACLE_COBALT_ID}}"
				ORACLE_SERVICE_NAME = "{{.TASK_ENV_ORACLE_SERVICE_NAME}}"
            }
			logs {
				max_files = {{.TASK_LOGS_MAX_FILES}}
    			max_file_size = {{.TASK_LOGS_MAX_FILES_SIZE}}
			}
            resources {
				cpu = {{.TASK_RESOURCES_CPU}}
				memory = {{.TASK_RESOURCES_MEMORY}}
				network {
					mbits = {{.TASK_RESOURCES_NETWORK_MBITS}}
					port "http_14840" {
					}
					port "http_8080" {
					}
					port "stats" {
					}
					port "cobalt"{
					}
				}
				iops = {{.TASK_RESOURCES_IOPS}}
			}
        }
		meta {
			is.entry = "{{.META_IS_ENTRY}}"
		}
	}
}