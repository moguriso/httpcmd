package main

import (
	"flag"
	"module/config"
	"module/httpcmd"
	"module/sensortag"
	"module/webcam"
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

	webcam.SetCameraIndex(conf.CameraIndex)

	sd, _ := sensortag.NewData(conf.SensorUrl, conf.SensorReadInterval)
	go sd.ReadButtonThread()

	httpcmd.Listen(":8089")
}
