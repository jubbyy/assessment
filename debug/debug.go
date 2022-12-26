package debug

import (
	"log"
)

var Enabled bool

func D(message string) {
	Enabled = false
	if Enabled {
		log.Println(message)
	}
}
