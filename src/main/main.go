package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/light", LightIndex)
	router.GET("/light/:id", Light)
	router.GET("/air", AirIndex)
	router.GET("/air/:id", Air)
	router.GET("/senpu", SenpuIndex)
	router.GET("/senpu/:id", Senpu)

	log.Fatal(http.ListenAndServe(":8089", router))
}

func runCommand(cmd, arg string) {
	if arg == "" {
		log.Fatal("arg error")
	} else {
		err := exec.Command(cmd, arg).Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func readSeq(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcmoe!")
}

func LightIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Light Welcmoe!")
}

func initCmd(w http.ResponseWriter, ps httprouter.Params) (code, cmd, arg string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	cd := ps.ByName("id")
	cm := "./cmd/remocon"
	ar := ""
	return cd, cm, ar
}

func Light(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//	fmt.Fprintf(w, "Light show: %s", ps.ByName("id"))
	code, cmd, arg := initCmd(w, ps)
	switch code {
	case "on":
		fmt.Fprintf(w, "Light show: on")
		arg = fmt.Sprintf("-d %s", readSeq("./cmd/light_all.txt"))
	case "off":
		fmt.Fprintf(w, "Light show: off")
		arg = fmt.Sprintf("-d %s", readSeq("./cmd/light_off.txt"))
	case "fav":
		fmt.Fprintf(w, "Light show: fav")
		arg = fmt.Sprintf("-d %s", readSeq("./cmd/light_fav.txt"))
	}
	runCommand(cmd, arg)
}

func AirIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Air Welcmoe!")
}

func Air(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//fmt.Fprintf(w, "Air Welcmoe!")
	code, cmd, arg := initCmd(w, ps)
	switch code {
	case "jositu":
		fmt.Fprintf(w, "Air show: on")
		arg = fmt.Sprintf("-d %s", readSeq("./cmd/air_jositu.txt"))
	case "reibo":
		fmt.Fprintf(w, "Air show: off")
		arg = fmt.Sprintf("-d %s", readSeq("./cmd/air_reibo.txt"))
	case "off":
		fmt.Fprintf(w, "Air show: fav")
		arg = fmt.Sprintf("-d %s", readSeq("./cmd/air_off.txt"))
	case "timer-on":
		fmt.Fprintf(w, "Air show: timer on")
		arg = fmt.Sprintf("-d %s", readSeq("./cmd/air_timer_on.txt"))
	case "timer-off":
		fmt.Fprintf(w, "Air show: timer off")
		arg = fmt.Sprintf("-d %s", readSeq("./cmd/air_timer_off.txt"))
	}
	runCommand(cmd, arg)
}

func SenpuIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Senpu Welcmoe!")
}

func Senpu(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//fmt.Fprintf(w, "Senpu show: %s", ps.ByName("id"))
	code, cmd, arg := initCmd(w, ps)
	switch code {
	case "on":
		fmt.Fprintf(w, "Senpu show: on")
		arg = fmt.Sprintf("-d %s", readSeq("./cmd/senpuuki_on-off.txt"))
	case "off":
		fmt.Fprintf(w, "Senpu show: off")
		arg = fmt.Sprintf("-d %s", readSeq("./cmd/senpuuki_on-off.txt"))
	case "timer":
		fmt.Fprintf(w, "Senpu show: timer")
		arg = fmt.Sprintf("-d %s", readSeq("./cmd/senpuuki_timer.txt"))
		runCommand(cmd, arg)
		time.Sleep(2 * time.Second)
	}
	runCommand(cmd, arg)
}
