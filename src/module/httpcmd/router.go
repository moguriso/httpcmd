package httpcmd

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"module/webcam"

	"github.com/julienschmidt/httprouter"
)

func Listen(port string) {

	webcam.Init()

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/light", LightIndex)
	router.GET("/light/:id", Light)
	router.GET("/air", AirIndex)
	router.GET("/air/:id", Air)
	router.GET("/senpu", SenpuIndex)
	router.GET("/senpu/:id", Senpu)
	router.GET("/webcam", WebCamIndex)
	router.GET("/webcam/:id", WebCam)

	log.Fatal(http.ListenAndServe(port, router))

	webcam.Deinit()
}

func modHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func getCode(ps httprouter.Params) string {
	return ps.ByName("id")
}

func preInit(w http.ResponseWriter, ps httprouter.Params) (string, string, string) {
	modHeader(w)
	code := getCode(ps)
	cmd, arg := initCmd()
	return code, cmd, arg
}

func preCameraInit(w http.ResponseWriter, ps httprouter.Params) (string, string, string) {
	modHeader(w)
	code := getCode(ps)
	return code, "", ""
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcmoe!")
}

func LightIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Light Welcmoe!")
}

func Light(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//	fmt.Fprintf(w, "Light show: %s", ps.ByName("id"))
	code, cmd, arg := preInit(w, ps)
	switch code {
	case "on":
		fmt.Fprintf(w, "Light show: on")
		arg = fmt.Sprintf("-d %s", ReadSeq("./cmd/light_all.txt"))
	case "off":
		fmt.Fprintf(w, "Light show: off")
		arg = fmt.Sprintf("-d %s", ReadSeq("./cmd/light_off.txt"))
	case "fav":
		fmt.Fprintf(w, "Light show: fav")
		arg = fmt.Sprintf("-d %s", ReadSeq("./cmd/light_fav.txt"))
	}
	RunCommand(cmd, arg)
}

func AirIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Air Welcmoe!")
}

func Air(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//fmt.Fprintf(w, "Air Welcmoe!")
	code, cmd, arg := preInit(w, ps)
	switch code {
	case "jositu":
		fmt.Fprintf(w, "Air show: jositu on")
		arg = fmt.Sprintf("-d %s", ReadSeq("./cmd/air_jositu.txt"))
	case "reibo":
		fmt.Fprintf(w, "Air show: reibo on")
		arg = fmt.Sprintf("-d %s", ReadSeq("./cmd/air_reibo.txt"))
	case "warm":
		fmt.Fprintf(w, "Air show: warm on")
		arg = fmt.Sprintf("-d %s", ReadSeq("./cmd/air_warm.txt"))
	case "off":
		fmt.Fprintf(w, "Air show: off")
		arg = fmt.Sprintf("-d %s", ReadSeq("./cmd/air_off.txt"))
	case "timer-on":
		fmt.Fprintf(w, "Air show: timer on")
		arg = fmt.Sprintf("-d %s", ReadSeq("./cmd/air_timer_on.txt"))
	case "timer-off":
		fmt.Fprintf(w, "Air show: timer off")
		arg = fmt.Sprintf("-d %s", ReadSeq("./cmd/air_timer_off.txt"))
	}
	RunCommand(cmd, arg)
}

func SenpuIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Senpu Welcmoe!")
}

func Senpu(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//fmt.Fprintf(w, "Senpu show: %s", ps.ByName("id"))
	code, cmd, arg := preInit(w, ps)
	switch code {
	case "on":
		fmt.Fprintf(w, "Senpu show: on")
		arg = fmt.Sprintf("-d %s", ReadSeq("./cmd/senpuuki_on-off.txt"))
	case "off":
		fmt.Fprintf(w, "Senpu show: off")
		arg = fmt.Sprintf("-d %s", ReadSeq("./cmd/senpuuki_on-off.txt"))
	case "timer":
		fmt.Fprintf(w, "Senpu show: timer")
		arg = fmt.Sprintf("-d %s", ReadSeq("./cmd/senpuuki_timer.txt"))
		RunCommand(cmd, arg)
		time.Sleep(500 * time.Millisecond)
	}
	RunCommand(cmd, arg)
}

func WebCamIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Webcam Welcome!")
}

func WebCam(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	code, _, _ := preCameraInit(w, ps)
	switch code {
	case "snap":
		log.Println("Webcam snap: on")
		fmt.Fprintf(w, "Webcam snap: on")
		webcam.Snap("lastsnap.jpg")
	}
}
