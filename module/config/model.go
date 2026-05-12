package config

import (
	coreConfig "sbb-golang-template/pkg/config"
)

type CustomConfig struct{}

type CustomConfigFile struct {
	coreConfig.ConfigFile `mapstructure:",squash"`
	CustomConfig          `mapstructure:",squash"`
}
