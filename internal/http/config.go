package http

import (
	"os"
	"strconv"
)

type Config struct {
	ServerPort uint16
}

func LoadConfig() Config {
	cfg := Config{
		ServerPort: 3000,
	}

	if portStr, exist := os.LookupEnv("PORT"); exist {
		if port, err := strconv.ParseUint(portStr, 10, 16); err == nil {
			cfg.ServerPort = uint16(port)
		}
	}
	return cfg
}
