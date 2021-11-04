package fileio

import "log"

func handleError(content string, err error) {

	if err != nil {
		log.Fatal(content, err)
	}
}
