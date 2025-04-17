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

// ParseConfigWithOptions creates an instance of needed struct passing some env.Options. It is necesary struct contains env and validate tags to be parsed correctly
func ParseConfigWithOptions[T any](opts env.Options) (T, error) {
	var instance T
	if err := env.ParseWithOptions(&instance, opts); err != nil {
		return instance, err
	}

	v := validator.NewGooKitValidator()
	if err := v.ValidateStruct(&instance); err != nil {
		return instance, err
	}
	return instance, nil
}

// ParseConfig creates an instance of needed struct. It is necesary struct contains env and validate tags to be parsed correctly
func ParseConfig[T any]() (T, error) {
	var instance T
	if err := env.Parse(&instance); err != nil {
		return instance, err
	}

	v := validator.NewGooKitValidator()
	if err := v.ValidateStruct(&instance); err != nil {
		return instance, err
	}
	return instance, nil
}
