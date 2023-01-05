package model

type Configuration struct {
	Init       bool
	GinRelease bool
	Port       string
	Noweb      bool
	Iface      string
	VerboseLog bool
	Mock       bool
}
