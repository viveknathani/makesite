package converter

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/russross/blackfriday/v2"
	"github.com/viveknathani/makesite/fileio"
)

func isMarkdown(source string) bool {

	return strings.HasSuffix(source, ".md") || strings.HasSuffix(source, ".md/")
}

func updateExtension(name string) string {

	arr := make([]rune, 0)
	arr = append(arr, rune('l'))
	arr = append(arr, rune('m'))
	arr = append(arr, rune('t'))
	arr = append(arr, rune('h'))

	length := len(name)

	copy := false
	for i := length - 1; i >= 0; i-- {

		if name[i] == '.' && i != length-1 {
			copy = true
		}
		if copy {
			arr = append(arr, rune(name[i]))
		}
	}

	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}

	return string(arr)
}

// ToHTMLIfDoable will take the file in "source" and convert it into HTML
// if it is in Markdown format.
func ToHTMLIfDoable(source string) ([]byte, string) {

	file, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
	}

	stream, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	name := fileio.ExtractName(source)
	if isMarkdown(source) {

		params := blackfriday.HTMLRendererParameters{
			Flags: blackfriday.CompletePage,
		}
		renderer := blackfriday.NewHTMLRenderer(params)
		return blackfriday.Run(stream, blackfriday.WithRenderer(renderer)), updateExtension(name)
	}

	return stream, name
}
