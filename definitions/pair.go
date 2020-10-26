package midiDefinitions

import (
	"fmt"
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/rtmididrv"
	"go-arpegiator/services"
)

type PortPair struct {
	midi.In
	midi.Out
}

func NewPortPair(name string, driver *rtmididrv.Driver) *PortPair {
	var err error
	portPair := PortPair{}
	portPair.In, err = driver.OpenVirtualIn(name)
	services.MustNot(err)
	portPair.Out, err = driver.OpenVirtualOut(name)
	services.MustNot(err)
	return &portPair
}

func (pair *PortPair) Close() {
	_ = pair.In.Close()
	_ = pair.Out.Close()
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
