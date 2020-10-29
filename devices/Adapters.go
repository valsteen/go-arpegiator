package devices

import (
	"gitlab.com/gomidi/midi"
	m "go-arpegiator/definitions"
	"go-arpegiator/services"
)

type ChannelMessageConsumer func(message m.ChannelMessage)
type NoteSetConsumer func(notes NoteSet)
type PressureMessageConsumer func(message m.PressureMessage)
type PitchBendMessageConsumer func(message m.PitchBendMessage)

func RawMessageToChannelMessageAdapter(in midi.In) (consumer func(ChannelMessageConsumer)) {
	consumers := make([]ChannelMessageConsumer, 0)

	err := in.SetListener(func(data []byte, deltaMicroseconds int64) {
		midiMessage := m.AsMidiMessage(data)
		for _, consumer := range consumers {
			if midiChannelMessage, ok := midiMessage.(m.ChannelMessage); ok {
				consumer(midiChannelMessage)
			}
		}
	})
	services.MustNot(err)

	return func(consumer ChannelMessageConsumer) {
		consumers = append(consumers, consumer)
	}
}

func PressureFilter(consumer PressureMessageConsumer) (receiver func(message m.ChannelMessage)) {
	return func(message m.ChannelMessage) {
		if pressureMessage, ok := message.(m.PressureMessage); ok {
			consumer(pressureMessage)
		}
	}
}

func PitchBendFilter(consumer PitchBendMessageConsumer) ChannelMessageConsumer {
	return func(message m.ChannelMessage) {
		if pressureMessage, ok := message.(m.PitchBendMessage); ok {
			consumer(pressureMessage)
		}
	}
}
