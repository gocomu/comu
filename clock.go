package comu

import (
	"time"
)

// type Clock interface {
// 	NewBPM(newTempo float64)
// }

// Tempo struct holds all tempo related information
type TempoClock struct {
	BPM            float64
	BPMchange      chan float64
	Beat           *time.Ticker
	BeatCounter    int64
	TimeSignature  []int
	BarCounter     int64
	TimeStarted    time.Time
	MStilNetxtBeat float64
}

// NewTempo returns a new Tempo struct
func NewClock(initBPM float64) *TempoClock {
	beatInMS := 60000 / initBPM
	timer := time.NewTicker(time.Duration(beatInMS) * time.Millisecond)
	tempoChangeCh := make(chan float64, 1)
	tempo := &TempoClock{
		BPM:            initBPM,
		BPMchange:      tempoChangeCh,
		Beat:           timer,
		TimeSignature:  []int{4, 4},
		MStilNetxtBeat: beatInMS,
	}
	go tempo.clock()

	return tempo
}

func (t *TempoClock) NewBPM(newTempo float64) {
	t.BPMchange <- newTempo
}

// func (t *Clock) TimeCalc() bool {
// 	return t.TimeStarted.After(time.Now())
// }

func (t *TempoClock) clock() {
	t.TimeStarted = time.Now()
	for {
		select {
		case <-t.Beat.C:
			t.BeatCounter++
			go t.barCounterUpdate()
		case newTempo := <-t.BPMchange:
			beatInMS := 60000 / newTempo
			newtimer := time.NewTicker(time.Duration(beatInMS) * time.Millisecond)
			t.Beat = newtimer
			t.BPM = newTempo
			t.MStilNetxtBeat = beatInMS
		}
	}
}

func (t *TempoClock) barCounterUpdate() {
	signature := float64(t.TimeSignature[0])
	// if user has not set TimeSignature default value to 1
	if signature == 0 {
		signature = 1.0
	}
	counter := float64(t.BeatCounter)
	if counter != 0 {
		currentBar := counter / signature
		if isIntegral(currentBar) == true {
			t.BarCounter++
		}
	}
}

func isIntegral(val float64) bool {
	return val == float64(int(val))
}
