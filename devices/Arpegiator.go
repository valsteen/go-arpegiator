package devices

import (
	"fmt"
	midiDefinitions "go-arpegiator/definitions"
)

/*
remaining work:
- given incoming arpswitch, arpegiator produces another set of notes on based on noteon on note in device, modulated
  by arp switch
- this is sent to notes out device which produces the note on/off message based on its own noteson state
 */
type Arpegiator struct {
	notes []midiDefinitions.NoteOnMessage
	ArpSwitchSet
	messageConsumers []MessageConsumer
}

func NewArpegiator(noteIn *NotesInDevice, arpIn *NotesInDevice) *Arpegiator {
	arpegiator := Arpegiator{
		notes:            make([]midiDefinitions.NoteOnMessage, 0, 12),
		ArpSwitchSet:     make(ArpSwitchSet, 12),
		messageConsumers: make([]MessageConsumer, 0, 10),
	}
	noteIn.AddNoteSetConsumer(arpegiator.consumeInNoteSet)
	arpIn.AddNoteSetConsumer(func(noteSet NoteSet) {
		arpegiator.consumeArpSwitchSet(newArpSwitchSet(noteSet))
	})
	return &arpegiator
}

func (a *Arpegiator) consumeInNoteSet(noteSet NoteSet) {
	a.notes = noteSet.Slice()
}

func (a *Arpegiator) consumeArpSwitchSet(arpSwitchSet ArpSwitchSet) {
	added, removed := a.ArpSwitchSet.Compare(arpSwitchSet)

	// TODO test removed, implement added ; try to refactor

	// TODO what we need here is express it the other way around: device out receives notes states, but outputs
	// messages to start or stop notes
	for _, arpSwitch := range removed {
		if int(arpSwitch.GetIndex()) < len(a.notes) {
			a.send(
				midiDefinitions.MakeNoteOffMessage(
					arpSwitch.GetChannel(),
					a.notes[arpSwitch.GetIndex()].GetPitch(),
					arpSwitch.GetVelocity(),
				),
			)
		}
	}

	// when added, find corresponding note(s) in noteset ( index can appear several times with different octaves )
	//that is the note in index ;
	a.ArpSwitchSet = arpSwitchSet

	fmt.Printf("added: %v removed: %v\n", added, removed)
}

func (a *Arpegiator) send(message []byte) {
	for _, consumer := range a.messageConsumers {
		consumer(message)
	}
}
