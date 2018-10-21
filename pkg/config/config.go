package config

import (
	"os"
)

// config from environment
var (
	MySQLDSN = os.Getenv("MYSQL_DSN")
)
