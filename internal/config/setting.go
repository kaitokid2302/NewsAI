package config

var (
	Global Config
)

type Config struct {
	Database
}

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}
