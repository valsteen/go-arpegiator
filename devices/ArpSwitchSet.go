package devices

import (
	midiDefinitions "go-arpegiator/definitions"
	"go-arpegiator/services"
	"go-arpegiator/services/set"
)

type ArpSwitchSet set.Set

func newArpSwitchSet(notes NoteSet) ArpSwitchSet {
	switches := make(ArpSwitchSet, len(notes))
	notes.Iterate(func(note midiDefinitions.Note) {
		set.Set(switches).Add(ArpSwitch{Note: note})
	})
	return switches
}

func convertElementToArpSwitch(e set.Element) ArpSwitch {
	arpSwitch, ok := e.(ArpSwitch)
	services.Must(ok)
	return arpSwitch
}

func newArpSwitchSetSlice(s []set.Element) []ArpSwitch {
	arpSwitches := make([]ArpSwitch, len(s))

	for i, e := range s {
		arpSwitches[i] = convertElementToArpSwitch(e)
	}

	return arpSwitches
}

func (s ArpSwitchSet) Compare(s2 ArpSwitchSet) ([]ArpSwitch, []ArpSwitch) {
	_added, _removed := set.Set(s).Compare(set.Set(s2))
	return newArpSwitchSetSlice(_added), newArpSwitchSetSlice(_removed)
}

func (s ArpSwitchSet) Iterate(cb func(e ArpSwitch)) {
	for _, e := range s {
		cb(convertElementToArpSwitch(e))
	}
}
