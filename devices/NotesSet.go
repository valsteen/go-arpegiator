package devices

import (
	midiDefinitions "go-arpegiator/definitions"
	"go-arpegiator/services"
	"go-arpegiator/services/set"
)

type NoteSet set.Set

func convertElementToNote(e set.Element) midiDefinitions.Note {
	note, ok := e.(midiDefinitions.Note)
	services.Must(ok)
	return note
}

func newNoteSetSlice(s []set.Element) []midiDefinitions.Note {
	notes := make([]midiDefinitions.Note, len(s))

	for i, e := range s {
		notes[i] = convertElementToNote(e)
	}

	return notes
}

func (s NoteSet) Compare(s2 NoteSet) (added []midiDefinitions.Note, removed []midiDefinitions.Note) {
	_added, _removed := set.Set(s).Compare(set.Set(s2))
	return newNoteSetSlice(_added), newNoteSetSlice(_removed)
}

func (s NoteSet) Iterate(cb func(e midiDefinitions.Note)) {
	for _, e := range s {
		cb(convertElementToNote(e))
	}
}
