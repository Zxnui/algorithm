package common

import (
	"log"
	"os"
)

func CheckError(err error) {
	if err != nil {
		log.Println(err.Error())
		os.Exit(2)
	}
}
