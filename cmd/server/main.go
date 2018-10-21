package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/spf13/pflag"
)

func main() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	glog.V(2).Info("Starting vericry server...")
}
