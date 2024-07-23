package microservice_sample

import (
	"github.com/jhseong7/ecl"
	"github.com/jhseong7/gimbap-sample/server-engine/salt"
	m "github.com/jhseong7/gimbap/microservice"
)

type (
	SampleMicroService struct {
		m.IMicroService
		logger ecl.Logger
		Salt   *salt.SaltService
	}
)

func (s *SampleMicroService) Start() {
	s.logger.Log("SampleMicroService Started " + s.Salt.GetName())
}

func (s *SampleMicroService) Stop() {
	s.logger.Log("SampleMicroService Started")
}

// Create a constructor with the desired dependencies from the DI container
func CreateSampleMicroService(salt *salt.SaltService) *SampleMicroService {
	return &SampleMicroService{
		logger: ecl.NewLogger(ecl.LoggerOption{
			Name: "SampleMicroService",
		}),
		Salt: salt,
	}
}

// Define a microservice provider with the utility function provided by GIMBAP
var MicroserviceProvider = m.DefineMicroService(m.MicroServiceProviderOption{
	Name:         "SampleMicroService",
	Instantiator: CreateSampleMicroService,
})
