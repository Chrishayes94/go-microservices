package registration

import (
	"log"
	"time"
)

func (s *Registration) UpdateTTL(check func() (bool, error)) {
	ticker := time.NewTicker(s.TTL / 2)
	for range ticker.C {
		s.update(c)
	}
}

func (s *Registration) update(check func() (bool, error)) {
	if agentErr := s.ConsulAgent.PassTTL("service:"+s.Name, ""); agentErr != nil {
		log.Print(agentErr)
	}
}


