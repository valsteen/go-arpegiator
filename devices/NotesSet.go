package devices

import (
	m "go-arpegiator/definitions"
	"go-arpegiator/services"
	"go-arpegiator/services/set"
)

type NoteSet struct {
	set.Set
}

func convertElementToNote(e set.Element) m.RichNote {
	note, ok := e.(m.RichNote)
	services.Must(ok)
	return note
}

func (s NoteSet) Compare(s2 NoteSet) (added NoteSet, removed NoteSet) {
	_added, _removed := s.Set.Compare(s2.Set)
	return NoteSet{_added}, NoteSet{_removed}
}

func (s NoteSet) Iterate(cb func(e m.RichNote)) {
	for _, e := range s.Set {
		cb(convertElementToNote(e))
	}
}

func (s NoteSet) Add(e m.RichNote) NoteSet {
	return NoteSet{s.Set.Add(e)}
}

func (s NoteSet) Delete(e m.Note) NoteSet {
	return NoteSet{s.Set.Delete(e)}
}

func (s NoteSet) At(i int) m.RichNote {
	return convertElementToNote(s.Set.At(i))
}

func NewNoteSet(cap int) NoteSet {
	return NoteSet{make(set.Set, 0, cap)}
}

func (s NoteSet) Count(condition func(message m.RichNote) bool) int {
	count := 0
	s.Iterate(func(e m.RichNote) {
		if condition(e) {
			count++
		}
	})
	return count
}
