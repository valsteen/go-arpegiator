package devices

import (
	midiDefinitions "go-arpegiator/definitions"
)

type Arpegiator struct {
	notes            NoteSet
	noteSetConsumers []NoteSetConsumer
}

func NewArpegiator(noteIn INotesInDevice, arpIn *NotesInDevice) *Arpegiator {
	arpegiator := Arpegiator{
		notes:            NewNoteSet(12),
		noteSetConsumers: make([]NoteSetConsumer, 0, 10),
	}
	noteIn.AddNoteSetConsumer(arpegiator.consumeInNoteSet)
	arpIn.AddNoteSetConsumer(func(noteSet NoteSet) {
		arpegiator.consumeArpSwitchSet(newArpSwitchSet(noteSet))
	})
	return &arpegiator
}

func (a *Arpegiator) consumeInNoteSet(noteSet NoteSet) {
	a.notes = noteSet
}

func (a *Arpegiator) consumeArpSwitchSet(arpSwitchSet ArpSwitchSet) {
	noteSet := NewNoteSet(a.notes.Length())
	arpSwitchSet.Iterate(func(e ArpSwitch) {
		index := int(e.GetIndex())
		if index < a.notes.Length() {
			note := a.notes.At(index)
			velocity := e.GetVelocity()
			if note.GetVelocity() == 0 {
				velocity = 0 // sticky dead note
			}

			noteOut := midiDefinitions.NewNoteOnMessage(note.GetChannel(), note.GetPitch(), velocity)
			noteSet = noteSet.Add(noteOut)
		}
	})
	a.send(noteSet)
}

func (a *Arpegiator) send(noteSet NoteSet) {
	for _, consumer := range a.noteSetConsumers {
		consumer(noteSet)
	}
}

func (a *Arpegiator) AddNoteSetConsumer(consumer NoteSetConsumer) {
	a.noteSetConsumers = append(a.noteSetConsumers, consumer)
}
