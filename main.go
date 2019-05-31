package main

import (
	"bytes"
	"flag"
	"log"
	"math/rand"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/subpathdev/CpuTempMqttClient/kubeClient"
)

var ipAddress, deviceID, user, password string
var simulate bool

func init() {
	//	flag := flag.NewFlagSet("Usage", flag.ExitOnError)
	flag.StringVar(&ipAddress, "ipAddress", "tcp://127.0.0.1:1884", "IPAddress with used protocol and port")
	flag.StringVar(&deviceID, "deviceID", "43098512438508132096394-a41fcb", "The unique ID of this device (created in the cloud)")
	flag.StringVar(&user, "user", "", "is the user, which should use by the mqtt connection")
	flag.StringVar(&password, "password", "", "is the password of the MQTT User")
	flag.BoolVar(&simulate, "simulate", false, "if you use this flag the input data will be simulated by random numbers and no sensor will be requested")
}

func main() {
	flag.Parse()
	kubeClient.Init(ipAddress, deviceID, user, password)
	for {
		var message string
		var out bytes.Buffer
		var core0 = false

		if simulate {
			message = strconv.Itoa(rand.Intn(200))
		} else {
			cmd := exec.Command("/usr/bin/sensors -Au")
			cmd.Stdout = &out
			cmd.Stderr = &out
			err := cmd.Start()
			if err != nil {
				log.Println("Error in command execution. Error: ", err)
			}
			err = cmd.Wait()
			if err != nil {
				log.Println("Error by waiting on command execution. Error: ", err)
			}
			str := strings.Split(out.String(), "\n")

			for _, element := range str {
				if core0 == true {
					if strings.Contains(element, "temp2_input") {
						val := strings.Split(element, ": ")
						message = val[1]
						core0 = false
					}
				}

				if strings.Contains(element, "Core 0:") {
					core0 = true
				}

			}
		}

		kubeClient.Update(message)
		time.Sleep(10 * time.Second)
	}
}
