package food

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jhseong7/gimbap"
)

type (
	FoodControllerFiber struct {
		gimbap.IController
		Food FoodService
	}
)

func (c *FoodControllerFiber) GetRouteSpecs() []gimbap.RouteSpec {
	return []gimbap.RouteSpec{
		{Method: "GET", Path: "/", Handler: c.GetFood},
		{Method: "POST", Path: "/", Handler: c.PostFood},
	}
}

func (c *FoodControllerFiber) GetFood(ctx *fiber.Ctx) error {
	ctx.Status(200).SendString(fmt.Sprintf("Food: %s, Salt: %s and combined is %s", c.Food.Name, c.Food.Salt.Name, c.Food.GetName()))

	return nil
}

func (c *FoodControllerFiber) PostFood(ctx *fiber.Ctx) error {
	ctx.Status(200).SendString("Post Food " + c.Food.GetName())

	return nil
}

func NewFoodControllerFiber(food *FoodService) *FoodControllerFiber {
	return &FoodControllerFiber{
		Food: *food,
	}
}

var FoodControllerFiberProvider = gimbap.DefineController(gimbap.ControllerOption{
	Name:         "FoodControllerFiber",
	Instantiator: NewFoodControllerFiber,
	RootPath:     "food",
})
