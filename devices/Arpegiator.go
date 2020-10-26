package devices

import "fmt"

type Arpegiator struct {
	NoteSet
	ArpSwitchSet
	messageConsumers []ChannelMessageConsumer
}

func NewArpegiator(noteIn *NoteDevice, arpIn *NoteDevice) *Arpegiator {
	arpegiator := Arpegiator{
		NoteSet:          make(NoteSet, 12),
		ArpSwitchSet:     make(ArpSwitchSet, 12),
		messageConsumers: make([]ChannelMessageConsumer, 0, 10),
	}
	noteIn.AddNoteSetConsumer(arpegiator.consumeInNoteSet)
	arpIn.AddNoteSetConsumer(func(noteSet NoteSet) {
		arpegiator.consumeArpSwitchSet(newArpSwitchSet(noteSet))
	})
	return &arpegiator
}

func (a *Arpegiator) consumeInNoteSet(noteSet NoteSet) {
	a.NoteSet = noteSet
}

func (a *Arpegiator) consumeArpSwitchSet(arpSwitchSet ArpSwitchSet) {
	added, removed := a.ArpSwitchSet.Compare(arpSwitchSet)
	a.ArpSwitchSet = arpSwitchSet
	fmt.Printf("added: %v removed: %v\n", added, removed)
}

func (a *Arpegiator) AddMessageConsumer(consumer ChannelMessageConsumer) {
	a.messageConsumers = append(a.messageConsumers, consumer)
}
