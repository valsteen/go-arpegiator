package devices

import (
	"fmt"
	"gitlab.com/gomidi/midi"
	midiDefinitions "go-arpegiator/definitions"
	"go-arpegiator/services"
	"go-arpegiator/services/set"
)

type Device struct {
	Notes          NoteSet
	notesConsumers []NotesConsumer
}

type ChannelMessageConsumer func(message midiDefinitions.ChannelMessage)

func pipeRawMessageToChannelMessage(in midi.In, consumer ChannelMessageConsumer) {
	err := in.SetListener(func(data []byte, deltaMicroseconds int64) {
		midiMessage := midiDefinitions.AsMidiMessage(data)
		if midiChannelMessage, ok := midiMessage.(midiDefinitions.ChannelMessage); ok {
			consumer(midiChannelMessage)
		}
	})
	services.MustNot(err)
}

func (device *Device) consume(message midiDefinitions.ChannelMessage) {
	if noteMessage, ok := message.(midiDefinitions.NoteMessage); ok {
		if noteMessage.IsNoteOn() {
			device.Notes.Add(noteMessage)
		} else {
			device.Notes.Delete(noteMessage)
		}

		for _, consumer := range device.notesConsumers {
			consumer(device.Notes)
		}
	} else {
		fmt.Println("ignored", message)
	}
}

func (device Device) String() string {
	return fmt.Sprintf("Device state: %v", device.Notes)
}

func NewDevice(in midi.In) *Device {
	device := Device{
		Notes:          make(NoteSet),
		notesConsumers: make([]NotesConsumer, 0, 10),
	}
	pipeRawMessageToChannelMessage(in, device.consume)
	return &device
}

type NotesConsumer func(notes NoteSet)

func (device *Device) AddConsumer(consumer NotesConsumer) {
	device.notesConsumers = append(device.notesConsumers, consumer)
}

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
