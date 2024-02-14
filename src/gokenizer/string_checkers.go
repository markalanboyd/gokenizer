package gokenizer

import "regexp"

func isKeyword(s string) bool {
	keywords := []string{
		"and", "break", "do", "else", "elseif",
		"end", "for", "function", "goto", "if",
		"in", "local", "nil", "not", "or", "repeat",
		"return", "then", "until", "while",
	}
	return matchesAny(s, keywords)
}

func isParenthesis(s string) bool {
	parens := []string{
		"(", ")",
	}
	return matchesAny(s, parens)
}

// func isUnaryOperator(s string) bool {

// }

func isBracket(s string) bool {
	brackets := []string{
		"[", "]",
	}
	return matchesAny(s, brackets)
}

func isMultiLineCommentOpen(s string, i int) bool {
	if i+4 <= len(s) {
		return s[i:i+4] == "--[["
	}
	return false
}

func isMultiLineCommentClose(s string, i int) bool {
	return s[i-1:i+1] == "]]"
}

func isBrace(s string) bool {
	braces := []string{
		"{", "}",
	}
	return matchesAny(s, braces)
}

func isDelimiter(s string) bool {
	delimiters := []string{
		".", ",", ";", ":",
	}
	return matchesAny(s, delimiters)
}

func isConcatOperator(s string, i int) bool {
	if i+3 <= len(s) {
		return s[i:i+2] == ".."
	}
	return false
}

func isStringLiteral(s string) bool {
	r := []rune(s)
	return r[0] == '"' || r[0] == '\''
}

func isNumericLiteral(s string) bool {
	reNumericLiteral := regexp.MustCompile(`(?:0[xX][\dA-Fa-f]+|0[bB][01]+|\d+(\.\d+)?([eE][+-]?\d+)?)`)
	return reNumericLiteral.MatchString(s)
}

func isBooleanLiteral(s string) bool {
	return s == "true" || s == "false"
}

func isBinaryOperator(s string) bool {
	binaryOperators := []string{
		"+", "-", "*", "/", "^", "%", "..",
		"<", "<=", ">", ">=", "==", "~=",
		"and", "or",
	}
	return matchesAny(s, binaryOperators)
}

func isAssignmentOperator(s string) bool {
	return s == "="
}

func isIdentifier(s string) bool {
	reIdentifier := regexp.MustCompile(`\w+`)
	return reIdentifier.MatchString(s)
}
