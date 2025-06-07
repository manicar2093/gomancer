package core

import (
	"github.com/labstack/echo/v4"
)

type Controller interface {
	SetUpRoutes(*echo.Group)
}

func RegisterControllers(group *echo.Group, contrs ...Controller) {
	for _, contr := range contrs {
		RegisterController(group, contr)
	}
}

func RegisterController(group *echo.Group, contr Controller) {
	contr.SetUpRoutes(group)
}
