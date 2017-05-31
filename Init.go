package main

import (
	"log"
)

func init() {
	err := DatabaseInit()
	if err != nil {
		log.Fatalln(err)
	}
}
