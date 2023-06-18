package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/taoti888/user/global"
)

func NewConsulRegistrar(ip, uuid string) error {

	// counsel配置信息
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.CONFIG.Consul.Address, global.CONFIG.Consul.Port)
	client, err := api.NewClient(cfg)
	if err != nil {
		panic("failed to init consul client agent, error: " + err.Error())
	}

	// 全局Agent
	global.Agent = client.Agent()

	check := &api.AgentServiceCheck{
		Interval: "5s",
		Timeout:  "1s",
		HTTP:     fmt.Sprintf("http://%s:8081/health", ip),
		//GRPC:                           fmt.Sprintf("%s:%d/%s/Check", ip, global.CONFIG.System.Port, "proto.Health"),
		DeregisterCriticalServiceAfter: "1m",
	}
	registration := &api.AgentServiceRegistration{
		ID:      uuid,
		Name:    global.CONFIG.System.Name,
		Port:    global.CONFIG.System.Port,
		Address: ip,
		Tags:    global.CONFIG.System.Tags,
		Check:   check,
	}
	return global.Agent.ServiceRegister(registration)
}

func Deregister(uuid string) error {
	return global.Agent.ServiceDeregister(uuid)
}
