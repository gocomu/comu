package pattern

import (
	"github.com/go-audio/generator"
	"github.com/gocomu/comu"
)

type Pattern struct {
	clock *comu.TempoClock
	osc   *generator.Osc
}

func NewPattern(clock *comu.TempoClock, osc *generator.Osc) *Pattern {
	return &Pattern{
		clock: clock,
		osc:   osc,
	}
}

func (p *Pattern) Four2TheFloor(pattern []int) {
	currentNote := 440.0
	for {
		select {
		case <-p.clock.Beat.C:
			p.osc.SetFreq(currentNote)
			currentNote = currentNote + 20.0
			if p.clock.BeatCounter == 8 {
				//fmt.Println("tempo change")
				p.clock.BPMchange <- 120.0
			}
			//log.Println(tempo.BarCounter)
			//tempo.BPMchange <- newTempo
			//newTempo = newTempo + 10.0
			//fmt.Println("received message ", newTempo)
		}
	}
}

//  start: First value of the sequence.
//  end:   The sequence is ended upon reaching the end value.
//  step:  step will be used as the increment between elements in the sequence.
//         step should be given as a positive number.
func stepSeq(start, end, step int) []int {
	if step <= 0 || end < start {
		return []int{}
	}
	s := make([]int, 0, 1+(end-start)/step)
	for start <= end {
		s = append(s, start)
		start += step
	}
	return s
}

func seq(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

// Returns a slice of elements with exact count.
// step will be used as the increment between elements in the sequence.
// step should be given as a positive, negative or zero number.
func ser(start, count, step int) []int {
	s := make([]int, count)
	for i := range s {
		s[i] = start
		start += step
	}
	return s
}
