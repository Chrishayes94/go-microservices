package registration

import (
	"time"
	consul "github.com/hashicorp/consul/api"
)

type Registration struct {
	Name			string
	TTL				time.Duration
	ConsulAgent		*consul.Agent
}

func New(integration string, ttl time.Duration, port int) (*Registration, error) {
	s := &Registration{
		Name: "Integration-" + integration,
		TTL: ttl,
	}
	config := consul.DefaultConfig()
	config.Address = "172.17.0.4"

	c, err := consul.NewClient(config)
	if err != nil {
		return nil, err
	}
	s.ConsulAgent = c.Agent()
	serviceDef := &consul.AgentServiceRegistration{
		Name:              s.Name,
		Tags:              []string{"development", "api"},
		Port:              port,
		Check: &consul.AgentServiceCheck{
			TTL: s.TTL.String(),
		},
	}
	if err := s.ConsulAgent.ServiceRegister(serviceDef); err != nil {
		return nil, err
	}
	//TODO check health of this node - go s.UpdateTTL(s.Check)
	return s, nil
}