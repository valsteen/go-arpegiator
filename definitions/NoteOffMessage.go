package midiDefinitions

import (
	"fmt"
	"go-arpegiator/services/set"
)

type NoteOffMessage rawMidiMessage

func (message NoteOffMessage) GetChannel() byte {
	return (message[0]-128)%16 + 1
}

func (message NoteOffMessage) GetPitch() byte {
	return message[1]
}

func (message NoteOffMessage) GetVelocity() byte {
	return message[2]
}

func (message NoteOffMessage) Hash() set.Hash {
	return set.Hash([]byte{message.GetChannel(), message.GetPitch()})
}

func (message NoteOffMessage) String() string {
	return fmt.Sprintf("Note off: channel %d pitch %d velocity %d", message.GetChannel(),
		message.GetPitch(),
		message.GetVelocity())
}

func MakeNoteOffMessage(channel, pitch, velocity byte) NoteOffMessage {
	message := make([]byte, 3)
	message[0] = channel + 128
	message[1] = pitch
	message[2] = velocity
	return message
}
