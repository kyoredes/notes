package internal

import (
	"github.com/spf13/viper"
)

type Config struct {
	Debug                  bool
	LoggingMode            string
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

func NewConfig() *Config {
	debug := viper.GetBool("DEBUG")
	loggingMode := viper.GetString("LOGGING_MODE")
	jwtSecretKey := viper.GetString("JWT_SECRET_KEY")
	accessTokenExpiration := viper.GetInt("ACCESS_TOKEN_EXPIRATION")
	refreshTokenExpiration := viper.GetInt("REFRESH_TOKEN_EXPIRATION")

	return &Config{
		Debug:                  debug,
		LoggingMode:            loggingMode,
		JWTSecretKey:           jwtSecretKey,
		AccessTokenExpiration:  accessTokenExpiration,
		RefreshTokenExpiration: refreshTokenExpiration,
	}
}

func NewDBConfig() *DBConfig {
	host := viper.GetString("DB_HOST")
	port := viper.GetString("DB_PORT")
	user := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASSWORD")
	dbName := viper.GetString("DB_NAME")
	dbTimeout := viper.GetInt("DB_TIMEOUT")

	return &DBConfig{
		Host:      host,
		Port:      port,
		User:      user,
		Password:  password,
		DBName:    dbName,
		DBTimeout: dbTimeout,
	}
}

func Init() {
	viper.AutomaticEnv()
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_PASSWORD", "postgres")
	viper.SetDefault("DB_NAME", "postgres")
	viper.SetDefault("DB_TIMEOUT", 10)

	viper.SetDefault("DEBUG", false)
	viper.SetDefault("LOGGING_MODE", "text")
	viper.SetDefault("JWT_SECRET_KEY", "secret")
	viper.SetDefault("ACCESS_TOKEN_EXPIRATION", 3600)
	viper.SetDefault("REFRESH_TOKEN_EXPIRATION", 86400)
}
