package processor

func isBlank(line string) bool {

	for i := 0; i < len(line); i++ {

		char := byte(line[i])
		if char != '\n' && char != ' ' && char != '\t' {
			return false
		}
	}

	return true
}

func endsWithTwoOrMoreSpaces(line string) bool {

	spaceCount := 0
	for i := len(line) - 1; i >= 0; i-- {

		char := rune(line[i])
		if char != ' ' {
			break
		}
		spaceCount++
	}

	return (spaceCount >= 2)
}

func isAtxStyleHeader(line string) bool {

	return len(line) > 0 && byte(line[0]) == '#'
}

func checkForSeTextStyleHeader(line string) int {

	level := 0

	for i := 0; i < len(line); i++ {

		char := byte(line[i])

		if char == '=' && level != 2 {
			level = 1
		} else if char == '-' && level != 2 {
			level = 2
		} else {
			return 0
		}
	}
	return level
}
