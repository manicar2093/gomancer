package echoer

import (
	"github.com/manicar2093/gomancer/bootstrap/core/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterEchoer(echoInstance *echo.Echo) error {
	echoInstance.GET("/echo", func(ctx echo.Context) error {
		var body = make(map[string]interface{})
		if err := ctx.Bind(&body); err != nil {
			logger.GetLogger().WithField("message", "error on /echo handler").Error(err)
			return err
		}
		logger.GetLogger().Info("echoed data received")
		return ctx.JSON(http.StatusAccepted, body)
	})
	return nil
}
