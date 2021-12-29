package fileio

import (
	"io/ioutil"
	"os"
)

// TraverseAndRun will recursively traverse "source" if it is a directory.
// When "source" is a file, it will run the "do" function on the file and save the
// returned bytes as a file in the "destination".
func TraverseAndRun(source string, destination string, do func(data []byte, justBody bool) []byte) {

	if isDirectory(source) {

		source := appendSeparator(source)
		destination := appendSeparator(destination)
		dirName := ExtractName(source)
		contents := getDirectoryListing(source)
		makeDirectory(destination, dirName)

		for _, element := range contents {
			TraverseAndRun(source+element, destination+dirName, do)
		}

		return
	}

	file, err := os.Open(source)
	handleError("TraverseAndRun: file open ", err)

	data, err := ioutil.ReadAll(file)
	handleError("TraverseAndRun: file to bytes ", err)
	file.Close()

	if !isMarkdown(source) {

		writeToDisk(destination+ExtractName(source), data)
		return
	}

	stream := do(data, false)
	writeToDisk(destination+updateExtension(ExtractName(source)), stream)
}
