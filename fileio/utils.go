package fileio

import "os"

const permissionRWRR = 0644
const permissionRWXRXRX = 0755

// ExtractName will get you the name of the file or the directory from the given
// variable 'path'
func ExtractName(path string) string {

	arr := make([]rune, 0)
	length := len(path)

	for i := length - 1; i >= 0; i-- {

		if path[i] == '/' && i != length-1 {
			break
		}

		arr = append(arr, rune(path[i]))
	}

	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}

	return string(arr)
}

// writeToDisk will write the given 'stream' of bytes to disk at the given 'path'
func writeToDisk(path string, stream []byte) {

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, permissionRWRR)
	handleError("writeToDisk: file read ", err)
	defer file.Close()

	_, err = file.Write(stream)
	handleError("writeToDisk: file write ", err)
}

// getDirectoryListing is basically your 'ls' command for the given directory
func getDirectoryListing(path string) []string {
	entries, err := os.ReadDir(path)
	handleError("getDirectoryListing: ", err)
	var list []string = make([]string, 0)
	for _, entry := range entries {
		list = append(list, entry.Name())
	}
	return list
}

// makeDirectory will make a 'name' directory under the given 'path'
func makeDirectory(path string, name string) {

	slash := ""
	if path[len(path)-1] != '/' {
		slash = "/"
	}
	dirPath := path + slash + name

	if _, er := os.Stat(dirPath); os.IsNotExist(er) {
		err := os.Mkdir(dirPath, permissionRWXRXRX)
		handleError("makeDirectory: ", err)
	}
}

// isDirectory will tell you if the given 'path' is a directory
func isDirectory(path string) bool {

	info, err := os.Stat(path)
	handleError("isDirectory: ", err)
	return info.IsDir()
}

// appendSeparator will add '/' to the end of the 'source' if it does not exist
func appendSeparator(source string) string {

	if len(source) != 0 && source[len(source)-1] != '/' {
		source += "/"
	}
	return source
}
