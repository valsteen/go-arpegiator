package midiDefinitions

import (
	"fmt"
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/rtmididrv"
	"go-arpegiator/services"
)

type PortPair struct {
	*rtmididrv.Driver
	midi.In
	midi.Out
}

func NewPortPair(name string) *PortPair {
	drv, err := rtmididrv.New()

	services.MustNot(err)

	in, err := drv.OpenVirtualIn(name)
	services.MustNot(err)
	out, err := drv.OpenVirtualOut(name)
	services.MustNot(err)

	return &PortPair{drv, in, out}
}

func (pair *PortPair) Close() {
	_ = pair.Driver.Close()
}

func (pair PortPair) MidiPassThrough(data []byte, deltaMicroseconds int64) {
	midiMessage := AsMidiMessage(data)
	fmt.Printf("just in: %v %s\n", data, midiMessage)
	if channelMessage, ok := midiMessage.(ChannelMessage); ok {
		fmt.Println(channelMessage.GetChannel())
	}
	_, err := pair.Write(data)
	services.MustNot(err)
}
