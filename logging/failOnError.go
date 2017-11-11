package logging

import (
	"fmt"
	"log"
)

//FailOnError prints error and message then kills process
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
