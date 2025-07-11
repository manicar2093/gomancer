package core

import (
	"embed"
	"github.com/invopop/ctxi18n"
	"github.com/invopop/ctxi18n/i18n"
	"github.com/labstack/echo/v4"
)

type SessionSecretKeyConfig struct {
	SessionSecretKey string `env:"SESSION_SECRET_KEY" validate:"required"`
}

func SessionSecretKeyMiddleware(sessionSecret SessionSecretKeyConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(SessionSecretKey, sessionSecret)
			return next(c)
		}
	}
}

func I18nMiddleware(embedLocalizations embed.FS, defaultLocale i18n.Code) echo.MiddlewareFunc {
	if defaultLocale == "" {
		defaultLocale = ctxi18n.DefaultLocale
	}
	if err := ctxi18n.LoadWithDefault(embedLocalizations, defaultLocale); err != nil {
		panic(err)
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			i18nContext, err := ctxi18n.WithLocale(c.Request().Context(), c.Request().Header.Get("Accept-Language"))
			if err != nil {
				return err
			}
			c.SetRequest(c.Request().WithContext(i18nContext))
			return next(c)
		}
	}
}
