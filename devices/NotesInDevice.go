package devices

import (
	"fmt"
	midiDefinitions "go-arpegiator/definitions"
	"go-arpegiator/services/set"
)

type INotesInDevice interface {
	AddNoteSetConsumer(consumer NoteSetConsumer)
}

type NotesInDevice struct {
	NoteSet
	noteSetConsumers []NoteSetConsumer
}

func (device *NotesInDevice) ConsumeMessage(channelMessage midiDefinitions.ChannelMessage) {
	switch message := channelMessage.(type) {
	case midiDefinitions.NoteOnMessage:
		device.NoteSet.Add(message)
	case midiDefinitions.NoteOffMessage:
		device.NoteSet.Delete(message)
	default:
		fmt.Println("ignored", channelMessage)
		return
	}
	device.send()
}

func (device *NotesInDevice) send() {
	for _, consumer := range device.noteSetConsumers {
		consumer(device.NoteSet)
	}
}

func NewNoteInDevice() *NotesInDevice {
	noteSetConsumers := make([]NoteSetConsumer, 0, 10)
	notesInDevice := &NotesInDevice{
		NoteSet:          NoteSet{make(set.Set, 12)},
		noteSetConsumers: noteSetConsumers,
	}
	_ = INotesInDevice(notesInDevice) // interface check
	return notesInDevice
}

func (device *NotesInDevice) AddNoteSetConsumer(consumer NoteSetConsumer) {
	device.noteSetConsumers = append(device.noteSetConsumers, consumer)
}
