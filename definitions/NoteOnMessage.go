package midiDefinitions

import (
	"fmt"
	"go-arpegiator/services"
	"go-arpegiator/services/set"
)

type NoteOnMessage rawMidiMessage

func (message NoteOnMessage) Less(element set.Element) bool {
	message2, ok := element.(Note)
	services.Must(ok)
	return message.GetPitch() < message2.GetPitch() ||
		(message.GetPitch() == message2.GetPitch() && message.GetChannel() < message2.GetChannel())
}

func (message NoteOnMessage) GetChannel() byte {
	return (message[0]-144)%16 + 1
}

func (message NoteOnMessage) GetPitch() byte {
	return message[1]
}

func (message NoteOnMessage) GetVelocity() byte {
	return message[2]
}

func (message NoteOnMessage) String() string {
	return fmt.Sprintf("Note on: channel %d pitch %d velocity %d", message.GetChannel(),
		message.GetPitch(),
		message.GetVelocity())
}

func NewNoteOnMessage(channel, pitch, velocity byte) NoteOnMessage {
	message := make([]byte, 3)
	message[0] = channel + 144
	message[1] = pitch
	message[2] = velocity
	return message
}
