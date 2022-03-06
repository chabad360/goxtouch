package goxtouch

// A Channel is used to define which Fader, Meter, Encoder or LCD should be modified.
type Channel uint8

const (
	Channel1 Channel = iota + 1
	Channel2
	Channel3
	Channel4
	Channel5
	Channel6
	Channel7
	Channel8
	// Master is only a fader and will do nothing if used to set an Encoder or a Meter.
	Master

	LenChannels = 9

	ValueMax = 127
	ValueMin = 0
)

var (
	Faders = map[Channel]uint8{
		1: 70,
		2: 71,
		3: 72,
		4: 73,
		5: 74,
		6: 75,
		7: 76,
		8: 77,
		9: 78,
	}
	Encoders = map[Channel]uint8{
		1: 80,
		2: 81,
		3: 82,
		4: 83,
		5: 84,
		6: 85,
		7: 86,
		8: 87,
	}
	Meters = map[Channel]uint8{
		1: 90,
		2: 91,
		3: 92,
		4: 93,
		5: 94,
		6: 95,
		7: 96,
		8: 97,
	}
	LCDs = map[Channel]byte{
		1: 0x00,
		2: 0x01,
		3: 0x02,
		4: 0x03,
		5: 0x04,
		6: 0x05,
		7: 0x06,
		8: 0x07,
	}
)
