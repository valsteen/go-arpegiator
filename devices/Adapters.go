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

func PressureFilter(consumerReceiver func(ChannelMessageConsumer)) func(consumer PressureMessageConsumer) {
	consumers := make([]PressureMessageConsumer, 0)

	consumerReceiver(func(message m.ChannelMessage) {
		if pressureMessage, ok := message.(m.PressureMessage); ok {
			for _, consumer := range consumers {
				consumer(pressureMessage)
			}
		}
	})

	return func(consumer PressureMessageConsumer) {
		consumers = append(consumers, consumer)
	}
}

func PitchBendFilter(consumer PitchBendMessageConsumer) ChannelMessageConsumer {
	return func(message m.ChannelMessage) {
		if pressureMessage, ok := message.(m.PitchBendMessage); ok {
			consumer(pressureMessage)
		}
	}
}

func FailOnWriteErrorAdapter(write func(b []byte) (int, error)) func(data []byte) {
	return func(data []byte) {
		_, err := write(data)
		services.MustNot(err)
	}
}

func FailOnPrintErrorAdapter(cb func(a ...interface{}) (n int, err error)) MessageConsumer {
	return func(data []byte) {
		_, err := cb(data)
		services.MustNot(err)
	}
}

func FailOnWritePressureAdapter(write func(b []byte) (int, error)) PressureMessageConsumer {
	return func(message m.PressureMessage) {
		_, err := write(message)
		services.MustNot(err)
	}
}
