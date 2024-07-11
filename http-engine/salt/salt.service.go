package salt

import "github.com/jhseong7/gimbap"

type (
	SaltService struct {
		Name string
	}
)

func (s *SaltService) GetName() string {
	return s.Name
}

func NewSalt() *SaltService {
	return &SaltService{
		Name: "Salty Salt!!",
	}
}

var SaltProvider = gimbap.DefineProvider(gimbap.ProviderOption{
	Name:         "SaltService",
	Instantiator: NewSalt,
})
