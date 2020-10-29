package runner

import (
	"fmt"
	midi "gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/rtmididrv"
	"go-arpegiator/definitions"
	"go-arpegiator/devices"
	s "go-arpegiator/services"
)

type ArpegiatorRunner struct {
	*rtmididrv.Driver
	midiNotesIn     midi.In
	patternPortPair *midiDefinitions.PortPair
}

func (a ArpegiatorRunner) Close() {
	_ = a.midiNotesIn.Close()
	a.patternPortPair.Close()
	if a.Driver != nil {
		_ = a.Driver.Close()
	}
}

func RunArpegiator(notesInName, arpName string) ArpegiatorRunner {
	driver, err := rtmididrv.New()
	s.MustNot(err)

	midiNotesIn, err := driver.OpenVirtualIn(notesInName)
	s.MustNot(err)

	arpegiatorRunner := ArpegiatorRunner{
		Driver:          driver,
		midiNotesIn:     midiNotesIn,
		patternPortPair: midiDefinitions.NewPortPair(arpName, driver),
	}

	//notesInDevice := devices.StickyNotesInDevice{NotesInDevice: devices.NewNoteInDevice()}
	notesInDevice := devices.NewNoteInDevice()
	devices.RawMessageToChannelMessageAdapter(arpegiatorRunner.midiNotesIn)(notesInDevice.ConsumeMessage)
	arpInDevice := devices.NewNoteInDevice()
	patternInAdapter := devices.RawMessageToChannelMessageAdapter(arpegiatorRunner.patternPortPair.In)
	patternInAdapter(arpInDevice.ConsumeMessage)

	arpegiator := devices.NewArpegiator(notesInDevice, arpInDevice)

	notesOutDevice := devices.NewNoteOutDevice()
	arpegiator.AddNoteSetConsumer(notesOutDevice.ConsumeNoteSet)
	notesOutDevice.AddMessageConsumer(func(data []byte) {
		fmt.Println(data)
		_, err = arpegiatorRunner.patternPortPair.Out.Write(data)
		s.MustNot(err)
	})

	patternInAdapter(
		devices.PressureFilter(func(message midiDefinitions.PressureMessage) {
			_, err = arpegiatorRunner.patternPortPair.Out.Write(message)
			s.MustNot(err)
		}),
	)

	return arpegiatorRunner
}
