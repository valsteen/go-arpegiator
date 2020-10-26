package devices

import (
	midiDefinitions "go-arpegiator/definitions"
	"go-arpegiator/services"
	"go-arpegiator/services/set"
)

type ArpSwitchSet set.Set

func newArpSwitchSet(notes NoteSet) ArpSwitchSet {
	switches := make(ArpSwitchSet)
	notes.Iterate(func(e midiDefinitions.Note) {
		switches.Add(ArpSwitch{Note: e})
	})
	return switches
}

func (s ArpSwitchSet) Delete(e ArpSwitch) {
	set.Set(s).Delete(e)
}

func (s ArpSwitchSet) Add(e ArpSwitch) {
	set.Set(s).Add(e)
}

func (s ArpSwitchSet) Diff(s2 ArpSwitchSet) (added []ArpSwitch, removed []ArpSwitch) {
	_added, _removed := set.Set(s).Diff(set.Set(s2))
	added = make([]ArpSwitch, len(_added))
	removed = make([]ArpSwitch, len(_removed))

	for i, e := range _added {
		arpSwitch, ok := e.(ArpSwitch)
		services.Must(ok)
		added[i] = arpSwitch
	}

	for i, e := range _removed {
		arpSwitch, ok := e.(ArpSwitch)
		services.Must(ok)
		removed[i] = arpSwitch
	}

	return
}

func (s ArpSwitchSet) Iterate(cb func(e ArpSwitch)) {
	for _, e := range s {
		note, ok := e.(ArpSwitch)
		services.Must(ok)
		cb(note)
	}
}
