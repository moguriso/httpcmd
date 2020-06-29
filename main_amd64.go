package main

import (
	"flag"
	"module/config"
	"module/httpcmd"
	"os"
)

var configFile = flag.String("conf", "", "config file name (fullpath)")

func main() {

	flag.Parse()

	if _, err := os.Stat(*configFile); err != nil {
		*configFile = "config.json"
	}
	conf, _ := config.NewParams()
	conf.LoadParam(*configFile)

	//sd, _ := sensortag.NewData(conf.SensorUrl, conf.SensorReadInterval)
	//go sd.ReadButtonThread()

	httpcmd.Listen(":8089")
}
