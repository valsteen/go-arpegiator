package midiDefinitions

import "fmt"

type CCMessage rawMidiMessage

func (message CCMessage) GetCc() byte {
	return message[1]
}

func (message CCMessage) GetValue() byte {
	return message[2]
}

func (message CCMessage) GetChannel() byte {
	return message[0] - CC
}

func (message CCMessage) String() string {
	return fmt.Sprintf("cc: channel %d cc %d value %d", message.GetChannel(), message.GetCc(),
		message.GetValue())
}
