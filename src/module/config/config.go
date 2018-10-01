package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Param struct {
	SensorUrl          string
	SensorReadInterval float64
}

func NewParams() (*Param, error) {
	p := &Param{
		SensorUrl:          "",
		SensorReadInterval: 0.0,
	}
	return p, nil
}

func (p *Param) LoadParam(fileName string) error {
	var config map[string]interface{}

	isSet := false
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println(fileName, " read error ", err)
		goto exit_func
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Println("json Unmarshal error ", err)
		goto exit_func
	}
	err = nil

	p.SensorUrl, isSet = config["sensor_url"].(string)
	if !isSet {
		p.SensorUrl = "http://localhost" /* tentative ... */
	}
	log.Println("sensor target URL: ", p.SensorUrl)

	p.SensorReadInterval, isSet = config["sensor_read_interval"].(float64)
	if !isSet {
		p.SensorReadInterval = 1000.0 /* default: 1sec */
	}
	log.Println("sensor read interval: ", p.SensorReadInterval)

exit_func:
	return err
}
