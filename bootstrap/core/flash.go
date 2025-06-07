package core

import (
	"encoding/gob"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

const (
	SessionSecretKey = "session_key"
	sessionName      = "flash_messages"
)

type FlashMessage struct {
	Variant string
	Message string
	Title   string
}

func (c *EchoContextTemplWrapper) GetFlash() []FlashMessage {
	session, _ := getCookieStore(c.Context).Get(c.Request(), sessionName)
	fm := session.Flashes()
	// If we have some messages.
	if len(fm) > 0 {
		session.Save(c.Request(), c.Response())
		// Initiate a strings slice to return messages.
		var flashes []FlashMessage
		for _, fl := range fm {
			flashes = append(flashes, fl.(FlashMessage))
		}

		return flashes
	}
	return nil
}

func SetFlash(ctx echo.Context, flash FlashMessage) {
	session, _ := getCookieStore(ctx).Get(ctx.Request(), sessionName)
	gob.Register(FlashMessage{})
	session.AddFlash(flash)

	if err := session.Save(ctx.Request(), ctx.Response()); err != nil {
		logrus.Error(err)
	}
}

func getCookieStore(ctx echo.Context) *sessions.CookieStore {
	// In real-world applications, use env variables to store the session key.
	sessionKey := ctx.Get(SessionSecretKey).(SessionSecretKeyConfig)
	return sessions.NewCookieStore([]byte(sessionKey.SessionSecretKey))
}
