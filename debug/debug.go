package debug

import (
	"log"
)

var Enabled bool

func D(message string) {
	if Enabled {
		log.Println(message)
	}
}
