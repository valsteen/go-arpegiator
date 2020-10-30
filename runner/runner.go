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

	// alternate method
	// notesInDevice := devices.StickyNotesInDevice{NotesInDevice: devices.NewNoteInDevice()}
	notesInDevice := devices.NewNoteInDevice()
	patternInDevice := devices.NewNoteInDevice()
	notesOutDevice := devices.NewNoteOutDevice()

	// give notes and pattern devices to arpegiator
	arpegiator := devices.NewArpegiator(notesInDevice, patternInDevice)
	// arpegiator outputs to notes output device
	arpegiator.AddNoteSetConsumer(notesOutDevice.ConsumeNoteSet)

	// adapter subscribes to midiNotesIn, then gives notesInDevice as receiver
	midiInAdapter := devices.RawMessageToChannelMessageAdapter(arpegiatorRunner.midiNotesIn)
	midiInAdapter(notesInDevice.ConsumeMessage)

	// adapter subscribes to pattern in, then give patternInDevice as receiver
	patternInAdapter := devices.RawMessageToChannelMessageAdapter(arpegiatorRunner.patternPortPair.In)
	patternInAdapter(patternInDevice.ConsumeMessage)

	// notes out device outputs to midi out and console
	notesOutDevice.AddMessageConsumer(devices.FailOnWriteErrorAdapter(arpegiatorRunner.patternPortPair.Out.Write))
	notesOutDevice.AddMessageConsumer(devices.FailOnPrintErrorAdapter(fmt.Println))

	// pressure is filtered out from notes and pattern devices, consume then from pattern in and output to midi out
	addPressureConsumer := devices.PressureFilter(patternInAdapter)
	addPressureConsumer(devices.FailOnWritePressureAdapter(arpegiatorRunner.patternPortPair.Out.Write))

	return arpegiatorRunner
}
