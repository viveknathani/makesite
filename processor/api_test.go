package processor

import (
	"log"
	"testing"
)

func TestConvertMarkdownToHTML(t *testing.T) {

	testCases := []struct {
		input    string
		output   string
		justBody bool
	}{
		{"hello", "<body><p>hello</p></body>", true},
		{"hello\nhey", "<body><p>hello</p><p>hey</p></body>", true},
		{"hello\n\nhey", "<body><p>hello</p><p>hey</p></body>", true},
		{"hello \n\nhey", "<body><p>hello </p><p>hey</p></body>", true},
		{"hello  \n\nhey", "<body><p>hello  </p><br/><p>hey</p></body>", true},
		{"hello   \n\nhey", "<body><p>hello   </p><br/><p>hey</p></body>", true},
		{"#hello", "<body><h1>hello</h1></body>", true},
		{"##hello", "<body><h2>hello</h2></body>", true},
		{"###hello", "<body><h3>hello</h3></body>", true},
		{"####hello", "<body><h4>hello</h4></body>", true},
		{"#####hello", "<body><h5>hello</h5></body>", true},
		{"######hello", "<body><h6>hello</h6></body>", true},
		{"hello\n=", "<body><h1>hello</h1></body>", true},
		{"hello\n-", "<body><h2>hello</h2></body>", true},
	}

	for _, testCase := range testCases {

		got := string(ConvertMarkdownToHTML([]byte(testCase.input), testCase.justBody))

		if testCase.output != got {
			log.Fatalf("input: %s, expected: %s, got: %s", testCase.input, testCase.output, got)
		}
	}
}
