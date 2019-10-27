package docopt

import (
	"regexp"
	"strings"
	"unicode"
)

type tokenList struct {
	tokens    []string
	errorFunc func(string, ...interface{}) error
	err       errorType
}
type token string

func newTokenList(source []string, err errorType) *tokenList {
	errorFunc := newError
	if err == errorUser {
		errorFunc = newUserError
	} else if err == errorLanguage {
		errorFunc = newLanguageError
	}
	return &tokenList{source, errorFunc, err}
}

func tokenListFromString(source string) *tokenList {
	return newTokenList(strings.Fields(source), errorUser)
}

func tokenListFromPattern(source string) *tokenList {
	p := regexp.MustCompile(`([\[\]\(\)\|]|\.\.\.)`)
	source = p.ReplaceAllString(source, ` $1 `)
	p = regexp.MustCompile(`\s+|(\S*<.*?>)`)
	split := p.Split(source, -1)
	match := p.FindAllStringSubmatch(source, -1)
	var result []string
	l := len(split)
	for i := 0; i < l; i++ {
		if len(split[i]) > 0 {
			result = append(result, split[i])
		}
		if i < l-1 && len(match[i][1]) > 0 {
			result = append(result, match[i][1])
		}
	}
	return newTokenList(result, errorLanguage)
}

func (t *token) eq(s string) bool {
	if t == nil {
		return false
	}
	return string(*t) == s
}
func (t *token) match(matchNil bool, tokenStrings ...string) bool {
	if t == nil && matchNil {
		return true
	} else if t == nil && !matchNil {
		return false
	}

	for _, tok := range tokenStrings {
		if tok == string(*t) {
			return true
		}
	}
	return false
}
func (t *token) hasPrefix(prefix string) bool {
	if t == nil {
		return false
	}
	return strings.HasPrefix(string(*t), prefix)
}
func (t *token) hasSuffix(suffix string) bool {
	if t == nil {
		return false
	}
	return strings.HasSuffix(string(*t), suffix)
}
func (t *token) isUpper() bool {
	if t == nil {
		return false
	}
	return isStringUppercase(string(*t))
}
func (t *token) String() string {
	if t == nil {
		return ""
	}
	return string(*t)
}

func (tl *tokenList) current() *token {
	if len(tl.tokens) > 0 {
		return (*token)(&(tl.tokens[0]))
	}
	return nil
}

func (tl *tokenList) length() int {
	return len(tl.tokens)
}

func (tl *tokenList) move() *token {
	if len(tl.tokens) > 0 {
		t := tl.tokens[0]
		tl.tokens = tl.tokens[1:]
		return (*token)(&t)
	}
	return nil
}

// returns true if all cased characters in the string are uppercase
// and there are there is at least one cased charcter
func isStringUppercase(s string) bool {
	if strings.ToUpper(s) != s {
		return false
	}
	for _, c := range []rune(s) {
		if unicode.IsUpper(c) {
			return true
		}
	}
	return false
}
