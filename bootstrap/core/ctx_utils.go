package core

import (
	"github.com/invopop/ctxi18n/i18n"
	"github.com/labstack/echo/v4"
	"github.com/manicar2093/gomancer/bootstrap/core/validator"
)

func BindAndValidate(ctx echo.Context, i any) error {
	if err := ctx.Bind(i); err != nil {
		return err
	}
	if err := Validate(ctx, i); err != nil {
		return err
	}
	return nil
}

func Validate(ctx echo.Context, i any) error {
	lang := i18n.GetLocale(ctx.Request().Context()).Code()
	return validator.ValidateAndTransform(
		validator.NewStructValidatorConfigured(i, string(lang)),
	)
}

func GetFromEchoContext[T any](ctx echo.Context, key string) T {
	return ctx.Get(key).(T)
}
