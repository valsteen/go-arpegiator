package devices

import (
	"gitlab.com/gomidi/midi"
	"go-arpegiator/definitions"
	"go-arpegiator/services"
)

type ChannelMessageConsumer func(message midiDefinitions.ChannelMessage)

func PipeRawMessageToChannelMessage(in midi.In, consumer ChannelMessageConsumer) {
	err := in.SetListener(func(data []byte, deltaMicroseconds int64) {
		midiMessage := midiDefinitions.AsMidiMessage(data)
		if midiChannelMessage, ok := midiMessage.(midiDefinitions.ChannelMessage); ok {
			consumer(midiChannelMessage)
		}
	})
	services.MustNot(err)
}

type NoteSetConsumer func(notes NoteSet)
