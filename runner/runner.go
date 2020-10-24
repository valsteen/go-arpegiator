package runner

import (
	"fmt"
	"go-arpegiator/definitions"
	"go-arpegiator/devices"
	s "go-arpegiator/services"
)

func POC() {
	pair := midiDefinitions.NewPortPair("Arp")
	defer pair.Close()

	err := pair.SetListener(pair.MidiPassThrough)
	s.MustNot(err)
}

type DeviceRunner struct {
	*devices.Device
	*midiDefinitions.PortPair
}

func RunDevice() DeviceRunner {
	pair := midiDefinitions.NewPortPair("Arp")
	deviceRunner := DeviceRunner{
		Device:   devices.New(pair.In),
		PortPair: pair,
	}

	stateChangesChan := make(devices.StateChangeConsumer)
	deviceRunner.AddConsumer(stateChangesChan)

	go func() {
		for state := range stateChangesChan {
			fmt.Println(state)
		}
	}()
	return deviceRunner
}
