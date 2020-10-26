package devices

import (
	"fmt"
	"gitlab.com/gomidi/midi"
	midiDefinitions "go-arpegiator/definitions"
	"go-arpegiator/services/set"
)

type NoteDevice struct {
	Notes            NoteSet
	noteSetConsumers []NoteSetConsumer
}

func (device *NoteDevice) consumeMessage(message midiDefinitions.ChannelMessage) {
	if noteMessage, ok := message.(midiDefinitions.NoteMessage); ok {
		if noteMessage.IsNoteOn() {
			set.Set(device.Notes).Add(noteMessage)
		} else {
			set.Set(device.Notes).Delete(noteMessage)
		}

		for _, consumer := range device.noteSetConsumers {
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
		Notes:            make(NoteSet),
		noteSetConsumers: make([]NoteSetConsumer, 0, 10),
	}
	pipeRawMessageToChannelMessage(in, device.consumeMessage)
	return &device
}

func (device *NoteDevice) AddNoteSetConsumer(consumer NoteSetConsumer) {
	device.noteSetConsumers = append(device.noteSetConsumers, consumer)
}
