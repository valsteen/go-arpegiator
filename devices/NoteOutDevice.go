package devices

import (
	midiDefinitions "go-arpegiator/definitions"
)

type NotesOutDevice struct {
	noteSet          NoteSet
	messageConsumers []MessageConsumer
}

func (device *NotesOutDevice) ConsumeNoteSet(noteSet NoteSet) {
	added, removed := device.noteSet.Compare(noteSet)

	for _, noteOnMessage := range removed {
		device.send(
			midiDefinitions.NewNoteOffMessage(
				noteOnMessage.GetChannel(),
				noteOnMessage.GetPitch(),
				// design issue: note off velocity cannot be implemented if just considering a set of active notes
				0,
			),
		)
	}

	for _, noteOnMessage := range added {
		device.send(
			midiDefinitions.NewNoteOnMessage(
				noteOnMessage.GetChannel(),
				noteOnMessage.GetPitch(),
				noteOnMessage.GetVelocity(),
			),
		)
	}

	device.noteSet = noteSet

	//	fmt.Printf("added: %v removed: %v\n", added, removed)
}

func NewNoteOutDevice() *NotesOutDevice {
	notesInDevice := NotesOutDevice{
		noteSet:          make(NoteSet, 12),
		messageConsumers: make([]MessageConsumer, 0, 10),
	}
	return &notesInDevice
}

func (device *NotesOutDevice) AddMessageConsumer(consumer MessageConsumer) {
	device.messageConsumers = append(device.messageConsumers, consumer)
}

type MessageConsumer func(data []byte)

func (device *NotesOutDevice) send(message []byte) {
	for _, consumer := range device.messageConsumers {
		consumer(message)
	}
}
