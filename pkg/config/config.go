package config

import (
	"os"

	"github.com/golang/glog"
)

// config from environment
var (
	MySQLDSN   string
	ListenAddr string

	envs = map[string]*string{
		"VERICRY_MYSQL_DSN":   &MySQLDSN,
		"VERICRY_LISTEN_ADDR": &ListenAddr,
	}
)

func init() {
	for k, v := range envs {
		*v = os.Getenv(k)
	}
}

// Logging all config
func Logging() {
	glog.V(4).Infof("Vericry config:")
	for k, v := range envs {
		glog.V(4).Infof("%s: %s", k, *v)
	}
}
