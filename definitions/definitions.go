package midiDefinitions

import (
	"go-arpegiator/services/set"
)

type MidiMessage interface{} // placeholder for now, can be cast to anything in this form
type rawMidiMessage []byte

type ChannelMessage interface {
	MidiMessage
	GetChannel() byte
}

type Note interface {
	ChannelMessage
	GetPitch() byte
	GetVelocity() byte
	set.Element
}

type CC interface {
	ChannelMessage
	GetCC() byte
	GetValue() byte
}

func AsMidiMessage(bytes []byte) MidiMessage {
	if bytes[0] >= 128 && bytes[0] < 144 {
		return NoteOffMessage(bytes)
	} else if bytes[0] >= 144 && bytes[0] < 160 {
		return NoteOnMessage(bytes)
	} else if bytes[0] >= 176 && bytes[0] < 192 {
		return CCMessage(bytes)
	}
	return nil
}
