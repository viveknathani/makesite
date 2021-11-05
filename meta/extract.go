package meta

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type MakeSiteMeta struct {
	CSS_URL        string
	DOCUMENT_TITLE string
}

const makeSiteMetaPrefix = "[meta]: # ("

// ExtractMeta will extract metadata from your markdown file and
// return it
func ExtractMeta(source string) *MakeSiteMeta {

	file, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	m := new(MakeSiteMeta)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	extracted := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, makeSiteMetaPrefix) {

			extracted = true

			startIndex := len(makeSiteMetaPrefix)
			endIndex := len(line) - 1

			valueArr := make([]rune, 0)
			keyArr := make([]rune, 0)
			okay := false
			for i := startIndex; i <= endIndex; i++ {
				if line[i] == ')' {
					break
				}
				if okay {
					valueArr = append(valueArr, rune(line[i]))
					continue
				}
				if line[i] == '=' {
					okay = true
					continue
				}
				keyArr = append(keyArr, rune(line[i]))
			}

			key := string(keyArr)
			value := string(valueArr)

			switch key {

			case "CSS_URL":
				m.CSS_URL = value
			case "DOCUMENT_TITLE":
				m.DOCUMENT_TITLE = value
			}
			continue
		}

		if extracted {
			break
		}
	}
	return m
}
