package core

import (
	"github.com/caarlos0/env/v10"
	"github.com/manicar2093/gomancer/bootstrap/core/validator"
)

// Config is a struct helper to get all configuration need by winter
type Config struct {
	// Environment is used to get data using ParseConfig
	Environment string `env:"ENVIRONMENT" validate:"required|in:prod,dev,test"`
	Port        int    `env:"PORT" validate:"required|min:1|max:65535"`
}

// ParseConfigWithOptions creates an instance of the necessary struct passing some env.Options. It is necessary the struct contains env and validate tags to be parsed correctly
func ParseConfigWithOptions[T any](opts env.Options) (T, error) {
	var instance T
	if err := env.ParseWithOptions(&instance, opts); err != nil {
		return instance, err
	}

	v := validator.NewStructValidatorConfigured(&instance, "en")
	if err := validator.ValidateAndTransform(v); err != nil {
		return instance, err
	}
	return instance, nil
}

// ParseConfig creates an instance of the necessary struct. It is necessary the struct contains env and validate tags to be parsed correctly
func ParseConfig[T any]() (T, error) {
	var instance T
	if err := env.Parse(&instance); err != nil {
		return instance, err
	}

	v := validator.NewStructValidatorConfigured(&instance, "en")
	if err := validator.ValidateAndTransform(v); err != nil {
		return instance, err
	}
	return instance, nil
}
