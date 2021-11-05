package main

import (
	"os"

	"github.com/viveknathani/makesite/converter"
	"github.com/viveknathani/makesite/fileio"
)

func main() {

	argLength := len(os.Args)
	if argLength == 1 || argLength != 3 {
		return
	}

	source := os.Args[1]
	destination := os.Args[2]
	fileio.TraverseAndRun(source, destination, converter.ToHTMLIfDoable)
}
