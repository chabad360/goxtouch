package goxtouch

import (
	"strings"
	"time"

	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/midimessage/channel"
	"gitlab.com/gomidi/midi/midimessage/sysex"
	"gitlab.com/gomidi/midi/writer"
)

var (
	header = []byte{0x00, 0x20, 0x32, 0x14}

	SysExMessages = map[string][]byte{
		"Query":       {0x00},
		"GoOffline":   {0x0F, 0x7F},
		"Version":     {0x13, 0x00},
		"ResetFaders": {0x61},
		"ResetLEDs":   {0x62},
		"Reset":       {0x63},
	}
)

// Reset resets and quickly triggers all the available features on the control surface.
// I recommend running this both to avoid some errors and as a sanity check to make sure that your entire control surface is working.
func Reset(wd *writer.Writer) {
	var m []midi.Message
	for i := 0; i < len(IDs); i++ {
		m = append(m, SetLED(Switch(i), StateOn))
	}
	for i := 0; i <= LenChannels; i++ {
		m = append(m, SetFaderPos(Channel(i), 127))
	}
	for i := 0; i <= LenChannels-1; i++ {
		m = append(m, SetEncoder(Channel(i), 127))
	}
	for i := 0; i <= LenChannels-1; i++ {
		m = append(m, SetMeter(Channel(i), MeterLevel8))
	}
	// for i := 0; i < LenDigits; i++ {
	// 	m = append(m, SetDigit(Digit(uint8(i)+0x40), Char0+DigitDot))
	// }
	for i := 0; i <= LenChannels-1; i++ {
		m = append(m, SetLCD(Channel(i), byte(ColorRed+InvertLine1+InvertLine2), "test123", "123test"))
	}

	writer.WriteMessages(wd, m)

	time.Sleep(time.Second)
	m = []midi.Message{}

	for i := 0; i < len(IDs); i++ {
		m = append(m, SetLED(Switch(i), StateOff))
	}
	for i := 0; i <= LenChannels; i++ {
		m = append(m, SetFaderPos(Channel(i), 0))
	}
	for i := 0; i <= LenChannels-1; i++ {
		m = append(m, SetEncoder(Channel(i), 0))
	}
	for i := 0; i <= LenChannels-1; i++ {
		m = append(m, SetMeter(Channel(i), MeterLevel0))
	}
	// for i := 0; i < LenDigits; i++ {
	// 	m = append(m, SetDigit(Digit(uint8(i)+0x40), SymbolSpace))
	// }
	for i := 0; i <= LenChannels-1; i++ {
		m = append(m, SetLCD(Channel(i), byte(ColorNone), "       ", "       "))
	}
	writer.WriteMessages(wd, m)
}

// SetLED sets the chosen button's LED to the chosen State.
func SetLED(led Switch, state State) midi.Message {
	return channel.Channel(0).NoteOn(uint8(led), byte(state))
}

// SetFaderPos sets the position of the chosen fader to a number between 0 (bottom) and 16382 (top).
func SetFaderPos(fader Channel, pos uint8) midi.Message {
	return channel.Channel(0).ControlChange(Faders[fader], pos)
}

// SetTimeDisplay sets multiple characters on the timecode display.
// Note: letters is limited to ten characters and is right aligned.
// Refer to timecode.Digit for valid characters.
func SetTimeDisplay(letters string) (m []midi.Message) {
	bytes := []byte(strings.ToUpper(letters))
	if len(bytes) > 10 {
		bytes = bytes[:10]
	}

	for i, char := range bytes {
		if char >= 0x40 && char <= 0x60 {
			bytes[i] = char - 0x40
		}
	}

	for i := len(bytes)/2 - 1; i >= 0; i-- {
		opp := len(bytes) - 1 - i
		bytes[i], bytes[opp] = bytes[opp], bytes[i]
	}

	for i := uint8(0); int(i) < len(bytes); i++ {
		m = append(m, channel.Channel(15).ControlChange(i+0x40, bytes[i]))
	}
	return
}

// SetDigit sets an individual digit on the timecode or Assignment section.
func SetDigits(char1, char2, char3, char4, char5, char6, char7, char8, char9, char10, char11, char12, dots1, dots2 byte) midi.Message {
	return sysex.SysEx(append(header, 0x37, char1, char2, char3, char4, char5, char6, char7, char8, char9, char10, char11, char12, dots1, dots2))
}

// SetLCD sets the text (an ASCII string) found on the LCD starting from the specified offset.
func SetLCD(lcd Channel, colorData byte, line1, line2 string) midi.Message {
	if len(line1+line2) != 14 {
		return nil
	}
	return sysex.SysEx(append(append(header, 0x4c, LCDs[lcd], colorData), line1+line2...))
}

// SetEncoder sets the postion of the Encoders.
func SetEncoder(ch Channel, pos uint8) midi.Message {
	return channel.Channel(0).ControlChange(Encoders[ch], pos)
}

// SetMeter sets the level meter for the selected Channel to the desired value.
func SetMeter(ch Channel, pos uint8) midi.Message {
	return channel.Channel(0).ControlChange(Meters[ch], pos)
}
