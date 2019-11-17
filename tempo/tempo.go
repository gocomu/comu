package tempo

import (
	"fmt"
	"time"
)

// Tempo struct holds all tempo related information
type Tempo struct {
	BPM        float64
	BPMchannel chan bool
	BPMchange  chan float64
	ticker     *time.Ticker
}

// NewTempo returns a new Tempo struct
func NewTempo(initBPM float64) *Tempo {
	tempo := &Tempo{
		BPM:        initBPM,
		BPMchannel: make(chan bool),
		BPMchange:  make(chan float64),
	}
	tempo.clock()

	return tempo
}

// func (t *Tempo) BPMchange(newBPM float64) {
// 	t.BPM = newBPM
// 	t.bPMchange <- newBPM
// }

func (t *Tempo) clock() {
	tick := 60000 / t.BPM
	t.ticker = time.NewTicker(time.Duration(tick) * time.Millisecond)

	go func() {
		for {
			select {
			case <-t.ticker.C:
				t.BPMchannel <- true
			case newTempo := <-t.BPMchange:
				t.ticker.Stop()
				tick = 60000 / newTempo
				fmt.Println(tick)
				t.ticker = time.NewTicker(time.Duration(tick) * time.Millisecond)
			}

		}

	}()
}

// go func() {
// 	// for currentNote := 440.0; currentNote < 800.0; currentNote++ {
// 	// 	fmt.Printf("switching oscillator to %.2f Hz\n", currentNote)
// 	// 	osc.SetFreq(currentNote)
// 	// }
// 	currentNote := 400.0
// 	time.Sleep(2 * time.Second)

// 	ticker := time.NewTicker(10 * time.Millisecond)
// 	done := make(chan bool)

// 	for {
// 		select {
// 		case <-done:
// 			return
// 		case t := <-ticker.C:
// 			fmt.Printf("switching oscillator to %.2f Hz\n %s", currentNote, t)
// 			osc.SetFreq(currentNote)

// 			currentNote = currentNote + 1
// 		}
// 	}

// }()
