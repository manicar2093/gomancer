package httphealthcheck

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthChecker struct {
	checkers []Checkable
}

func NewHealthChecker(checkers ...Checkable) *HealthChecker {
	return &HealthChecker{
		checkers: checkers,
	}
}

func (c *HealthChecker) SetUpRoutes(group *echo.Group) {
	group.GET("/health_check", c.echoController)
}

// @Summary		Gets services status
// @Description	Gets services status of registered services
// @Tags			healthcheck
// @Produce		json
// @Success		200	{object}	healthResponse			"Health status data"
// @Failure		410	{object}	apperrors.MessagedError	"A sevice has an error"
// @Failure		500	"Something unidentified has occurred"
// @Router			/health_check [get]
func (c *HealthChecker) echoController(ctx echo.Context) error {
	healths := make(map[string]HealthStatusData)
	errs := make(map[string]error)
	msg := "everything seems ok"
	code := http.StatusOK

	for _, checker := range c.checkers {
		health, err := checker.ServiceHealth()
		if err != nil {
			msg = "a checker has an unexpected error"
			code = http.StatusGone
			errs[checker.ServiceName()] = err
			continue
		}
		healths[checker.ServiceName()] = health
	}

	return ctx.JSON(code, healthResponse{
		Errors:  errs,
		Healths: healths,
		Msg:     msg,
	})
}
