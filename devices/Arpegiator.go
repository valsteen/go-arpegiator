package devices

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
		arpegiator.consumePattern(NewPattern(noteSet))
	})
	return &arpegiator
}

func (a *Arpegiator) consumeInNoteSet(noteSet NoteSet) {
	a.notes = noteSet
}

func (a *Arpegiator) consumePattern(pattern Pattern) {
	noteSet := NewNoteSet(a.notes.Length())
	pattern.Iterate(func(e PatternIterm) {
		index := int(e.GetIndex())
		if index < a.notes.Length() {
			note := a.notes.At(index)
			noteOut := e.Transpose(note)
			if noteOut != nil {
				noteSet = noteSet.Add(noteOut)
			}
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
