package devices

import (
	"fmt"
	"gitlab.com/gomidi/midi"
	midiDefinitions "go-arpegiator/definitions"
	"go-arpegiator/services"
)

type midiChannelMessageChan chan midiDefinitions.ChannelMessage
type notes []midiDefinitions.Note
type StateChangeConsumer chan notes

type Device struct {
	notes
	stateConsumers []StateChangeConsumer
}

func pipeRawMessageToChannelMessage(in midi.In) (channel midiChannelMessageChan) {
	channel = make(midiChannelMessageChan)
	err := in.SetListener(func(data []byte, deltaMicroseconds int64) {
		midiMessage := midiDefinitions.AsMidiMessage(data)
		if midiChannelMessage, ok := midiMessage.(midiDefinitions.ChannelMessage); ok {
			channel <- midiChannelMessage
		}
	})
	services.MustNot(err)
	return
}

func (notes notes) insert(note midiDefinitions.Note, pos int) (out notes) {
	out = append(notes, nil)
	copy(out[pos+1:], notes[pos:])
	out[pos] = note
	return out
}

func (notes notes) remove(pos int) notes {
	if pos < len(notes)-1 {
		copy(notes[pos:], notes[pos+1:])
	}
	return notes[:len(notes)-1]
}

func (device *Device) consume(messageChan midiChannelMessageChan) {
	for message := range messageChan {
		if noteMessage, ok := message.(midiDefinitions.NoteMessage); ok {
			pitchIn := noteMessage.GetPitch()

			if noteMessage.IsNoteOn() {
				for i, note := range device.notes {
					switch pitch := note.GetPitch(); {
					case pitch == pitchIn:
						return
					case pitch < pitchIn:
						device.notes = device.notes.insert(noteMessage, i)
						return
					}
				}
				device.notes = device.notes.insert(noteMessage, 0)
			} else {
				for i, note := range device.notes {
					pitch := note.GetPitch()
					if pitch == pitchIn {
						device.notes = device.notes.remove(i)
					}
				}
			}

			for _, consumer := range device.stateConsumers {
				consumer <- device.notes
			}
		} else {
			fmt.Println("ignored", message)
		}

	}
}

func (device Device) String() string {
	return fmt.Sprintf("Device state: %s", device.notes)
}

func New(in midi.In) *Device {
	device := Device{
		notes:          make(notes, 0, 12),
		stateConsumers: make([]StateChangeConsumer, 0, 10),
	}
	go device.consume(pipeRawMessageToChannelMessage(in))
	return &device
}

func (device *Device) AddConsumer(consumer StateChangeConsumer) {
	device.stateConsumers = append(device.stateConsumers, consumer)
}
