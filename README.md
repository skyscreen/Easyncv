
EasyNCV version 0.2

This is a implementation tool for deploying cloud on nomad/consul/vault

1. can support nomad deploy using your custom template, destroy job
2. can support job/group/task template config on page(so far only 1)
3. can support consul key/value create/delete
4. can support vault init->unseal->enableAuth


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

      config consul.json for consul configuration e.g.

      {
        "framework":"easyncv",
        "version":"0.2",
        "url":"10.173.76.57:8500"

      }

      2) page run
      start main.go
      http://localhost:8080
      input values and submit to start job

      http://localhost:8080/stop to stop job

       command line
         run deploy.go
         run destroy.go



