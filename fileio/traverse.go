package fileio

// TraverseAndRun will recursively traverse "source" if it is a directory.
// When "source" is a file, it will run the "do" function on the file and save the
// returned bytes as a file in the "destination".
func TraverseAndRun(source string, destination string, do func(source string) []byte) {

	if isDirectory(source) {

		source := appendSeparator(source)
		destination := appendSeparator(destination)
		dirName := extractName(source)
		contents := getDirectoryListing(source)
		makeDirectory(destination, dirName)

		for _, element := range contents {
			TraverseAndRun(source+element, destination+dirName, do)
		}

		return
	}

	stream := do(source)
	writeToDisk(destination+extractName(source), stream)
}
