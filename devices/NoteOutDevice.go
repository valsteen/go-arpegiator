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

	removed.Iterate(func(noteOnMessage midiDefinitions.NoteOnMessage) {
		device.send(
			midiDefinitions.NewNoteOffMessage(
				noteOnMessage.GetChannel(),
				noteOnMessage.GetPitch(),
				// design issue: note off velocity cannot be implemented if just considering a set of active notes
				0,
			),
		)
	})

	added.Iterate(func(noteOnMessage midiDefinitions.NoteOnMessage) {
		if noteOnMessage.GetVelocity() > 0 {
			// velocity 0 is a sticky dead note,
			// we keep other notes in position and don't play this one
			device.send(
				midiDefinitions.NewNoteOnMessage(
					noteOnMessage.GetChannel(),
					noteOnMessage.GetPitch(),
					noteOnMessage.GetVelocity(),
				),
			)
		}
	})

	device.noteSet = noteSet

	//	fmt.Printf("added: %v removed: %v\n", added, removed)
}

func NewNoteOutDevice() *NotesOutDevice {
	notesInDevice := NotesOutDevice{
		noteSet:          NewNoteSet(12),
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
