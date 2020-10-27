package devices

import (
	midiDefinitions "go-arpegiator/definitions"
	"go-arpegiator/services/set"
)

type Arpegiator struct {
	notes            []midiDefinitions.NoteOnMessage
	noteSetConsumers []NoteSetConsumer
}

func NewArpegiator(noteIn *NotesInDevice, arpIn *NotesInDevice) *Arpegiator {
	arpegiator := Arpegiator{
		notes:            make([]midiDefinitions.NoteOnMessage, 0, 12),
		noteSetConsumers: make([]NoteSetConsumer, 0, 10),
	}
	noteIn.AddNoteSetConsumer(arpegiator.consumeInNoteSet)
	arpIn.AddNoteSetConsumer(func(noteSet NoteSet) {
		arpegiator.consumeArpSwitchSet(newArpSwitchSet(noteSet))
	})
	return &arpegiator
}

func (a *Arpegiator) consumeInNoteSet(noteSet NoteSet) {
	a.notes = noteSet.Slice()
}

func (a *Arpegiator) consumeArpSwitchSet(arpSwitchSet ArpSwitchSet) {
	noteSet := make(NoteSet)
	arpSwitchSet.Iterate(func(e ArpSwitch) {
		index := int(e.GetIndex())
		if index < len(a.notes) {
			note := a.notes[index]
			noteOut := midiDefinitions.NewNoteOnMessage(note.GetChannel(), note.GetPitch(), e.GetVelocity())
			set.Set(noteSet).Add(noteOut)
		}
	})
	a.send(noteSet)
}

func (a *Arpegiator) send(set NoteSet) {
	for _, consumer := range a.noteSetConsumers {
		consumer(set)
	}
}

func (a *Arpegiator) AddNoteSetConsumer(consumer NoteSetConsumer) {
	a.noteSetConsumers = append(a.noteSetConsumers, consumer)
}
