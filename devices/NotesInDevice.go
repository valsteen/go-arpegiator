package devices

import (
	"fmt"
	d "go-arpegiator/definitions"
)

type INotesInDevice interface {
	AddNoteSetConsumer(consumer NoteSetConsumer)
}

type NotesInDevice struct {
	NoteSet
	noteSetConsumers []NoteSetConsumer
	pressures []d.PressureMessage
	pitchBends []d.PitchBendMessage
}

func (device *NotesInDevice) ConsumeMessage(channelMessage d.ChannelMessage) {
	switch message := channelMessage.(type) {
	case d.NoteOnMessage:
		device.NoteSet = device.NoteSet.Add(d.RichNote{
			NoteOnMessage:    message,
			PressureMessage:  device.pressures[message.GetChannel()],
			PitchBendMessage: device.pitchBends[message.GetChannel()],
		})
	case d.NoteOffMessage:
		device.NoteSet = device.NoteSet.Delete(message)
	case d.PressureMessage:
		device.pressures[message.GetChannel()] = message
		device.NoteSet.Iterate(func(e d.RichNote) {
			if e.GetChannel() == message.GetChannel() {
				e.PressureMessage = message
			}
		})
	case d.PitchBendMessage:
		device.pitchBends[message.GetChannel()] = message
		device.NoteSet.Iterate(func(e d.RichNote) {
			if e.GetChannel() == message.GetChannel() {
				e.PitchBendMessage = message
			}
		})
	default:
		fmt.Println("ignored", channelMessage)
		return
	}
	device.send()
}

func (device *NotesInDevice) send() {
	for _, consumer := range device.noteSetConsumers {
		consumer(device.NoteSet)
	}
}

func NewNoteInDevice() *NotesInDevice {
	noteSetConsumers := make([]NoteSetConsumer, 0, 10)
	notesInDevice := &NotesInDevice{
		NoteSet:          NewNoteSet(12),
		noteSetConsumers: noteSetConsumers,
		pressures: make([]d.PressureMessage, 12),
		pitchBends: make([]d.PitchBendMessage, 12),
	}
	_ = INotesInDevice(notesInDevice) // interface check
	return notesInDevice
}

func (device *NotesInDevice) AddNoteSetConsumer(consumer NoteSetConsumer) {
	device.noteSetConsumers = append(device.noteSetConsumers, consumer)
}
