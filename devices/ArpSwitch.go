package devices

import (
	"fmt"
	"go-arpegiator/definitions"
)

type ArpSwitch struct {
	midiDefinitions.Note
}

func (a ArpSwitch) GetIndex() byte {
	return a.GetPitch() % 12
}

func (a ArpSwitch) GetOctave() int8 {
	// C4 is considered octave 0
	return (int8(a.GetPitch()) - 60) / 12
}

func (a ArpSwitch) String() string {
	return fmt.Sprintf("switch = (%d %v %d)", a.GetChannel(), a.GetOctave(), a.GetIndex())
}
