package devices

import (
	"fmt"
	midiDefinitions "go-arpegiator/definitions"
)

type ArpSwitch struct {
	midiDefinitions.Note
}

type switches map[midiDefinitions.NoteHash]ArpSwitch

type ArpInDevice struct {
	switches
}

func (d ArpInDevice) Consume(notes Notes) {
	d.switches = make(switches) // just reset for now
	for hash, note := range notes {
		d.switches[hash] = ArpSwitch{Note: note}
	}
	fmt.Println(d.switches)
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

func NewArpInDevice() *ArpInDevice {
	return &ArpInDevice{
		switches: make(switches),
	}
}
