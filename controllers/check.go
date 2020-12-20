package controllers

import "log"

func Check(err error) {
	if err != nil {
		log.Printf("%+v\n", err)
	}
}
