package goxtouch

type Color byte
type Invert byte

const (
	ColorNone Color = iota
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite

	NoInvert    = 0
	InvertLine1 = 0x10
	InvertLine2 = 0x20
)
