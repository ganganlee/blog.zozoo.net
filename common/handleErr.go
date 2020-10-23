package common

import (
	"log"
)

func HandleErr(err error,msg string)  {
	if err != nil {
		log.Fatalf("%v err:%V",msg,err.Error())
	}
}