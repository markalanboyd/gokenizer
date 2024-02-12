package main

import (
	"fmt"
)

// Types

type token struct {
	T   string
	Val string
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

func assignTokenType(tokens []token) {
	for i := range tokens {
		var t string
		v := tokens[i].Val
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

func stringToTokenVal(s string) []token {
	tokens := []token{}
	var rSlice []rune
	var insideStringLiteral, insideNumericLiteral, insideComment, insideMultiLineComment bool
	var stringDelimiter rune

	for i, r := range s {
		switch {
		// STRING LITERAL //
		case !insideStringLiteral && !insideMultiLineComment && rIsQuoteMark(r):
			insideStringLiteral = true
			stringDelimiter = r
			rSlice = append(rSlice, r)
		case insideStringLiteral && r == stringDelimiter:
			insideStringLiteral = false
			rSlice = append(rSlice, r)
			tokens = append(tokens, token{Val: string(rSlice)})
			rSlice = rSlice[:0]
		case insideStringLiteral:
			rSlice = append(rSlice, r)

		// NUMERIC LITERAL //
		case !insideNumericLiteral && rIsInteger(r):
			insideNumericLiteral = true
			rSlice = append(rSlice, r)
		case insideNumericLiteral && !rIsNumComponent(r):
			insideNumericLiteral = false
			tokens = append(tokens, token{Val: string(rSlice)})
			rSlice = rSlice[:0]
		case insideNumericLiteral:
			rSlice = append(rSlice, r)

		// SINGLE-LINE COMMENT //
		case !insideComment && rIsDash(r) && rIsDash(lastOf(rSlice)):
			insideComment = true
			rSlice = rSlice[:len(rSlice)-1]
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

		// TODO Fix concat op
		// NORMAL TOKENIZATION AND DELIMITERS //
		default:
			if rIsDelimiter(r) && len(rSlice) > 0 {
				tokens = append(tokens, token{Val: string(rSlice)})
				rSlice = rSlice[:0]
			}
			if !rIsWhiteSpace(r) && !rIsDelimiter(r) {
				rSlice = append(rSlice, r)
			} else if rIsDelimiter(r) && !rIsWhiteSpace(r) {
				tokens = append(tokens, token{Val: string(r)})
			}
		}
	}

	// FINAL TOKEN //
	if len(rSlice) > 0 {
		tokens = append(tokens, token{Val: string(rSlice)})
	}

	return tokens
}

func stringToTokens(s string) []token {
	tokens := stringToTokenVal(s)
	assignTokenType(tokens)
	return tokens
}

func printTokens(tokens []token) {
	for _, token := range tokens {
		fmt.Printf("%s: %s\n", token.T, token.Val)
	}
}
