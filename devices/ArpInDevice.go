package devices

import (
	"fmt"
	midiDefinitions "go-arpegiator/definitions"
	"go-arpegiator/services"
	"go-arpegiator/services/set"
)

type ArpSwitch struct {
	midiDefinitions.Note
}

type ArpInDevice struct {
	switches set.Set
}

func (d * ArpInDevice) Consume(notes set.Set) {
	newSwitches := NewSwitches(notes)
	added, removed := d.switches.Diff(newSwitches)
	fmt.Printf("added: %v removed: %v\n", added, removed)
	d.switches = newSwitches
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
		switches: make(set.Set),
	}
}

func NewSwitches(notes set.Set) set.Set {
	switches := make(set.Set)
	for _, element := range notes {
		note, ok := element.(midiDefinitions.Note)
		services.Must(ok)
		switches.Add(ArpSwitch{Note: note})
	}
	return switches
}