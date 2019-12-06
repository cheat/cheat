package j

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// JSON lexer.
var JSON = internal.Register(MustNewLexer(
	&Config{
		Name:         "JSON",
		Aliases:      []string{"json"},
		Filenames:    []string{"*.json"},
		MimeTypes:    []string{"application/json"},
		NotMultiline: true,
		DotAll:       true,
	},
	Rules{
		"whitespace": {
			{`\s+`, Text, nil},
		},
		"simplevalue": {
			{`(true|false|null)\b`, KeywordConstant, nil},
			{`-?(0|[1-9]\d*)(\.\d+[eE](\+|-)?\d+|[eE](\+|-)?\d+|\.\d+)`, LiteralNumberFloat, nil},
			{`-?(0|[1-9]\d*)`, LiteralNumberInteger, nil},
			{`"(\\\\|\\"|[^"])*"`, LiteralStringDouble, nil},
		},
		"objectattribute": {
			Include("value"),
			{`:`, Punctuation, nil},
			{`,`, Punctuation, Pop(1)},
			{`\}`, Punctuation, Pop(2)},
		},
		"objectvalue": {
			Include("whitespace"),
			{`"(\\\\|\\"|[^"])*"`, NameTag, Push("objectattribute")},
			{`\}`, Punctuation, Pop(1)},
		},
		"arrayvalue": {
			Include("whitespace"),
			Include("value"),
			{`,`, Punctuation, nil},
			{`\]`, Punctuation, Pop(1)},
		},
		"value": {
			Include("whitespace"),
			Include("simplevalue"),
			{`\{`, Punctuation, Push("objectvalue")},
			{`\[`, Punctuation, Push("arrayvalue")},
		},
		"root": {
			Include("value"),
		},
	},
))
