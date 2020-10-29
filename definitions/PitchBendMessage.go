package midiDefinitions

import "fmt"

type PitchBendMessage rawMidiMessage

func (message PitchBendMessage) GetValue() byte {
	return message[1]
}

func (message PitchBendMessage) GetChannel() byte {
	return message[0] - PITCHBEND
}

func (message PitchBendMessage) String() string {
	return fmt.Sprintf("pitchbend: channel %d value %d", message.GetChannel(), message.GetValue())
}
