package debug

import (
	"log"
)

var Enabled bool

func D(message string) {
	Enabled = true
	if Enabled {
		log.Println(message)
	}
}
