package devices

import (
	"fmt"
	midiDefinitions "go-arpegiator/definitions"
)

type NotesOutDevice struct {
	NoteSet
	messageConsumers []MessageConsumer
}

func (device *NotesOutDevice) consumeNoteSet(noteSet NoteSet) {
	added, removed := device.NoteSet.Compare(noteSet)

	for _, noteOnMessage := range removed {
		device.send(
			midiDefinitions.MakeNoteOffMessage(
				noteOnMessage.GetChannel(),
				noteOnMessage.GetPitch(),
				// design issue: note off velocity cannot be implemented if just considering a set of active notes
				0,
			),
		)
	}

	for _, noteOnMessage := range added {
		device.send(
			midiDefinitions.MakeNoteOnMessage(
				noteOnMessage.GetChannel(),
				noteOnMessage.GetPitch(),
				noteOnMessage.GetVelocity(),
			),
		)
	}

	device.NoteSet = noteSet

	fmt.Printf("added: %v removed: %v\n", added, removed)
}

func NewNoteOutDevice() *NotesOutDevice {
	notesInDevice := NotesOutDevice{
		NoteSet:      make(NoteSet, 12),
		messageConsumers: make([]MessageConsumer, 0, 10),
	}
	//pipeRawMessageToChannelMessage(out, notesInDevice.consumeMessage)
	return &notesInDevice
}

func (device *NotesOutDevice) AddMessageConsumer(consumer MessageConsumer) {
	device.messageConsumers = append(device.messageConsumers, consumer)
}

type MessageConsumer func([]byte)

func (device *NotesOutDevice) send(message []byte) {
	for _, consumer := range device.messageConsumers {
		consumer(message)
	}
}
