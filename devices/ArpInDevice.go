package devices

type switches []byte

type ArpInDevice struct {
	switches
}

func (d ArpInDevice) Consume(notes Notes) {
	// TODO replace note array with set
	/*
	operations needed:
	- diff between current/previous state
	- obviously need a set structure
	- remove DIY operations and use go's set
	 */
}

func NewArpInDevice() * ArpInDevice {
	return &ArpInDevice{
		switches: make(switches, 0, 12),
	}
}
