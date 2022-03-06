package goxtouch

// MeterLevel is the height (in decibels) that the meter should be set to.
type MeterLevel int

const (
	MeterLevel0 = iota * 14
	MeterLevel1
	MeterLevel2
	MeterLevel3
	MeterLevel4
	MeterLevel5 = iota * 15
	MeterLevel6
	MeterLevel7
	MeterLevel8
)
