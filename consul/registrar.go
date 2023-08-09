package consul

import (
	stdconsul "github.com/hashicorp/consul/api"

	l4g "github.com/libra9z/log4go"
)

// Registrar registers service instance liveness information to Consul.
type Registrar struct {
	client       Client
	registration *stdconsul.AgentServiceRegistration
	logger       l4g.Logger
}

// NewRegistrar returns a Consul Registrar acting on the provided catalog
// registration.
func NewRegistrar(client Client, r *stdconsul.AgentServiceRegistration, logger l4g.Logger) *Registrar {
	return &Registrar{
		client:       client,
		registration: r,
		logger:       logger,
	}
}

// Register implements sd.Registrar interface.
func (p *Registrar) Register() {
	if err := p.client.Register(p.registration); err != nil {
		p.logger.Error("error=%v", err)
	} else {
		p.logger.Info("action=register")
	}
}

// Deregister implements sd.Registrar interface.
func (p *Registrar) Deregister() {
	if err := p.client.Deregister(p.registration); err != nil {
		p.logger.Error("error=%v", err)
	} else {
		p.logger.Info("action=deregister")
	}
}
