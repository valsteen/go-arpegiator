package devices

import (
	"fmt"
	"gitlab.com/gomidi/midi"
	midiDefinitions "go-arpegiator/definitions"
	"go-arpegiator/services"
)

type midiChannelMessageChan chan midiDefinitions.ChannelMessage
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

func (device *Device) consume(messageChan midiChannelMessageChan) {
	for message := range messageChan {
		if noteMessage, ok := message.(midiDefinitions.NoteMessage); ok {
			if noteMessage.IsNoteOn() {
				device.notes = device.notes.insert(noteMessage)
			} else {
				device.notes = device.notes.remove(noteMessage)
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
