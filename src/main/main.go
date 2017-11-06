package main

import (
	"flag"
	"module/httpcmd"
	"module/webcam"
	"os"
)

var configFile = flag.String("conf", "", "config file name (fullpath)")

func main() {

	flag.Parse()

	if _, err := os.Stat(*configFile); err != nil {
		*configFile = "config.json"
	}

	webcam.LoadParam(*configFile)
	httpcmd.Listen(":8089")
}
