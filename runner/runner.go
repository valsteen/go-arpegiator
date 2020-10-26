package runner

import (
	"fmt"
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/rtmididrv"
	"go-arpegiator/definitions"
	"go-arpegiator/devices"
	s "go-arpegiator/services"
)

type ArpegiatorRunner struct {
	*rtmididrv.Driver
	midi.In
	arpInPortPair *midiDefinitions.PortPair
}

func (a ArpegiatorRunner) Close() {
	_ = a.In.Close()
	a.arpInPortPair.Close()
	if a.Driver != nil {
		_ = a.Driver.Close()
	}
}

func RunArpegiator(notesInName, arpName string) ArpegiatorRunner {
	driver, err := rtmididrv.New()
	s.MustNot(err)

	in, err := driver.OpenVirtualIn(notesInName)
	s.MustNot(err)

	arpegiatorRunner := ArpegiatorRunner{
		Driver:        driver,
		In:            in,
		arpInPortPair: midiDefinitions.NewPortPair(arpName, driver),
	}

	noteInDevice := devices.NewNoteDevice(arpegiatorRunner.In)
	arpInDevice := devices.NewNoteDevice(arpegiatorRunner.arpInPortPair.In)
	arpegiator := devices.NewArpegiator(noteInDevice, arpInDevice)

	arpegiator.AddMessageConsumer(func(message midiDefinitions.ChannelMessage) {
		fmt.Println("Arp out message", message)
	})

	return arpegiatorRunner
}
