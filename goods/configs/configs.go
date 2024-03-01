package configs

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Addrs    string `env:"ADDRS"`
	Port     string `env:"PORT"`
	Host     string
	Username string
	DBName   string
	Password string
	SSLMode  string

	// TODO: add external confs
	// TODO: add db and api separation
}

func GetConfigs(filePath string) (Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig(filePath, &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
