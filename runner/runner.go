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
	*devices.NoteDevice
	*midiDefinitions.PortPair
}

func RunNoteDevice(name string, consumer devices.NoteSetConsumer) DeviceRunner {
	pair := midiDefinitions.NewPortPair(name)
	deviceRunner := DeviceRunner{
		NoteDevice: devices.NewNoteDevice(pair.In),
		PortPair:   pair,
	}
	deviceRunner.AddNoteSetConsumer(consumer)
	return deviceRunner
}

func RunArpInDevice(name string) DeviceRunner {
	arpInDevice := devices.NewArpInDevice()
	deviceRunner := RunNoteDevice(name, arpInDevice.ConsumeNoteSet)
	return deviceRunner
}
