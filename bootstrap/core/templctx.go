package core

import (
	"context"
	"github.com/labstack/echo/v4"
	"time"
)

const EchoContextTemplWrapperKey = "echoCtx"

type EchoContextTemplWrapper struct {
	echo.Context
}

func (c *EchoContextTemplWrapper) Deadline() (deadline time.Time, ok bool) {
	return c.Context.Request().Context().Deadline()
}

func (c *EchoContextTemplWrapper) Done() <-chan struct{} {
	return c.Request().Context().Done()
}

func (c *EchoContextTemplWrapper) Err() error {
	return c.Context.Request().Context().Err()
}

func (c *EchoContextTemplWrapper) Value(key any) any {
	return c.Request().Context().Value(key)
}

func (c *EchoContextTemplWrapper) GetCqrfToken() string {
	return c.Context.Get("csrf").(string)
}

func GetEchoTemplContext(ctx context.Context) *EchoContextTemplWrapper {
	return ctx.Value(EchoContextTemplWrapperKey).(*EchoContextTemplWrapper)
}
