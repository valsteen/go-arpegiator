package runner

import (
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

func RunDevice(name string, consumer devices.NotesConsumer) DeviceRunner {
	pair := midiDefinitions.NewPortPair(name)
	deviceRunner := DeviceRunner{
		Device:   devices.NewDevice(pair.In),
		PortPair: pair,
	}
	deviceRunner.AddConsumer(consumer)
	return deviceRunner
}

func RunArpInDevice(name string) DeviceRunner {
	arpInDevice := devices.NewArpInDevice()
	deviceRunner := RunDevice(name, arpInDevice.Consume)
	return deviceRunner
}
