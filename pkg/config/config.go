package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cfg *ConfigFile

func Init(env string, loadEnvFile bool, loadConfFile bool) error {
	cfg = &ConfigFile{
		Server: ServerConfig{Port: 9090},
	}

	if loadEnvFile {
		loadDotEnv(".env")
	}

	if loadConfFile {
		path := fmt.Sprintf(".static/configs/files/%s.json", env)
		if data, err := os.ReadFile(path); err == nil {
			if err := json.Unmarshal(data, cfg); err != nil {
				return fmt.Errorf("config file parse error: %w", err)
			}
		}
	}

	overrideFromEnv()
	return nil
}

func loadDotEnv(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		return
	}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			val := strings.TrimSpace(parts[1])
			if os.Getenv(key) == "" {
				os.Setenv(key, val)
			}
		}
	}
}

func overrideFromEnv() {
	if port := os.Getenv("CONFIG_PORT"); port != "" {
		if p, err := strconv.Atoi(port); err == nil {
			cfg.Server.Port = p
		}
	}
	if port := os.Getenv("PORT"); port != "" {
		if p, err := strconv.Atoi(port); err == nil {
			cfg.Server.Port = p
		}
	}
	if name := os.Getenv("CONFIG_SERVICE_NAME"); name != "" {
		cfg.Server.Name = name
	}
	if dbList := os.Getenv("CONFIG_DATABASE_LIST"); dbList != "" {
		cfg.Configs.DatabaseList = dbList
		// Auto-build database configs from env vars when no config file is loaded
		if len(cfg.Databases) == 0 {
			for _, name := range strings.Split(dbList, ",") {
				name = strings.TrimSpace(name)
				dialect := os.Getenv("DB_DIALECT")
				if dialect == "" {
					dialect = name
				}
				dbPort := 0
				if p := os.Getenv("DB_PORT"); p != "" {
					dbPort, _ = strconv.Atoi(p)
				}
				cfg.Databases = append(cfg.Databases, Database{
					Name: name,
					Config: DatabaseConfig{
						Dialect:  dialect,
						Host:     os.Getenv("DB_HOST"),
						Port:     dbPort,
						Username: os.Getenv("DB_USERNAME"),
						Password: os.Getenv("DB_PASSWORD"),
						Name:     os.Getenv("DB_NAME"),
						SSLMode:  os.Getenv("DB_SSLMODE"),
					},
				})
			}
		}
	}
	if migration := os.Getenv("CONFIG_IS_MIGRATION"); migration == "true" {
		cfg.Configs.IsMigration = true
	}
	if target := os.Getenv("CONFIG_TARGET_MIGRATION_DATABASE"); target != "" {
		cfg.Configs.TargetMigrationDatabase = target
	}
	if display := os.Getenv("CONFIG_DISPLAY_RESPONSE_ERROR"); display == "true" {
		cfg.Configs.DisplayResponseError = true
	}
	if origin := os.Getenv("CONFIG_ORIGIN_ALLOWED"); origin != "" {
		if cfg.Server.Cors == nil {
			cfg.Server.Cors = &CorsConfig{}
		}
		cfg.Server.Cors.AllowOrigin = strings.Split(origin, ",")
	}
}

func GetConfig() *ConfigFile {
	if cfg == nil {
		cfg = &ConfigFile{Server: ServerConfig{Port: 9090}}
	}
	return cfg
}

func SetConfig(c *ConfigFile) {
	cfg = c
}
