package devices

import (
	"fmt"
	"gitlab.com/gomidi/midi"
	midiDefinitions "go-arpegiator/definitions"
	"go-arpegiator/services"
)

type NoteDevice struct {
	Notes          NoteSet
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

func (device *NoteDevice) consumeNotes(message midiDefinitions.ChannelMessage) {
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

func (device NoteDevice) String() string {
	return fmt.Sprintf("NoteDevice state: %v", device.Notes)
}

func NewNoteDevice(in midi.In) *NoteDevice {
	device := NoteDevice{
		Notes:          make(NoteSet),
		notesConsumers: make([]NotesConsumer, 0, 10),
	}
	pipeRawMessageToChannelMessage(in, device.consumeNotes)
	return &device
}

type NotesConsumer func(notes NoteSet)

func (device *NoteDevice) AddConsumer(consumer NotesConsumer) {
	device.notesConsumers = append(device.notesConsumers, consumer)
}
