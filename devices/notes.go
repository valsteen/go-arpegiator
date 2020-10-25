package devices

import midiDefinitions "go-arpegiator/definitions"

type Notes []midiDefinitions.Note

func (notes Notes) insertAt(note midiDefinitions.Note, pos int) (out Notes) {
	out = append(notes, nil)
	copy(out[pos+1:], notes[pos:])
	out[pos] = note
	return out
}

func (notes Notes) removeAt(pos int) Notes {
	if pos < len(notes)-1 {
		copy(notes[pos:], notes[pos+1:])
	}
	return notes[:len(notes)-1]
}

func (notes Notes) insert(noteMessage midiDefinitions.NoteMessage) Notes {
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

func (notes Notes) remove(noteMessage midiDefinitions.NoteMessage) Notes {
	pitchIn := noteMessage.GetPitch()
	for i, note := range notes {
		pitch := note.GetPitch()
		if pitch == pitchIn {
			return notes.removeAt(i)
		}
	}
	return notes
}
