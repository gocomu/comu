package cio

import (
	"testing"
	"time"

	"github.com/hypebeast/go-osc/osc"
)

func TestOscIO(t *testing.T) {
	var oscio *OscIO

	t.Run("", func(t *testing.T) {
		oscio = NewOscIO("8765", "localhost", "8765")
		time.Sleep(4 * time.Second)
	})

	t.Run("", func(t *testing.T) {
		oscio.Server.Handle("/message/address", func(msg *osc.Message) {
			osc.PrintMessage(msg)
		})
	})

	t.Run("", func(t *testing.T) {
		oscio.Message("/message/address", int32(666), true, "Hello")
		time.Sleep(1 * time.Second)
	})
}
