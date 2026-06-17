package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/likhithkp/ping/application/user/handler"
)

type Controller struct {
	getProfileHandler  *handler.GetProfileHandler
	editProfileHandler *handler.EditProfileHandler
}

func NewController(
	getProfileHandler *handler.GetProfileHandler,
	editProfileHandler *handler.EditProfileHandler,
) *Controller {
	return &Controller{
		getProfileHandler:  getProfileHandler,
		editProfileHandler: editProfileHandler,
	}
}

func (controller *Controller) SignUp(c *fiber.Ctx) error {
	return controller.getProfileHandler.GetProfile(c)
}

func (controller *Controller) SignIn(c *fiber.Ctx) error {
	return controller.editProfileHandler.EditProfile(c)
}
