package redis

import "fmt"

type Config struct {
	Host     string `mapstructure:"host" json:"host"`
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
	Port     int    `mapstructure:"port" json:"port"`
	Prefix   string `mapstructure:"prefix" json:"prefix"`
}

func (cfg *Config) GetUri() (uri string) {
	usernamePass := ""
	if cfg.Username != "" {
		usernamePass += fmt.Sprintf("%s", cfg.Username)
	}

	if cfg.Password != "" {
		usernamePass += fmt.Sprintf(":%s", cfg.Password)
	}

	if usernamePass != "" {
		usernamePass += "@"
	}

	uri = fmt.Sprintf("redis://%s%s:%d", usernamePass, cfg.Host, cfg.Port)
	return
}
