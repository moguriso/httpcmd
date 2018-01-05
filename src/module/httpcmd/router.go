// +build arm arm64

package httpcmd

import (
	"fmt"
	"log"
	"net/http"

	"module/gpio"

	"github.com/julienschmidt/httprouter"
)

func Listen(port string) {

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/led", LedIndex)
	router.GET("/led/:id", Led)

	log.Fatal(http.ListenAndServe(port, router))
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
	return code, "", ""
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcmoe!")
}

func LedIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Led Welcmoe!")
}

func Led(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//	fmt.Fprintf(w, "Led show: %s", ps.ByName("id"))
	code, _, _ := preInit(w, ps)
	switch code {
	case "on":
		fmt.Fprintf(w, "Led show: on")
		gpio.LedOn()
	case "off":
		fmt.Fprintf(w, "Led show: off")
		gpio.LedOff()
	case "blue":
		fmt.Fprintf(w, "Led show: blue")
		gpio.LedBlue()
	case "green":
		fmt.Fprintf(w, "Led show: green")
		gpio.LedGreen()
	case "lime":
		fmt.Fprintf(w, "Led show: lime")
		gpio.LedLime()
	case "perple":
		fmt.Fprintf(w, "Led show: perple")
		gpio.LedPerple()
	case "pink":
		fmt.Fprintf(w, "Led show: pink")
		gpio.LedPink()
	case "red":
		fmt.Fprintf(w, "Led show: red")
		gpio.LedRed()
	case "white":
		fmt.Fprintf(w, "Led show: white")
		gpio.LedWhite()
	case "yellow":
		fmt.Fprintf(w, "Led show: yellow")
		gpio.LedYellow()
	}
}
