package devices

import (
	midiDefinitions "go-arpegiator/definitions"
	"go-arpegiator/services"
	"go-arpegiator/services/set"
	"sort"
)

type NoteSet set.Set

func convertElementToNote(e set.Element) midiDefinitions.NoteOnMessage {
	note, ok := e.(midiDefinitions.NoteOnMessage)
	services.Must(ok)
	return note
}

func newNoteSetSlice(s []set.Element) []midiDefinitions.NoteOnMessage {
	notes := make([]midiDefinitions.NoteOnMessage, len(s))

	for i, e := range s {
		notes[i] = convertElementToNote(e)
	}

	return notes
}

func (s NoteSet) Compare(s2 NoteSet) (added []midiDefinitions.NoteOnMessage, removed []midiDefinitions.NoteOnMessage) {
	_added, _removed := set.Set(s).Compare(set.Set(s2))
	return newNoteSetSlice(_added), newNoteSetSlice(_removed)
}

func (s NoteSet) Iterate(cb func(e midiDefinitions.NoteOnMessage)) {
	for _, e := range s {
		cb(convertElementToNote(e))
	}
}

// TODO maybe a set should not be a map but a sorted slice
func (s NoteSet) Slice() []midiDefinitions.NoteOnMessage {
	slice := make([]midiDefinitions.NoteOnMessage, 0, len(s))
	s.Iterate(func(e midiDefinitions.NoteOnMessage) {
		slice = append(slice, e)
	})
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].GetPitch() < slice[j].GetPitch()
	})
	return slice
}
