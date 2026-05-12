package database

import (
	"fmt"
	"os"

	"sbb-golang-template/pkg/config"
	"sbb-golang-template/pkg/interfaces"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseService struct {
	databases   map[string]*gorm.DB
	defaultName string
}

var dbService *DatabaseService

type databaseModule struct {
	service *DatabaseService
}

func (m *databaseModule) Init() error {
	return initDatabases()
}

func (m *databaseModule) SetRoutes() error {
	return nil
}

func GetInstance() interfaces.IModule {
	mod := &databaseModule{}
	if err := mod.Init(); err != nil {
		fmt.Println("database init error:", err)
	}
	return mod
}

func GetDatabaseService() *DatabaseService {
	return dbService
}

func (s *DatabaseService) DefaultDatabase() *gorm.DB {
	if s == nil {
		return nil
	}
	if s.defaultName != "" {
		return s.databases[s.defaultName]
	}
	for _, db := range s.databases {
		return db
	}
	return nil
}

func (s *DatabaseService) GetDatabase(name string) *gorm.DB {
	if s == nil {
		return nil
	}
	return s.databases[name]
}

func initDatabases() error {
	cfg := config.GetConfig()
	if len(cfg.Databases) == 0 {
		return nil
	}

	dbService = &DatabaseService{
		databases: make(map[string]*gorm.DB),
	}

	gormCfg := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}

	for _, dbCfg := range cfg.Databases {
		db, err := openDatabase(dbCfg, gormCfg)
		if err != nil {
			return fmt.Errorf("failed to connect to database %s: %w", dbCfg.Name, err)
		}
		dbService.databases[dbCfg.Name] = db

		if dbService.defaultName == "" || dbCfg.Name == cfg.Configs.DatabaseList {
			dbService.defaultName = dbCfg.Name
		}
	}

	return nil
}

func openDatabase(dbCfg config.Database, gormCfg *gorm.Config) (*gorm.DB, error) {
	switch dbCfg.Config.Dialect {
	case "sqlite":
		dbPath := dbCfg.Config.Name
		if dbPath == "" {
			dbPath = "app.db"
		}
		return gorm.Open(sqlite.Open(dbPath), gormCfg)

	case "postgres":
		dsn := buildPostgresDSN(dbCfg.Config)
		return gorm.Open(postgres.Open(dsn), gormCfg)

	case "mysql":
		dsn := buildMySQLDSN(dbCfg.Config)
		return gorm.Open(mysql.Open(dsn), gormCfg)

	default:
		return nil, fmt.Errorf("unsupported dialect: %s", dbCfg.Config.Dialect)
	}
}

func buildPostgresDSN(cfg config.DatabaseConfig) string {
	host := getEnvOrDefault("DB_HOST", cfg.Host)
	port := getEnvOrDefault("DB_PORT", fmt.Sprintf("%d", cfg.Port))
	user := getEnvOrDefault("DB_USERNAME", cfg.Username)
	pass := getEnvOrDefault("DB_PASSWORD", cfg.Password)
	name := getEnvOrDefault("DB_NAME", cfg.Name)
	ssl := getEnvOrDefault("DB_SSLMODE", cfg.SSLMode)
	if ssl == "" {
		ssl = "disable"
	}
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, pass, name, ssl)
}

func buildMySQLDSN(cfg config.DatabaseConfig) string {
	host := getEnvOrDefault("DB_HOST", cfg.Host)
	port := getEnvOrDefault("DB_PORT", fmt.Sprintf("%d", cfg.Port))
	user := getEnvOrDefault("DB_USERNAME", cfg.Username)
	pass := getEnvOrDefault("DB_PASSWORD", cfg.Password)
	name := getEnvOrDefault("DB_NAME", cfg.Name)
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)
}

func getEnvOrDefault(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
