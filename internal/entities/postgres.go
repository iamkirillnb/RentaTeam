package entities

import (
	"fmt"
)

type DbRepo struct {
	DbAppName string `yaml:"application_name"`
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Database  string `yaml:"database"`
	SslMode   string `yaml:"ssl_mode"`
}

func (d *DbRepo) DbUrl() string {
	dbUrl := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v&fallback_application_name=%v",
		d.User, d.Password, d.Host, d.Port, d.Database, d.SslMode, d.DbAppName)
	return dbUrl
}