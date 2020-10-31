package devices

import (
	midiDefinitions "go-arpegiator/definitions"
	"go-arpegiator/services"
	"go-arpegiator/services/set"
)

type Pattern struct {
	set.Set
}

func NewPattern(notes NoteSet) Pattern {
	patterns := make(set.Set, 0, len(notes.Set))
	notes.Iterate(func(note midiDefinitions.RichNote) {
		patterns = patterns.Add(PatternItem(note))
	})
	return Pattern{patterns}
}

func convertElementToPattern(e set.Element) PatternItem {
	pattern, ok := e.(PatternItem)
	services.Must(ok)
	return pattern
}

func newPatternSetSlice(s []set.Element) []PatternItem {
	patternes := make([]PatternItem, 0, len(s))

	for i, e := range s {
		patternes[i] = convertElementToPattern(e)
	}

	return patternes
}

func (s Pattern) Compare(s2 Pattern) ([]PatternItem, []PatternItem) {
	_added, _removed := s.Set.Compare(s2.Set)
	return newPatternSetSlice(_added), newPatternSetSlice(_removed)
}

func (s Pattern) Iterate(cb func(e PatternItem)) {
	for _, e := range s.Set {
		cb(convertElementToPattern(e))
	}
}
