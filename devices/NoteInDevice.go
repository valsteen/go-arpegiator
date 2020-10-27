package devices

import (
	"fmt"
	"gitlab.com/gomidi/midi"
	midiDefinitions "go-arpegiator/definitions"
	"go-arpegiator/services/set"
)

type NotesInDevice struct {
	NoteSet
	noteSetConsumers []NoteSetConsumer
}

func (device *NotesInDevice) consumeMessage(channelMessage midiDefinitions.ChannelMessage) {
	switch message := channelMessage.(type) {
	case midiDefinitions.NoteOnMessage:
		set.Set(device.NoteSet).Add(message)
	case midiDefinitions.NoteOffMessage:
		set.Set(device.NoteSet).Delete(message)
	default:
		fmt.Println("ignored", channelMessage)
		return
	}

	for _, consumer := range device.noteSetConsumers {
		consumer(device.NoteSet)
	}
}

func NewNoteInDevice(in midi.In) *NotesInDevice {
	notesInDevice := NotesInDevice{
		NoteSet:      make(NoteSet, 12),
		noteSetConsumers: make([]NoteSetConsumer, 0, 10),
	}
	pipeRawMessageToChannelMessage(in, notesInDevice.consumeMessage)
	return &notesInDevice
}

func (device *NotesInDevice) AddNoteSetConsumer(consumer NoteSetConsumer) {
	device.noteSetConsumers = append(device.noteSetConsumers, consumer)
}
