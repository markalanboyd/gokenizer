package gokenizer

import (
	"fmt"
)

// Token represents a lexical token, which consists of a type (T) and a corresponding value (V). The type (T) indicates the category or classification of the token, such as "Identifier", "Keyword", "BinaryOperator", "StringLiteral", etc. The value (V) provides the specific content or data associated with the token, such as the actual identifier name, keyword text, operator symbol, or string literal contents.
type Token struct {
	T string
	V string
}

// Generics

func matchesAny[T comparable](item T, slice []T) bool {
	for _, elem := range slice {
		if item == elem {
			return true
		}
	}
	return false
}

func lastOf[T any](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	return s[len(s)-1]
}

// Util Functions

func assignTokenType(tokens []Token) {
	for i := range tokens {
		var t string
		v := tokens[i].V
		switch {
		case isKeyword(v):
			t = "Keyword"
		case isParenthesis(v):
			t = "Parenthesis"
		case isDelimiter(v):
			t = "Delimiter"
		case isBracket(v):
			t = "Bracket"
		case isBrace(v):
			t = "Brace"
		case isAssignmentOperator(v):
			t = "AssignmentOperator"
		case isBinaryOperator(v):
			t = "BinaryOperator"
		case isBooleanLiteral(v):
			t = "BooleanLiteral"
		case isStringLiteral(v):
			t = "StringLiteral"
		case isNumericLiteral(v):
			t = "NumericLiteral"
		case isIdentifier(v):
			t = "Identifier"
		}
		tokens[i].T = t
	}
}

func stringToTokenVal(s string) []Token {
	tokens := []Token{}
	var rS []rune
	var insideStringLiteral, insideNumericLiteral, insideComment, insideMultiLineComment, skipNext bool
	var stringDelimiter rune

	for i, r := range s {
		switch {
		// STRING LITERAL //
		case skipNext:
			skipNext = false
		case !insideStringLiteral && !insideMultiLineComment && rIsQuoteMark(r):
			insideStringLiteral = true
			stringDelimiter = r
			rS = append(rS, r)
		case insideStringLiteral && r == stringDelimiter:
			insideStringLiteral = false
			rS = append(rS, r)
			tokens = append(tokens, Token{V: string(rS)})
			rS = rS[:0]
		case insideStringLiteral:
			rS = append(rS, r)

		// NUMERIC LITERAL //
		case !insideNumericLiteral && rIsInteger(r):
			insideNumericLiteral = true
			rS = append(rS, r)
		case insideNumericLiteral && !rIsNumComponent(r):
			insideNumericLiteral = false
			tokens = append(tokens, Token{V: string(rS)})
			rS = rS[:0]
			if rIsPunctuation(r) {
				tokens = append(tokens, Token{V: string(r)})
			}
		case insideNumericLiteral:
			rS = append(rS, r)

		// COMMENT //
		case !insideComment && rIsDash(r) && rIsDash(lastOf(rS)):
			insideComment = true
			rS = rS[:len(rS)-1]
			continue
		case insideComment && rIsNewline(r):
			insideComment = false
		case insideComment:
			continue

		// MULTI-LINE COMMENT //
		case !insideMultiLineComment && !insideStringLiteral && isMultiLineCommentOpen(s, i):
			insideMultiLineComment = true
		case insideMultiLineComment && isMultiLineCommentClose(s, i):
			insideMultiLineComment = false
		case insideMultiLineComment:
			continue

		// NORMAL TOKENIZATION AND DELIMITERS //
		default:
			switch {
			case rIsDelimiter(r):
				if isConcatOperator(s, i) {
					tokens = append(tokens, Token{V: ".."})
					skipNext = true
					continue
				}
				if len(rS) > 0 {
					tokens = append(tokens, Token{V: string(rS)})
					rS = rS[:0]
				}
				if rIsPunctuation(r) {
					tokens = append(tokens, Token{V: string(r)})
				}
			default:
				rS = append(rS, r)
			}
		}
	}

	// FINAL TOKEN //
	if len(rS) > 0 {
		tokens = append(tokens, Token{V: string(rS)})
	}

	return tokens
}

func stringToTokens(s string) []Token {
	tokens := stringToTokenVal(s)
	assignTokenType(tokens)
	return tokens
}

// PrintTokens pretty prints tokens to the console
func PrintTokens(tokens []Token) {
	for i, token := range tokens {
		pad := 19 - len(token.T)
		fmt.Printf("%v T: %s %*s V:%s\n", i, token.T, pad, "", token.V)
	}
}
