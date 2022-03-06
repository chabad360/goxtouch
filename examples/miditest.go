package main

import (
	"fmt"
	"time"

	xtouch "github.com/chabad360/goxtouch"
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/reader"

	"gitlab.com/gomidi/midi/writer"
	driver "gitlab.com/gomidi/rtmididrv"
)

func must(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// This example demonstrates using github.com/chabad360/goxtouch to send and receive signals using the CTRL mode interface for the Behringer X-Touch Universal
func main() {
	drv, err := driver.New()
	must(err)

	// make sure to close all open ports at the end
	defer drv.Close()

	ins, err := drv.Ins()
	must(err)

	outs, err := drv.Outs()
	must(err)

	fmt.Println(ins, outs)

	in, out := ins[0], outs[1]

	must(in.Open())
	must(out.Open())

	wr := writer.New(out)
	rd := reader.New(
		reader.NoteOn(noteon(wr)))
	xtouch.Reset(wr)

	go rd.ListenTo(in)

	m := []midi.Message{xtouch.SetDigits(xtouch.Char2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)}
	m = append(m, xtouch.SetLCD(xtouch.Channel1, byte(xtouch.ColorGreen), string([]byte{'c', 'h', '1', '2', 0x00, 0x00, 0x00}), "   -0dB"))
	m = append(m, xtouch.SetLED(109, xtouch.StateOn))
	writer.WriteMessages(wr, m)

	time.Sleep(time.Second * 20)
}

func noteon(wr *writer.Writer) func(p *reader.Position, c, k, v uint8) {
	return func(p *reader.Position, c, k, v uint8) {
		switch xtouch.Switch(k) {
		case xtouch.Play:
			wr.Write(xtouch.SetLED(xtouch.Play, xtouch.StateOn))
		}
	}
}
