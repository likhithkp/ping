package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/likhithkp/ping/application/user/handler"
)

type UserController struct {
	signUpHandler *handler.SignUpHandler
	signInHandler *handler.SignInHandler
}

func NewUserController(
	signUpHandler *handler.SignUpHandler,
	signInHandler *handler.SignInHandler,
) *UserController {
	return &UserController{
		signUpHandler: signUpHandler,
		signInHandler: signInHandler,
	}
}

func (userController *UserController) SignUp(c *fiber.Ctx) error {
	return userController.signUpHandler.SignUp(c)
}

func (userController *UserController) SignIn(c *fiber.Ctx) error {
	return userController.signInHandler.SignIn(c)
}
