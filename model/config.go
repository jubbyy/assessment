package model

type Configuration struct {
	Init       bool
	GinRelease bool
	Port       string
	Action     string
	Iface      string
	VerboseLog bool
	Mock       bool
}
