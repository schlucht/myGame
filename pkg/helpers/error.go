package helpers

import "log"

func ErrHandler(desc string, err error) {
	if err != nil {
		log.Fatalf("%s\n, %v", desc, err)
	}
}
