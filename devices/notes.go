package devices

import midiDefinitions "go-arpegiator/definitions"

type notes []midiDefinitions.Note

func (notes notes) insertAt(note midiDefinitions.Note, pos int) (out notes) {
	out = append(notes, nil)
	copy(out[pos+1:], notes[pos:])
	out[pos] = note
	return out
}

func (notes notes) removeAt(pos int) notes {
	if pos < len(notes)-1 {
		copy(notes[pos:], notes[pos+1:])
	}
	return notes[:len(notes)-1]
}

func (notes notes) insert(noteMessage midiDefinitions.NoteMessage) notes {
	pitchIn := noteMessage.GetPitch()
	for i, note := range notes {
		switch pitch := note.GetPitch(); {
		case pitch == pitchIn:
			return notes
		case pitch > pitchIn:
			return notes.insertAt(noteMessage, i)
		}
	}

	return append(notes, noteMessage)
}

func (notes notes) remove(noteMessage midiDefinitions.NoteMessage) notes {
	pitchIn := noteMessage.GetPitch()
	for i, note := range notes {
		pitch := note.GetPitch()
		if pitch == pitchIn {
			return notes.removeAt(i)
		}
	}
	return notes
}
