package http

import (
	"fmt"
	"time"
)

const (
	defaultServerName         = "HttpServer"
	defaultServerHost         = "0.0.0.0"
	defaultServerPort         = "8080"
	defaultServerReadTimeout  = 15 * time.Second
	defaultServerWriteTimeout = 30 * time.Second
)

type ServerConfig struct {
	Name         string        `yaml:"name"`
	Host         string        `yaml:"host"`
	Port         string        `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

func (c *ServerConfig) Address() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}
