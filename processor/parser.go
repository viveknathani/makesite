package processor

import (
	"bytes"
	"strconv"
)

func parse(input []byte, justBody bool) *Tag {

	if justBody {
		return parseBody(input)
	}

	tag := new(Tag)
	tag.Name = "html"
	tag.attributes = nil

	head := parseHead(input)
	body := parseBody(input)

	tag.children = []*Tag{
		head, body,
	}

	return tag
}

func parseHead(input []byte) *Tag {

	tag := new(Tag)
	tag.Name = "head"
	tag.IsSelfClosing = false
	tag.attributes = nil

	meta := extractMeta(input)

	tag.children = []*Tag{
		getMetaCharset(),
		getMetaViewPort(),
		getTitleTag(meta.DOCUMENT_TITLE),
		getCSSTag(meta.CSS_URL),
	}

	return tag
}

func parseBody(input []byte) *Tag {

	tag := new(Tag)
	tag.Name = "body"

	children := make([]*Tag, 0)
	ps := bytes.Split(input, []byte("\n"))
	for i := 0; i < len(ps); i++ {

		p := ps[i]
		line := string(p)
		if !isBlank(line) {

			if i+1 < len(ps) {
				level := checkForSeTextStyleHeader(string(ps[i+1]))
				if level != 0 {
					children = append(children, parseSetextHeader(p, level))
					i++
					continue
				}
			}

			if isAtxStyleHeader(line) {
				children = append(children, parseAtxHeader(p))
				continue
			}

			children = append(children, parseParagraph(p))

			if endsWithTwoOrMoreSpaces(line) {
				children = append(children, getBRTag())
			}
		}
	}
	tag.children = children
	return tag
}

func parseSetextHeader(input []byte, level int) *Tag {

	return &Tag{
		Name:          "h" + strconv.Itoa(level),
		InnerText:     string(input),
		IsSelfClosing: false,
		attributes:    nil,
		children:      nil,
	}
}

func parseAtxHeader(input []byte) *Tag {

	level := 0

	start := -1
	end := len(input) - 1
	for i := 0; i < len(input); i++ {

		if level == 6 || input[i] != '#' {
			start = i
			break
		}
		level++
	}

	for i := len(input) - 1; i >= start; i-- {

		if input[i] != '#' {
			end = i
			break
		}
	}

	content := input[start : end+1]

	return &Tag{
		Name:          "h" + strconv.Itoa(level),
		InnerText:     string(content),
		IsSelfClosing: false,
		attributes:    nil,
		children:      nil,
	}
}

func parseParagraph(input []byte) *Tag {

	return &Tag{
		Name:          "p",
		InnerText:     string(input),
		IsSelfClosing: false,
		attributes:    nil,
		children:      nil,
	}
}

func getMetaCharset() *Tag {

	return &Tag{
		Name: "meta",
		attributes: map[string]string{
			"charset": "utf-8",
		},
		children:      nil,
		IsSelfClosing: true,
	}
}

func getMetaViewPort() *Tag {

	return &Tag{
		Name: "meta",
		attributes: map[string]string{
			"name":    "viewport",
			"content": "width=device-width, initial-scale=1",
		},
		children:      nil,
		IsSelfClosing: true,
	}
}

func getTitleTag(title string) *Tag {

	return &Tag{
		Name:          "title",
		InnerText:     title,
		IsSelfClosing: false,
		children:      nil,
		attributes:    nil,
	}
}

func getCSSTag(url string) *Tag {

	return &Tag{
		Name: "link",
		attributes: map[string]string{
			"rel":  "stylesheet",
			"type": "text/css",
			"href": url,
		},
		children:      nil,
		IsSelfClosing: true,
	}
}

func getBRTag() *Tag {

	return &Tag{
		Name:          "br",
		IsSelfClosing: true,
		attributes:    nil,
		children:      nil,
	}
}
