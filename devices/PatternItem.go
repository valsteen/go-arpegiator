package devices

import (
	"fmt"
	m "go-arpegiator/definitions"
	"go-arpegiator/services"
	"go-arpegiator/services/set"
)

type PatternItem m.RichNote

func (a PatternItem) GetIndex() byte {
	return m.RichNote(a).GetPitch() % 12
}

func (a PatternItem) GetOctave() int8 {
	// C4 is considered octave 0
	return int8(m.RichNote(a).GetPitch()) / 12 - 4
}

type TransposeError string

func (e TransposeError) Error() string {
	return string(e)
}

func (a PatternItem) Transpose(note m.RichNote) (m.RichNote, error) {
	pitch := int(note.GetPitch()) + int(a.GetOctave())*12
	if pitch > 127 || pitch < 0 {
		return m.RichNote{}, TransposeError("transpose out of bounds")
	}

	if note.IsDeadNote() {
		return m.NewDeadNote(
			a.GetChannel(),
			byte(pitch),
		), nil
	}

	return m.RichNote{
		NoteOnMessage: m.NewNoteOnMessage(
			a.GetChannel(),
			byte(pitch),
			a.GetVelocity(),
		),
		PressureMessage:  a.PressureMessage,
		PitchBendMessage: a.PitchBendMessage,
	}, nil
}

func (a PatternItem) GetChannel() byte {
	return m.RichNote(a).GetChannel()
}

func (a PatternItem) GetVelocity() byte {
	return m.RichNote(a).GetVelocity()
}

func (a PatternItem) String() string {
	return fmt.Sprintf("switch = (%d %v %d)", m.RichNote(a).GetChannel(), a.GetOctave(), a.GetIndex())
}

func (a PatternItem) Less(element set.Element) bool {
	a2, ok := element.(PatternItem)
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
