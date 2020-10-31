package midiDefinitions


type RichNote struct {
	NoteOnMessage
	PressureMessage
	PitchBendMessage
}

func (n RichNote) GetChannel() byte {
	return n.NoteOnMessage.GetChannel()
}

func NewDeadNote(channel, pitch byte) RichNote {
	return RichNote{
		NoteOnMessage{
			channel + NOTEON,
			pitch,
			0,
		},
		nil,
		nil,
	}
}

func (n RichNote) IsDeadNote() bool {
	return n.GetVelocity() == 0
}
