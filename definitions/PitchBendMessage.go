package midiDefinitions

import (
	"fmt"
	"math"
)

type PitchBendMessage rawMidiMessage

func (message PitchBendMessage) GetPitchBend() byte {
	return message[1]
}

func (message PitchBendMessage) GetChannel() byte {
	return message[0] - PITCHBEND
}

func (message PitchBendMessage) String() string {
	return fmt.Sprintf("pitchbend: channel %d value %d", message.GetChannel(), message.GetPitchBend())
}

func NewPitchBendMessage(channel byte, semitones float64) PitchBendMessage {
	// scale up 96 values to 128, then shift 7 bits
	pitchBendValue := int(math.Round(semitones * 128 * 128 / 96 / 1000))
	return []byte{
		PITCHBEND + channel,
		byte(pitchBendValue & 0x7F),
		byte(pitchBendValue >> 7),
	}
}
