package structs

import (
_"github.com/grafov/bcast"
)


const (
	CONSUL = "consul"
)

//global instructs.

type NCVConfiguration struct {
	NomadJobDockerNum  int
	NomadJobExecRawNum int
	WorkerTimeout      int
	WorkerUndoTimeout  int
	Registry           map[string]string
	EnvInfraType       string
	EnvOrchType        string
	EnvNetworkType     string
	RegUsername        string
	RegPassword        string
	DefaultConsulURL   string
	DefaultConsulPort  string
	DefaultNomadURL    string
	DefaultNomadPort   string
	VaultAddr          string
}

// Orch define the interfaces for Nomad api client to implement
type Orch interface {
	//input nomad config object/template
	SubmitJob(interface{}) (error, interface{})
	//input jobid and service name
	KillJob(string, string) (error, bool)
}
