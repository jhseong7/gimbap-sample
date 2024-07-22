package salt

import (
	"fmt"

	"github.com/jhseong7/gimbap"
	"github.com/labstack/echo/v4"
)

type (
	SaltControllerEcho struct {
		gimbap.IController
		saltSvc SaltService
	}
)

func (c *SaltControllerEcho) GetRouteSpecs() []gimbap.RouteSpec {
	return []gimbap.RouteSpec{
		{Method: "GET", Path: "/salt", Handler: c.GetSaltEcho},
		{Method: "GET", Path: "/salt/:id", Handler: c.GetSaltEcho},
		{Method: "POST", Path: "/salt", Handler: c.PostSaltEcho},
	}
}

func (c *SaltControllerEcho) GetSaltEcho(ctx echo.Context) error {
	id := ctx.Param("id")

	if id == "" {
		id = "default"
	}

	ctx.String(200, fmt.Sprintf("Get Salt: %s %s", id, c.saltSvc.GetName()))

	return nil
}

func (c *SaltControllerEcho) PostSaltEcho(ctx echo.Context) error {
	ctx.String(200, "Post Salt "+c.saltSvc.GetName())

	return nil
}

func NewSaltControllerEcho(saltSvc *SaltService) *SaltControllerEcho {
	return &SaltControllerEcho{
		saltSvc: *saltSvc,
	}
}

var SaltControllerEchoProvider = gimbap.DefineController(
	gimbap.ControllerOption{
		Name:         "SaltControllerEcho",
		Instantiator: NewSaltControllerEcho,
		RootPath:     "salt",
	})
