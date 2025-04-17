package core

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/manicar2093/echoroutesview"
	"github.com/manicar2093/gomancer/bootstrap/core/apperrors"
	"github.com/manicar2093/gomancer/bootstrap/core/echoer"
	"github.com/manicar2093/gomancer/bootstrap/core/logger"
	"github.com/manicar2093/gomancer/bootstrap/core/stages"
	"github.com/manicar2093/gomancer/bootstrap/core/validator"

	echoSwagger "github.com/swaggo/echo-swagger"
)

type (
	Server struct {
		*echo.Echo
		gookitValidator *validator.GooKitValidator
		baseEndpoint    *echo.Group
		controllers     []Controller
		stage           string
	}

	ServerConfig struct {
		EchoInstance *echo.Echo
		BaseEndpoint string
	}

	Controller interface {
		SetUpRoutes(*echo.Group)
	}
)

func NewServer(
	conf ServerConfig,
	controllers ...Controller,
) *Server {
	stage := getStage()
	server := &Server{
		Echo:            conf.EchoInstance,
		gookitValidator: validator.NewGooKitValidator(),
		baseEndpoint:    conf.EchoInstance.Group(conf.BaseEndpoint),
		controllers:     controllers,
		stage:           stage,
	}

	server.configEcho()
	server.configControllers()

	return server
}

func (c *Server) Start(address string) error {
	return c.Echo.Start(address)
}

func (c *Server) configEcho() {
	c.HideBanner = true
	c.Use(middleware.Recover())
	c.Use(middleware.RequestID())
	c.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"msg":"request to ${uri}","data":{"latency_human":"${latency_human}","id":"${id}","method":"${method}","uri":"${uri}","status":"${status}"},"time":"${time_rfc3339_nano}"}` + "\n",
	}))
	c.Use(middleware.CORS())
	c.Validator = c.gookitValidator
	c.HTTPErrorHandler = apperrors.HandlerWEcho
	logger.GetLogger().Debug("swagger docs available")
	c.GET("/swagger/*", echoSwagger.WrapHandler)
}

func (c *Server) configControllers() {
	for _, controller := range c.controllers {
		controller.SetUpRoutes(c.baseEndpoint)
	}
	if c.stage == stages.Dev {
		logger.GetLogger().Info("/echo and /registered-routes endpoint registered")
		echoer.RegisterEchoer(c.Echo)               //nolint:errcheck
		echoroutesview.RegisterRoutesViewer(c.Echo) //nolint:errcheck
	}
}

func getStage() string {
	stage, err := stages.GetCurrentStage()
	if err != nil {
		panic(err)
	}
	return stage
}
