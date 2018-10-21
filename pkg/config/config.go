package config

import (
	"os"

	"github.com/golang/glog"
)

// config from environment
var (
	MySQLDSN string

	envs = map[string]*string{
		"VERICRY_MYSQL_DSN": &MySQLDSN,
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
	glog.V(4).Infof("VERICRY_MYSQL_DSN: %s", MySQLDSN)
}
