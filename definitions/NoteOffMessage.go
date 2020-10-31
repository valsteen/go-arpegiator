package midiDefinitions

import (
	"fmt"
	"go-arpegiator/services"
	"go-arpegiator/services/set"
)

type NoteOffMessage rawMidiMessage

func (message NoteOffMessage) GetChannel() byte {
	return (message[0] - NOTE_OFF) % 0x10
}

func (message NoteOffMessage) GetPitch() byte {
	return message[1]
}

func (message NoteOffMessage) GetVelocity() byte {
	return message[2]
}

func (message NoteOffMessage) String() string {
	return fmt.Sprintf("Note off: channel %d pitch %d velocity %d", message.GetChannel(),
		message.GetPitch(),
		message.GetVelocity())
}

func NewNoteOffMessage(channel, pitch, velocity byte) NoteOffMessage {
	return []byte{
		channel + NOTE_OFF,
		pitch,
		velocity,
	}
}

func (message NoteOffMessage) Less(element set.Element) bool {
	message2, ok := element.(Note)
	services.Must(ok)
	return message.GetPitch() < message2.GetPitch() ||
		(message.GetPitch() == message2.GetPitch() && message.GetChannel() < message2.GetChannel())
}
