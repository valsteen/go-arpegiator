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

	notesInDevice := devices.StickyNotesInDevice{NotesInDevice: devices.NewNoteInDevice()}
	//notesInDevice := devices.NewNoteInDevice()
	devices.PipeRawMessageToChannelMessage(arpegiatorRunner.In, notesInDevice.ConsumeMessage)
	arpInDevice := devices.NewNoteInDevice()
	devices.PipeRawMessageToChannelMessage(arpegiatorRunner.arpInPortPair.In, arpInDevice.ConsumeMessage)

	arpegiator := devices.NewArpegiator(notesInDevice, arpInDevice)

	notesOutDevice := devices.NewNoteOutDevice()
	arpegiator.AddNoteSetConsumer(notesOutDevice.ConsumeNoteSet)
	notesOutDevice.AddMessageConsumer(func(data []byte) {
		fmt.Println(data)
		_, err = arpegiatorRunner.arpInPortPair.Out.Write(data)
		s.MustNot(err)
	})

	return arpegiatorRunner
}
