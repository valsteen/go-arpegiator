package devices

import (
	m "go-arpegiator/definitions"
)

type StickyNotesInDevice struct {
	*NotesInDevice
}

func (device *StickyNotesInDevice) allDeadNotes() bool {
	return device.NoteSet.Count(m.NoteOnMessage.IsDeadNote) == 0
}

func (device *StickyNotesInDevice) ConsumeMessage(channelMessage m.ChannelMessage) {
	switch message := channelMessage.(type) {
	case m.NoteOnMessage:
		deadNote := m.NewDeadNoteMessage(message.GetChannel(), message.GetPitch())
		device.NoteSet = device.NoteSet.Delete(deadNote) // delete dead note if matching
		// TODO -- new notes can replace dead notes , but how to do it ?
	case m.NoteOffMessage:
		if len(device.NoteSet.Set) > 0 {
			device.NoteSet = device.NoteSet.Delete(message)
			if device.allDeadNotes() {
				device.NoteSet = NewNoteSet(12)
			} else {
				deadNote := m.NewNoteOnMessage(message.GetChannel(), message.GetPitch(), 0)
				device.NoteSet = device.NoteSet.Add(deadNote)
			}
			device.send()
			return
		}
	}
	device.NotesInDevice.ConsumeMessage(channelMessage)
}
