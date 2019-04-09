package main

import (
	"flag"
	"math/rand"
	"strconv"

	"./kubeClient"
)

var ipAddress, deviceID, user, password string

func init() {
//	flag := flag.NewFlagSet("Usage", flag.ExitOnError)
	flag.StringVar(&ipAddress,"ipAddress" , "tcp://127.0.0.1:1884" , "IPAddress with used protocol and port")
	flag.StringVar(&deviceID, "deviceID", "43098512438508132096394-a41fcb", "The unique ID of this device (created in the cloud)")
	flag.StringVar(&user, "user", "", "is the user, which should use by the mqtt connection")
	flag.StringVar(&password, "password", "", "is the password of the MQTT User")
}

func main() {
	flag.Parse()
	kubeClient.Init(ipAddress, deviceID, user, password)
	for {
		var message string
		message = strconv.Itoa(rand.Int())
		kubeClient.Update(message)
	}
}
