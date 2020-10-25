package devices

type switches []byte

type ArpInDevice struct {
	switches
}

func (d ArpInDevice) Consume(notes Notes) {
	// TODO
}

func NewArpInDevice() * ArpInDevice {
	return &ArpInDevice{
		switches: make(switches, 0, 12),
	}
}
