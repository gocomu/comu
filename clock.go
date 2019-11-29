package comu

import (
	"sync"
	"time"
)

// TempoClock struct holds all tempo related information
type TempoClock struct {
	// BPM provides an altervative way to poll current BPM without using channels
	BPM float64

	// BPMchange is the main way to alter the tempo on the fly
	// it is a buffered channel
	BPMchange chan float64

	// Beat is comu's universal clock ticker
	Beat *time.Ticker

	// BeatCounter counts how many beats have been played since the program started
	BeatCounter int64

	// TimeSignature is optional and holds the time signature of the piece
	TimeSignature []int

	// BarCounter to work needs a non-empty TimeSignature
	BarCounter int64

	// TimeStarted holds time.Now() from when the piece started
	TimeStarted time.Time

	// MStilNetxtBeat is the time until next beat in milliseconds
	MStilNetxtBeat float64

	mu *sync.Mutex
}

// NewClock returns a new TempoClock struct
func NewClock(initBPM float64) *TempoClock {
	tempo := &TempoClock{
		BPM:       (60000 / initBPM) * 0.5,
		BPMchange: make(chan float64, 1),
		Beat:      time.NewTicker(time.Duration((60000/initBPM)*0.5) * time.Millisecond),
		// default time signature is 4/4
		TimeSignature:  []int{4, 4},
		MStilNetxtBeat: (60000 / initBPM) * 0.5,
		mu:             &sync.Mutex{},
	}
	// start the clock
	tempo.clock()

	return tempo
}

// NewBPM provides a function based way of changing the tempo
func (t *TempoClock) NewBPM(newTempo float64) {
	t.BPMchange <- newTempo
}

func (t *TempoClock) clock() {
	t.TimeStarted = time.Now()
	go func() {
		for {
			// on each new beat update BeatCounter and BarCounter
			<-t.Beat.C
			t.BeatCounter++
			t.barCounterUpdate()
		}
	}()

	go func() {
		for {
			// when BPMchange channels receives a value a new ticker is set
			// along with new BPM and MStilNetxtBeat
			newTempo := <-t.BPMchange
			beatInMS := (60000 / newTempo) * 0.5
			t.Beat = time.NewTicker(time.Duration(beatInMS) * time.Millisecond)
			t.BPM = newTempo
			t.MStilNetxtBeat = beatInMS
		}
	}()

}

func (t *TempoClock) barCounterUpdate() {
	signature := float64(t.TimeSignature[0])
	// if user has not set TimeSignature default value to 1
	if signature == 0 {
		signature = 1.0
	}
	counter := float64(t.BeatCounter)
	// prevent counter from multiplying with 0
	if counter != 0 {
		currentBar := counter / signature
		if isIntegral(currentBar) == true {
			t.BarCounter++
		}
	}
}

// check if a value is int
func isIntegral(val float64) bool {
	return val == float64(int(val))
}
