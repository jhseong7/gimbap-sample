package microservice_sample

import (
	"time"

	"github.com/jhseong7/ecl"
	"github.com/jhseong7/gimbap-sample/server-engine/food"
	"github.com/jhseong7/gimbap-sample/server-engine/salt"
	m "github.com/jhseong7/gimbap/microservice"
)

type (
	OtherSampleMicroService struct {
		m.IMicroService
		logger ecl.Logger
		Salt   *salt.SaltService
		Food   *food.FoodService
	}
)

func (s *OtherSampleMicroService) Start() {
	// Intentionally sleep for 2 seconds to simulate a long-running process --> this will show a warning message
	time.Sleep(7 * time.Second)
	s.logger.Log("OtherSampleMicroService Started " + s.Salt.GetName() + s.Food.GetName())
}

func (s *OtherSampleMicroService) Stop() {
	// Intentionally sleep for 2 seconds to simulate a long-running process
	time.Sleep(2 * time.Second)

	s.logger.Log("OtherSampleMicroService Stopped")
}

// Create a constructor with the desired dependencies from the DI container
func CreateOtherSampleMicroService(salt *salt.SaltService, food *food.FoodService) *OtherSampleMicroService {
	return &OtherSampleMicroService{
		logger: ecl.NewLogger(ecl.LoggerOption{
			Name: "OtherSampleMicroService",
		}),
		Salt: salt,
		Food: food,
	}
}

// Define a microservice provider with the utility function provided by GIMBAP
var OtherSampleMicroServiceProvider = m.DefineMicroService(m.MicroServiceProviderOption{
	Name:         "OtherSampleMicroService",
	Instantiator: CreateOtherSampleMicroService,
})
