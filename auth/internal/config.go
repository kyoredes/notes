package internal

import (
	"os"
)

type Config struct {
	Debug                  bool
	JWTSecretKey           string
	AccessTokenExpiration  int
	RefreshTokenExpiration int
}

type DBConfig struct {
	Host      string
	Port      string
	User      string
	Password  string
	DBName    string
	DBTimeout int
}

func getEnv(key string, defaultValue string) any {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defaultValue
}

func NewConfig() *Config {
	debug := getEnv("DEBUG", "false").(bool)
	jwtSecretKey := getEnv("JWT_SECRET_KEY", "secret").(string)
	accessTokenExpiration := getEnv("ACCESS_TOKEN_EXPIRATION", "3600").(int)
	refreshTokenExpiration := getEnv("REFRESH_TOKEN_EXPIRATION", "86400").(int)

	return &Config{
		Debug:                  debug,
		JWTSecretKey:           jwtSecretKey,
		AccessTokenExpiration:  accessTokenExpiration,
		RefreshTokenExpiration: refreshTokenExpiration,
	}
}

func NewDBConfig() *DBConfig {
	host := getEnv("DB_HOST", "localhost").(string)
	port := getEnv("DB_PORT", "5432").(string)
	user := getEnv("DB_USER", "postgres").(string)
	password := getEnv("DB_PASSWORD", "postgres").(string)
	dbName := getEnv("DB_NAME", "postgres").(string)
	dbTimeout := getEnv("DB_TIMEOUT", "10").(int)

	return &DBConfig{
		Host:      host,
		Port:      port,
		User:      user,
		Password:  password,
		DBName:    dbName,
		DBTimeout: dbTimeout,
	}
}
