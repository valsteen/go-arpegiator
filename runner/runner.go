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

	notesInDevice := devices.NewNoteInDevice(arpegiatorRunner.In)
	arpInDevice := devices.NewNoteInDevice(arpegiatorRunner.arpInPortPair.In)
	arpegiator := devices.NewArpegiator(notesInDevice, arpInDevice)

	arpegiator.AddMessageConsumer(func(data []byte) {
		_, err = arpegiatorRunner.arpInPortPair.Out.Write(data)
		s.MustNot(err)
		fmt.Println("Arp out message", data)
	})

	return arpegiatorRunner
}
