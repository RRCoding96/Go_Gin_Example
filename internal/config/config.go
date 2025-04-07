package config

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewConfig() *Config {
	return &Config{
		DB: DBConfig{
			Host:     "localhost",
			Port:     "33066",
			User:     "root",
			Password: "1234",
			DBName:   "gin_example",
		},
	}
}
