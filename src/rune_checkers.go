package main

func rIsDelimiter(r rune) bool {
	delimiters := []rune{
		' ', '\t', '\n', // Whitespace characters
		'(', ')', // Parentheses
		'[', ']', // Square brackets
		'{', '}', // Curly braces
		'.', ',', ':', ';', // Punctuation marks
	}
	return matchesAny(r, delimiters)
}

func rIsInteger(r rune) bool {
	integers := []rune{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	}
	return matchesAny(r, integers)
}

func rIsDash(r rune) bool {
	return r == '-'
}

func rIsNewline(r rune) bool {
	return r == '\n'
}

func rIsNumComponent(r rune) bool {
	numComponents := []rune{
		'A', 'a', 'B', 'b', 'C', 'c', 'D', 'd', 'E', 'e', 'F', 'f', 'X', 'x',
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'+', '-', '.',
	}
	return matchesAny(r, numComponents)
}

func rIsWhiteSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n'
}

func rIsQuoteMark(r rune) bool {
	return r == '"' || r == '\''
}
