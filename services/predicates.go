package services

import "log"

func MustNot(err interface{}) {
	if err != nil {
		log.Fatal(err)
	}
}
