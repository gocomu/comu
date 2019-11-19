package comuio

import (
	"fmt"
	"strconv"

	"github.com/hypebeast/go-osc/osc"
)

type OSCio struct {
	Server *osc.Server
	client *osc.Client
}

func NewOSCio(serverPort, clientAddr, clientPort string) *OSCio {
	oscio := &OSCio{}
	if serverPort != "" {
		addr := "0.0.0.0:" + serverPort
		server := &osc.Server{Addr: addr}
		oscio.Server = server
		go oscio.Server.ListenAndServe()
	}

	if clientAddr != "" && clientPort != "" {
		oscio.client = osc.NewClient(clientAddr, stringToInt(clientPort))
	}

	return oscio
}

func (o *OSCio) Message(messageAddress string, data ...interface{}) {
	msg := osc.NewMessage(messageAddress)
	for _, value := range data {
		msg.Append(value)
	}
	o.client.Send(msg)
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	return i
}
