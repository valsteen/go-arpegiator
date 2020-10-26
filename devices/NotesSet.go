package devices

import (
	midiDefinitions "go-arpegiator/definitions"
	"go-arpegiator/services"
	"go-arpegiator/services/set"
)

type NoteSet set.Set

func (s NoteSet) Delete(e midiDefinitions.Note) {
	set.Set(s).Delete(e)
}

func (s NoteSet) Add(e midiDefinitions.Note) {
	set.Set(s).Add(e)
}

func (s NoteSet) Diff(s2 NoteSet) (added []midiDefinitions.Note, removed []midiDefinitions.Note) {
	_added, _removed := set.Set(s).Diff(set.Set(s2))
	added = make([]midiDefinitions.Note, len(_added))
	removed = make([]midiDefinitions.Note, len(_removed))

	for i, e := range _added {
		note, ok := e.(midiDefinitions.Note)
		services.Must(ok)
		added[i] = note
	}

	for i, e := range _removed {
		note, ok := e.(midiDefinitions.Note)
		services.Must(ok)
		removed[i] = note
	}

	return
}

func (s NoteSet) Iterate(cb func(e midiDefinitions.Note)) {
	for _, e := range s {
		note, ok := e.(midiDefinitions.Note)
		services.Must(ok)
		cb(note)
	}
}
