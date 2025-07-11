package core

import (
	"context"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)
	renderContext := context.WithValue(ctx.Request().Context(), EchoContextTemplWrapperKey, &EchoContextTemplWrapper{ctx})

	if err := t.Render(renderContext, buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}
