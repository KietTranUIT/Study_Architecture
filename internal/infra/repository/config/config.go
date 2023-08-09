package config

import "fmt"

type ConfigDatabase struct {
	Driver   string
	Username string
	Password string
	Host     string
	Port     uint
	DbName   string
}

func (conf ConfigDatabase) URL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.DbName)
}
