package salt

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jhseong7/gimbap"
)

type (
	SaltControllerFiber struct {
		gimbap.IController
		saltSvc SaltService
	}
)

func (c *SaltControllerFiber) GetRouteSpecs() []gimbap.RouteSpec {
	return []gimbap.RouteSpec{
		{Method: "GET", Path: "/salt", Handler: c.GetSaltFiber},
		{Method: "GET", Path: "/salt/:id", Handler: c.GetSaltFiber},
		{Method: "POST", Path: "/salt", Handler: c.PostSaltFiber},
	}
}

func (c *SaltControllerFiber) GetSaltFiber(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if id == "" {
		id = "default"
	}

	ctx.Status(200).SendString(fmt.Sprintf("Get Salt: %s %s", id, c.saltSvc.GetName()))

	return nil
}

func (c *SaltControllerFiber) PostSaltFiber(ctx *fiber.Ctx) error {
	ctx.Status(200).SendString("Post Salt " + c.saltSvc.GetName())

	return nil
}

func NewSaltControllerFiber(saltSvc *SaltService) *SaltControllerFiber {
	return &SaltControllerFiber{
		saltSvc: *saltSvc,
	}
}

var SaltControllerFiberProvider = gimbap.DefineController(
	gimbap.ControllerOption{
		Name:         "SaltControllerFiber",
		Instantiator: NewSaltControllerFiber,
		RootPath:     "salt",
	})
