package config

type CorsConfig struct {
	AllowHeaders []string `json:"allow_headers" mapstructure:"allow_headers"`
	AllowOrigin  []string `json:"allow_origin" mapstructure:"allow_origin"`
}

type TLSConfig struct {
	Enable bool   `json:"enable" mapstructure:"enable"`
	Cert   string `json:"cert" mapstructure:"cert"`
	Key    string `json:"key" mapstructure:"key"`
}

type ServerConfig struct {
	Name string      `json:"name" mapstructure:"name"`
	Port int         `json:"port" mapstructure:"port"`
	Cors *CorsConfig `json:"cors" mapstructure:"cors"`
	TLS  *TLSConfig  `json:"tls" mapstructure:"tls"`
}

type DatabaseConfig struct {
	Dialect  string `json:"dialect" mapstructure:"dialect"`
	Host     string `json:"host" mapstructure:"host"`
	Port     int    `json:"port" mapstructure:"port"`
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	Name     string `json:"name" mapstructure:"name"`
	SSLMode  string `json:"sslmode" mapstructure:"sslmode"`
}

type Database struct {
	Name   string         `json:"name" mapstructure:"name"`
	Config DatabaseConfig `json:"config" mapstructure:"config"`
}

type Configs struct {
	DatabaseList            string `json:"databaseList" mapstructure:"databaseList"`
	IsMigration             bool   `json:"isMigration" mapstructure:"isMigration"`
	TargetMigrationDatabase string `json:"targetMigrationDatabase" mapstructure:"targetMigrationDatabase"`
	DisplayResponseError    bool   `json:"displayResponseError" mapstructure:"displayResponseError"`
}

type ConfigFile struct {
	Server    ServerConfig `json:"server" mapstructure:"server"`
	Databases []Database   `json:"databases" mapstructure:"databases"`
	Configs   Configs      `json:"configs" mapstructure:"configs"`
}
