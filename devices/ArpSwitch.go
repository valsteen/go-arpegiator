package devices

import (
	"fmt"
	"go-arpegiator/definitions"
	"go-arpegiator/services"
	"go-arpegiator/services/set"
)

type ArpSwitch midiDefinitions.NoteOnMessage

func (a ArpSwitch) GetIndex() byte {
	return midiDefinitions.NoteOnMessage(a).GetPitch() % 12
}

func (a ArpSwitch) GetOctave() int8 {
	// C4 is considered octave 0
	return (int8(midiDefinitions.NoteOnMessage(a).GetPitch()) - 60) / 12
}

func (a ArpSwitch) GetChannel() byte {
	return midiDefinitions.NoteOnMessage(a).GetChannel()
}

func (a ArpSwitch) GetVelocity() byte {
	return midiDefinitions.NoteOnMessage(a).GetVelocity()
}

func (a ArpSwitch) String() string {
	return fmt.Sprintf("switch = (%d %v %d)", midiDefinitions.NoteOnMessage(a).GetChannel(), a.GetOctave(), a.GetIndex())
}

func (a ArpSwitch) Less(element set.Element) bool {
	a2, ok := element.(ArpSwitch)
	services.Must(ok)
	return a.GetIndex() < a2.GetIndex() || (a.GetIndex() == a2.GetIndex() && a.GetChannel() < a2.GetChannel())
}
