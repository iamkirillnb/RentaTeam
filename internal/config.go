package internal

import (
	"fmt"
	"github.com/iamkirillnb/Rentateam/internal/entities"
	"github.com/iamkirillnb/Rentateam/pkg/http"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
	"time"
)

type Config struct {
	DbConfig entities.DbRepo   `yaml:"postgres_local"`
	Server   http.ServerConfig `yaml:"server_http"`
}

type ServerConfig struct {
	Name         string        `yaml:"name"`
	Host         string        `yaml:"host"`
	Port         string        `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

var instance *Config
var once sync.Once

func (c *ServerConfig) Address() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func GetConfig(path string) *Config {
	once.Do(func() {
		instance = &Config{}

		if err := cleanenv.ReadConfig(path, instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Println(help)
			log.Fatal(err)
		}
	})
	return instance
}
