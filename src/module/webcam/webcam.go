package webcam

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/lazywei/go-opencv/opencv"
)

var gCap *opencv.Capture

func checkAndMoveFile(f string) {
	if _, err := os.Stat(f); err == nil {
		if err := os.Chmod(f, 0644); err != nil {
			fmt.Println(err)
		}
	}
}

var cameraIndex int

func LoadParam(fileName string) {
	isSet := false
	var config map[string]interface{}
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println(fileName, " read error")
		return
	}

	err = json.Unmarshal(data, &config)
	var ftmp float64
	ftmp, isSet = config["camera_index"].(float64)
	if !isSet {
		cameraIndex = 0
	} else {
		cameraIndex = int(ftmp)
	}

	log.Println("parameter: ", cameraIndex)
}

func Init() {
	gCap = opencv.NewCameraCapture(cameraIndex)
	if gCap == nil {
		panic("can not open camera")
	}
}

func Deinit() {
	gCap.Release()
}

func Snap(f string) {
	count := 0
	for {
		if gCap.GrabFrame() {
			imgCV := gCap.RetrieveFrame(10)
			if imgCV != nil {
				const layout = "2006-01-02 15:04:05"
				t := time.Now()
				imgCV = opencv.Resize(imgCV, 320, 240, 0)
				col := opencv.NewScalar(255, 255, 255, 255)
				font := opencv.InitFont(opencv.CV_FONT_HERSHEY_PLAIN, 1.2, 1.2, 0, 1, 8)
				point := opencv.NewCvPoint(5.0, imgCV.Height()-5.0)
				font.PutText(imgCV, t.Format(layout), point.ToPoint(), col)
				opencv.SaveImage(f, imgCV, nil)
				checkAndMoveFile(f)
				log.Println("snapshot Image")
				if count++; count > 3 {
					break
				}
			} else {
				log.Println("Image is nil")
			}
		}
	}
}
