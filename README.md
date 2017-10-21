
EasyNCV version 0.1

This is a implementation tool for deploying cloud on nomad/consul/vault

1. can support nomad deploy using your custom template, destroy job
2. can support job/group/task template config on page(so far only 1)


Usage:
      1)
       config hcl.json to start deploy e.g.
      {
        "run":"start",
        "jobid":"an-service-2017",
        "nomadurl": "10.173.76.57:8500",
        "hclfile": "hcl/example.hcl"
      }
       config hclstop.json to stop job e.g.
      {
        "run":"stop",
        "jobid":"an-service-2017",
        "nomadurl": "10.173.76.57:8500",
        "hclfile": "hcl/example.hcl"
      }
      2) page run
      start main.go
      http://localhost:8080
      input values and submit to start job

      http://localhost:8080/stop to stop job

      3) command line
         run deploy.go
         run destroy.go
