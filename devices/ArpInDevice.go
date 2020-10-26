package devices

import (
	"fmt"
	midiDefinitions "go-arpegiator/definitions"
)

type ArpSwitch struct {
	midiDefinitions.Note
}

type ArpInDevice struct {
	switches ArpSwitchSet
}

func (d *ArpInDevice) ConsumeNoteSet(notes NoteSet) {
	newSwitches := newArpSwitchSet(notes)
	added, removed := d.switches.Compare(newSwitches)
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
		switches: make(ArpSwitchSet),
	}
}
