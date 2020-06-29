package webcam

import (
	"fmt"
	"log"
	"os"

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

func Init() {
	for ii := 0; ii < 9; ii++ {
		gCap = opencv.NewCameraCapture(ii)
		if gCap != nil {
			log.Println("using camera ", ii)
			break
		} else {
			log.Println("camera ", ii, " not found")
		}
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
				imgCV = opencv.Resize(imgCV, 320, 240, 0)

				/* Web上の表示と同期しなくて見苦しいので消す {
					const layout = "2006-01-02 15:04:05"
					t := time.Now()
					col := opencv.NewScalar(255, 255, 255, 255)
					font := opencv.InitFont(opencv.CV_FONT_HERSHEY_PLAIN, 1.2, 1.2, 0, 1, 8)
					point := opencv.NewCvPoint(5.0, imgCV.Height()-5.0)
					font.PutText(imgCV, t.Format(layout), point.ToPoint(), col)
				} */

				opencv.SaveImage(f, imgCV, nil)
				checkAndMoveFile(f)
				//log.Println("snapshot Image")
				if count++; count > 3 {
					break
				}
			} else {
				log.Println("Image is nil")
			}
		}
	}
}
