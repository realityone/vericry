package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/spf13/pflag"

	"github.com/realityone/vericry/pkg/config"
)

func main() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	config.Logging()
	glog.V(2).Info("Starting vericry server...")
}
