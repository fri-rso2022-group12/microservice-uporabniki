package initializers

import (
	"github.com/hashicorp/consul/api"
	"os"
)

var ConsulClient *api.Client
var ConsulKV *api.KV

func InitializeConsul() {
	// Get a new client
	var err error
	ConsulClient, err = api.NewClient(&api.Config{Address: os.Getenv("CONSUL_SERVER")})
	if err != nil {
		panic(err)
	}
	ConsulKV = ConsulClient.KV()
}
