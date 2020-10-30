package midiDefinitions

import (
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/rtmididrv"
	"go-arpegiator/services"
)

func NewPortPair(name string, driver *rtmididrv.Driver) (midi.In, midi.Out) {
	in, err := driver.OpenVirtualIn(name)
	services.MustNot(err)
	out, err := driver.OpenVirtualOut(name)
	services.MustNot(err)
	return in, out
}
