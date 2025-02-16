package helper

import (
	"log"
)

func HandleError(e error) {
	if e != nil {
		log.Println(e)
	}
}
