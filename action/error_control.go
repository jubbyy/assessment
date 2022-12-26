package action

import (
	"log"
)

func HasErr(errorcode int, e any) {
	if e != nil {
		log.Printf("%v %v\n", errorcode, e)
		panic(e)
	}
}
