package comu

import (
	"testing"
	"time"
)

func TestCli(t *testing.T) {
	t.Run("Run SetBannerFunction()", func(t *testing.T) {
		clock := NewClock(100.0)
		clock.TimeSignature = []int{0, 0}
		clock.NewBPM(120.0)
		time.Sleep(4 * time.Second)
	})
}
