package gokenizer

// Tokenize splits a Lua script into individual tokens with type and value information, ignoring comments.
func Tokenize(s string) []Token {
	return stringToTokens(s)
}
