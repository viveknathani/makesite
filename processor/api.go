package processor

// ConvertMarkdownToHTML will take input buffer containing markdown text
// and produce the equivalent HTML
func ConvertMarkdownToHTML(input []byte, justBody bool) []byte {

	tag := parse(input, justBody)
	return tag.produceHTML()
}
