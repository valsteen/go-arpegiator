package devices

import (
	"fmt"
	m "go-arpegiator/definitions"
	"go-arpegiator/services"
	"go-arpegiator/services/set"
)

type PatternIterm m.NoteOnMessage

func (a PatternIterm) GetIndex() byte {
	return m.NoteOnMessage(a).GetPitch() % 12
}

func (a PatternIterm) GetOctave() int8 {
	// C4 is considered octave 0
	return (int8(m.NoteOnMessage(a).GetPitch()) - 48) / 12
}

func (a PatternIterm) Transpose(note m.NoteOnMessage) m.NoteOnMessage {
	pitch := int(note.GetPitch()) + int(a.GetOctave())*12
	if pitch > 127 || pitch < 0 {
		return nil
	}

	if note.IsDeadNote() {
		return m.NewDeadNoteMessage(
			a.GetChannel(),
			byte(pitch),
		)
	}

	return m.NewNoteOnMessage(
		a.GetChannel(),
		byte(pitch),
		a.GetVelocity(),
	)
}

func (a PatternIterm) GetChannel() byte {
	return m.NoteOnMessage(a).GetChannel()
}

func (a PatternIterm) GetVelocity() byte {
	return m.NoteOnMessage(a).GetVelocity()
}

func (a PatternIterm) String() string {
	return fmt.Sprintf("switch = (%d %v %d)", m.NoteOnMessage(a).GetChannel(), a.GetOctave(), a.GetIndex())
}

func (a PatternIterm) Less(element set.Element) bool {
	a2, ok := element.(PatternIterm)
	services.Must(ok)
	if a.GetIndex() < a2.GetIndex() {
		return true
	} else if a.GetIndex() > a2.GetIndex() {
		return false
	}

	if a.GetOctave() < a2.GetOctave() {
		return true
	} else if a.GetOctave() > a2.GetOctave() {
		return false
	}
	return a.GetChannel() < a2.GetChannel()
}
