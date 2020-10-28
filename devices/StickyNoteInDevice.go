package devices

import (
	midiDefinitions "go-arpegiator/definitions"
)

type StickyNotesInDevice struct {
	*NotesInDevice
}

func (device *StickyNotesInDevice) allDeadNotes() bool {
	alive := 0
	device.NoteSet.Iterate(func(e midiDefinitions.NoteOnMessage) {
		if e.GetVelocity() > 0 {
			alive += 1
		}
	})
	return alive == 0
}

func (device *StickyNotesInDevice) ConsumeMessage(channelMessage midiDefinitions.ChannelMessage) {
	switch message := channelMessage.(type) {
	case midiDefinitions.NoteOnMessage:
		deadNote := midiDefinitions.NewNoteOnMessage(message.GetChannel(), message.GetPitch(), 0)
		device.NoteSet = device.NoteSet.Delete(deadNote) // delete dead note if matching
		// TODO -- new notes can replace dead notes , but how to do it ?
	case midiDefinitions.NoteOffMessage:
		if len(device.NoteSet.Set) > 0 {
			device.NoteSet = device.NoteSet.Delete(message)
			if device.allDeadNotes() {
				device.NoteSet = NewNoteSet(12)
			} else {
				deadNote := midiDefinitions.NewNoteOnMessage(message.GetChannel(), message.GetPitch(), 0)
				device.NoteSet = device.NoteSet.Add(deadNote)
			}
			device.send()
			return
		}
	}
	device.NotesInDevice.ConsumeMessage(channelMessage)
}
