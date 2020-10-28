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
	return (n[0]-144)%16 + 1
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

func NewNoteOnMessage(channel, pitch, velocity byte) NoteOnMessage {
	message := make([]byte, 3)
	message[0] = channel + 144
	message[1] = pitch
	message[2] = velocity
	return message
}

func NewDeadNoteMessage(channel, pitch byte) NoteOnMessage {
	message := make([]byte, 3)
	message[0] = channel + 144
	message[1] = pitch
	message[2] = 0
	return message
}

func (n NoteOnMessage) IsDeadNote() bool {
	return n.GetVelocity() == 0
}
