package midiDefinitions

import "fmt"

type MidiMessage interface{} // placeholder for now, can be cast to anything in this form

type rawMidiMessage []byte

type noteMessage rawMidiMessage
type ccMessage rawMidiMessage
type NoteHash string

type ChannelMessage interface {
	MidiMessage
	GetChannel() byte
}

type Note interface {
	ChannelMessage
	GetPitch() byte
	GetVelocity() byte
	GetNoteHash() NoteHash
}

type NoteMessage interface {
	Note
	IsNoteOn() bool
}

type CC interface {
	ChannelMessage
	GetCC() byte
	GetValue() byte
}

func (message noteMessage) GetChannel() byte {
	return (message[0]-144)%16 + 1
}

func (message noteMessage) GetPitch() byte {
	return message[1]
}

func (message noteMessage) GetVelocity() byte {
	return message[2]
}

func (message noteMessage) IsNoteOn() bool {
	return message[0] >= 144 && message[0] < 160
}

func (message noteMessage) GetNoteHash() NoteHash {
	return NoteHash([]byte{message.GetChannel(), message.GetPitch()})
}

func (message ccMessage) GetCC() byte {
	return message[1]
}

func (message ccMessage) GetValue() byte {
	return message[2]
}

func (message ccMessage) GetChannel() byte {
	return message[0] - 176 + 1
}

func AsMidiMessage(bytes []byte) MidiMessage {
	if bytes[0] >= 128 && bytes[0] < 160 {
		return noteMessage(bytes)
	} else if bytes[0] >= 176 && bytes[0] < 192 {
		return ccMessage(bytes)
	}
	return nil
}

func (message noteMessage) String() string {
	var onOff string
	if message.IsNoteOn() {
		onOff = "on"
	} else {
		onOff = "off"
	}
	return fmt.Sprintf("Note %s: channel %d pitch %d velocity %d", onOff, message.GetChannel(),
		message.GetPitch(),
		message.GetVelocity())
}

func (message ccMessage) String() string {
	return fmt.Sprintf("cc: channel %d cc %d value %d", message.GetChannel(), message.GetCC(),
		message.GetValue())
}

func (noteHash NoteHash) String() string {
	return fmt.Sprintf("(%d %d)", noteHash[0], noteHash[1])
}