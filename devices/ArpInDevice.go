package devices

import (
	"fmt"
	midiDefinitions "go-arpegiator/definitions"
	"go-arpegiator/services"
	"go-arpegiator/services/set"
)

type ArpSwitch struct {
	midiDefinitions.Note
}

type ArpInDevice struct {
	switches ArpSwitchSet
}

func (d *ArpInDevice) Consume(notes NoteSet) {
	newSwitches := NewSwitches(notes)
	added, removed := d.switches.Diff(newSwitches)
	fmt.Printf("added: %v removed: %v\n", added, removed)
	d.switches = newSwitches
}

func (a ArpSwitch) GetIndex() byte {
	return a.GetPitch() % 12
}

func (a ArpSwitch) GetOctave() int8 {
	// C4 is considered octave 0
	return (int8(a.GetPitch()) - 60) / 12
}

func (a ArpSwitch) String() string {
	return fmt.Sprintf("switch = (%d %v %d)", a.GetChannel(), a.GetOctave(), a.GetIndex())
}

func NewArpInDevice() *ArpInDevice {
	return &ArpInDevice{
		switches: make(ArpSwitchSet),
	}
}

func NewSwitches(notes NoteSet) ArpSwitchSet {
	switches := make(ArpSwitchSet)
	notes.Iterate(func(e midiDefinitions.Note) {
		switches.Add(ArpSwitch{Note: e})
	})
	return switches
}

type ArpSwitchSet set.Set

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
