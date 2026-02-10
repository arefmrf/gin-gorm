package config

type Config struct {
	App    App
	Server Server
	DB     DB `mapstructure:"db"`
	TestDB DB `mapstructure:"test_db"`
}

type App struct {
	Name string
}

type Server struct {
	Host string
	Port string
}

type DB struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

type TestDB struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}
