package config

var (
	Global Config
)

type Config struct {
	Database
	GoogleEmail `mapstructure:"google_email"`
	Redis
	JWT
}

type GoogleEmail struct {
	AppPassword string `mapstructure:"app_password"`
	From        string `mapstructure:"from"`
}

type JWT struct {
	Key string
}

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type Redis struct {
	Host string
	Port int
}
