package cio

import (
	"strconv"

	"github.com/hypebeast/go-osc/osc"
)

//TODO: clock auto receivers
//TODO: bundle messages
//TODO: clock auto messages

// OscIO holds OSC server/client connections
// OscIO is a thin wrapper around 'github.com/hypebeast/go-osc'
// providing an even more simplified & integrated osc usage for gocomu users
type OscIO struct {
	Server *osc.Server
	client *osc.Client
}

// NewOSCio returns a new OSC connection
// note: if serverPort arg is left empty "" the only a client will start
// similarly if clientAddr & clientPort are left blank ""
// only a server connection will start
// when starting a server or client comu automatically exposes TempoClock
// to send/receive tempo related information
// for more details about the API check comu's documentation on github
func NewOscIO(serverPort, clientAddr, clientPort string) *OscIO {
	oscio := &OscIO{}
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

// Message sends an OSC message to 'message/address' with given arguments
func (o *OscIO) Message(messageAddress string, data ...interface{}) {
	msg := osc.NewMessage(messageAddress)
	for _, value := range data {
		msg.Append(value)
	}
	o.client.Send(msg)
}

func stringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
