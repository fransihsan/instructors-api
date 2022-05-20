package configs

type AppConfig struct {
	PORT        int16
	DB_DRIVER   string
	DB_NAME     string
	DB_PORT     int16
	DB_HOST     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_LOC      string
}