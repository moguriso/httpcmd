package httpcmd

import (
	"io/ioutil"
	"log"
	"os/exec"
)

func initCmd() (cmd, arg string) {
	cm := "./cmd/remocon"
	ar := ""
	return cm, ar
}

func runCommand(cmd, arg string) {
	log.Println("call cmd = ", cmd)
	if arg == "" {
		log.Fatal("arg error")
	} else {
		out, err := exec.Command(cmd, arg).Output()
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println(out)
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
