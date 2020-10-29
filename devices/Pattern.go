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
	switches := make(set.Set, 0, len(notes.Set))
	notes.Iterate(func(note midiDefinitions.NoteOnMessage) {
		switches = switches.Add(PatternIterm(note))
	})
	return Pattern{switches}
}

func convertElementToPattern(e set.Element) PatternIterm {
	pattern, ok := e.(PatternIterm)
	services.Must(ok)
	return pattern
}

func newPatternSetSlice(s []set.Element) []PatternIterm {
	patternes := make([]PatternIterm, 0, len(s))

	for i, e := range s {
		patternes[i] = convertElementToPattern(e)
	}

	return patternes
}

func (s Pattern) Compare(s2 Pattern) ([]PatternIterm, []PatternIterm) {
	_added, _removed := s.Set.Compare(s2.Set)
	return newPatternSetSlice(_added), newPatternSetSlice(_removed)
}

func (s Pattern) Iterate(cb func(e PatternIterm)) {
	for _, e := range s.Set {
		cb(convertElementToPattern(e))
	}
}
