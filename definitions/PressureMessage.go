package midiDefinitions

import "fmt"

type PressureMessage rawMidiMessage

func (message PressureMessage) GetValue() byte {
	return message[2]
}

func (message PressureMessage) GetChannel() byte {
	return message[0] - PRESSURE
}

func (message PressureMessage) String() string {
	return fmt.Sprintf("pressure: channel %d value %d", message.GetChannel(), message.GetValue())
}
