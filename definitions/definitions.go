package midiDefinitions

import (
	"go-arpegiator/services/set"
)

const (
	NOTEOFF   = 0x80
	NOTEON    = 0x90
	CC        = 0xB0
	PRESSURE  = 0xD0
	PITCHBEND = 0xE0
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

type Cc interface {
	ChannelMessage
	GetCc() byte
	GetValue() byte
}

type Pressure interface {
	ChannelMessage
	GetChannel() byte
	GetValue() byte
}

type PitchBend interface {
	ChannelMessage
	GetChannel() byte
	GetValue() byte
}

func AsMidiMessage(bytes []byte) MidiMessage {
	if bytes[0] >= NOTEOFF && bytes[0] < NOTEOFF+0x10 {
		return NoteOffMessage(bytes)
	} else if bytes[0] >= NOTEON && bytes[0] < NOTEON+0x10 {
		return NoteOnMessage(bytes)
	} else if bytes[0] >= CC && bytes[0] < CC+0x10 {
		return CCMessage(bytes)
	} else if bytes[0] >= PITCHBEND && bytes[0] < PITCHBEND+0x10 {
		return PitchBendMessage(bytes)
	} else if bytes[0] >= PRESSURE && bytes[0] < PRESSURE+0x10 {
		return PressureMessage(bytes)
	}
	return nil
}
