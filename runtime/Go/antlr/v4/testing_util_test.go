package antlr

import (
	"fmt"
	"strings"
)

// newTestCommonToken create common token with tokenType, text and channel
// notice: test purpose only
func newTestCommonToken(tokenType int, text string, channel int) *CommonToken {
	t := new(CommonToken)
	t.tokenType = tokenType
	t.channel = channel
	t.text = text
	t.line = 0
	t.column = -1
	return t
}

// tokensToString returns []Tokens string
// notice: test purpose only
func tokensToString(tokens []Token) string {
	buf := make([]string, len(tokens))
	for i, token := range tokens {
		buf[i] = fmt.Sprintf("%v", token)
	}
	
	return "[" + strings.Join(buf, ", ") + "]"
}
