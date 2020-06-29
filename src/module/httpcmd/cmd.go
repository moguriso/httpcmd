// +build amd64

package httpcmd

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func initCmd() (cmd, arg string) {
	cm := "./cmd/remocon"
	ar := ""
	return cm, ar
}

var lock sync.Mutex

func CtrlDehumidifier() error {
	lock.Lock()
	script_path := os.Getenv("SWITCH_BOT_SCRIPT_PATH")
	ble_addr := os.Getenv("SWITCH_BOT_DEVICE_ADDRESS")
	out, err := exec.Command("/usr/bin/python3",
		script_path, "-d", ble_addr, "press").Output()
	lock.Unlock()
	if err != nil {
		log.Println("CtrlDehumidifier: ")
		log.Println(err)
	} else {
		s := string(out)
		if strings.Index(s, "RuntimeError") != -1 {
			log.Println("CtrlDehumidifier: command execution failed")
			return errors.New("CtrlDehumidifier: command execution failed")
		}
	}
	return nil
}

func RunCommand(cmd, arg string) {
	log.Println("call cmd = ", cmd, " arg = ", arg)
	if arg == "" {
		log.Fatal("arg error")
	} else {
		out, err := exec.Command(cmd, arg).Output()
		if err != nil {
			log.Println(out)
			log.Panic(err)
		} else {
			log.Println(out)
		}
	}
}

func ReadSeq(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
