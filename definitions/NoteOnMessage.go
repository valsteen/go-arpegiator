package midiDefinitions

import (
	"fmt"
	"go-arpegiator/services"
	"go-arpegiator/services/set"
)

type NoteOnMessage rawMidiMessage

func (n NoteOnMessage) Less(element set.Element) bool {
	message2, ok := element.(Note)
	services.Must(ok)
	return n.GetPitch() < message2.GetPitch() ||
		(n.GetPitch() == message2.GetPitch() && n.GetChannel() < message2.GetChannel())
}

func (n NoteOnMessage) GetChannel() byte {
	return (n[0] - NOTEON) % 0x10
}

func (n NoteOnMessage) GetPitch() byte {
	return n[1]
}

func (n NoteOnMessage) GetVelocity() byte {
	return n[2]
}

func (n NoteOnMessage) String() string {
	return fmt.Sprintf("Note on: channel %d pitch %d velocity %d", n.GetChannel(),
		n.GetPitch(),
		n.GetVelocity())
}

func NewNoteOnMessage(channel, pitch, velocity byte) []byte {
	return []byte{
		channel + NOTEON,
		pitch,
		velocity,
	}
}
