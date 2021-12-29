package processor

import "bytes"

type Tag struct {
	Name          string
	IsSelfClosing bool
	InnerText     string
	children      []*Tag
	attributes    map[string]string
}

func (tag *Tag) produceHTML() []byte {

	var output bytes.Buffer

	if tag.Name == "html" {
		output.WriteString("<!DOCTYPE html>")
	}
	output.WriteString("<" + tag.Name)

	if tag.attributes != nil {
		for key, value := range tag.attributes {
			output.WriteString(" " + key + "=" + "\"" + value + "\"")
		}
	}

	if tag.IsSelfClosing {
		output.WriteString("/>")
		return output.Bytes()
	}

	output.WriteString(">")
	output.WriteString(tag.InnerText)

	if tag.children != nil {
		for _, child := range tag.children {
			output.WriteString(string(child.produceHTML()))
		}
	}

	output.WriteString("</" + tag.Name + ">")
	return output.Bytes()
}
