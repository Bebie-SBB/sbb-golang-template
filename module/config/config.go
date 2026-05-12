package config

import (
	coreConfig "sbb-golang-template/pkg/config"
)

func GetProjectConfig() *CustomConfigFile {
	base := coreConfig.GetConfig()
	if base == nil {
		return nil
	}
	return &CustomConfigFile{
		ConfigFile: *base,
	}
}
