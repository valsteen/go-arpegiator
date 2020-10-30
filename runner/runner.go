package runner

import (
	"fmt"
	"gitlab.com/gomidi/rtmididrv"
	"go-arpegiator/definitions"
	"go-arpegiator/devices"
	s "go-arpegiator/services"
)

type IClosable interface {
	Close() error
}

type Closables []IClosable

func (closables Closables) Close() {
	for _, closable := range closables {
		_ = closable.Close()
	}
}

func NewClosables(closables ...IClosable) Closables {
	return closables
}

func RunArpegiator(notesInName, arpName string) Closables {
	driver, err := rtmididrv.New()
	s.MustNot(err)

	midiNotesIn, err := driver.OpenVirtualIn(notesInName)
	s.MustNot(err)

	patternMidiIn, patternMidiOut := midiDefinitions.NewPortPair(arpName, driver)

	closer := NewClosables(driver, midiNotesIn, patternMidiIn, patternMidiOut)

	// alternate method
	// notesInDevice := devices.StickyNotesInDevice{NotesInDevice: devices.NewNoteInDevice()}
	notesInDevice := devices.NewNoteInDevice()
	patternInDevice := devices.NewNoteInDevice()
	notesOutDevice := devices.NewNoteOutDevice(
		// notes out device outputs to midi out and console
		devices.FailOnWriteErrorAdapter(patternMidiOut.Write),
		devices.FailOnPrintErrorAdapter(fmt.Println),
	)

	// give notes and pattern devices to arpegiator, outputs to notes output device
	devices.NewArpegiator(notesInDevice, patternInDevice, notesOutDevice.ConsumeNoteSet)

	devices.RawMessageToChannelMessageAdapter(midiNotesIn, notesInDevice.ConsumeMessage)
	devices.RawMessageToChannelMessageAdapter(
		patternMidiIn,
		patternInDevice.ConsumeMessage,
		// pressure is filtered out from notes and pattern devices, consume then from pattern in and output to midi out
		devices.PressureFilter(devices.FailOnWritePressureAdapter(patternMidiOut.Write)),
	)

	return closer
}
