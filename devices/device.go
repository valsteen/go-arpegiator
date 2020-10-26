package devices

import (
	"fmt"
	"gitlab.com/gomidi/midi"
	midiDefinitions "go-arpegiator/definitions"
	"go-arpegiator/services"
	"go-arpegiator/services/set"
)

type Device struct {
	Notes set.Set
	notesConsumers []NotesConsumer
}

type ChannelMessageConsumer func(message midiDefinitions.ChannelMessage)

func pipeRawMessageToChannelMessage(in midi.In, consumer ChannelMessageConsumer) {
	err := in.SetListener(func(data []byte, deltaMicroseconds int64) {
		midiMessage := midiDefinitions.AsMidiMessage(data)
		if midiChannelMessage, ok := midiMessage.(midiDefinitions.ChannelMessage); ok {
			consumer(midiChannelMessage)
		}
	})
	services.MustNot(err)
}

func (device *Device) consume(message midiDefinitions.ChannelMessage) {
	// oops type assertion fails I don't even know why
	if noteMessage, ok := message.(midiDefinitions.NoteMessage); ok {
		if noteMessage.IsNoteOn() {
			device.Notes.Add(noteMessage)
		} else {
			device.Notes.Delete(noteMessage)
		}

		for _, consumer := range device.notesConsumers {
			consumer(device.Notes)
		}
	} else {
		fmt.Println("ignored", message)
	}
}

func (device Device) String() string {
	return fmt.Sprintf("Device state: %v", device.Notes)
}

func NewDevice(in midi.In) *Device {
	device := Device{
		Notes:          make(set.Set),
		notesConsumers: make([]NotesConsumer, 0, 10),
	}
	pipeRawMessageToChannelMessage(in, device.consume)
	return &device
}

type NotesConsumer func(notes set.Set)

func (device *Device) AddConsumer(consumer NotesConsumer) {
	device.notesConsumers = append(device.notesConsumers, consumer)
}
