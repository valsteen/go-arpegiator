package devices

import (
	midiDefinitions "go-arpegiator/definitions"
	"go-arpegiator/services"
	"go-arpegiator/services/set"
)

type ArpSwitchSet struct {
	set.Set
}

func newArpSwitchSet(notes NoteSet) ArpSwitchSet {
	switches := make(set.Set, len(notes.Set))
	notes.Iterate(func(note midiDefinitions.NoteOnMessage) {
		switches.Add(ArpSwitch(note))
	})
	return ArpSwitchSet{switches}
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
	_added, _removed := s.Set.Compare(s2.Set)
	return newArpSwitchSetSlice(_added), newArpSwitchSetSlice(_removed)
}

func (s ArpSwitchSet) Iterate(cb func(e ArpSwitch)) {
	for _, e := range s.Set {
		cb(convertElementToArpSwitch(e))
	}
}
