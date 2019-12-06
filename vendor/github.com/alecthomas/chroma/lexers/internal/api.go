// Package internal contains common API functions and structures shared between lexer packages.
package internal

import (
	"path/filepath"
	"sort"
	"strings"

	"github.com/danwakefield/fnmatch"

	"github.com/alecthomas/chroma"
)

// Registry of Lexers.
var Registry = struct {
	Lexers  chroma.Lexers
	byName  map[string]chroma.Lexer
	byAlias map[string]chroma.Lexer
}{
	byName:  map[string]chroma.Lexer{},
	byAlias: map[string]chroma.Lexer{},
}

// Names of all lexers, optionally including aliases.
func Names(withAliases bool) []string {
	out := []string{}
	for _, lexer := range Registry.Lexers {
		config := lexer.Config()
		out = append(out, config.Name)
		if withAliases {
			out = append(out, config.Aliases...)
		}
	}
	sort.Strings(out)
	return out
}

// Get a Lexer by name, alias or file extension.
func Get(name string) chroma.Lexer {
	candidates := chroma.PrioritisedLexers{}
	if lexer := Registry.byName[name]; lexer != nil {
		candidates = append(candidates, lexer)
	}
	if lexer := Registry.byAlias[name]; lexer != nil {
		candidates = append(candidates, lexer)
	}
	if lexer := Registry.byName[strings.ToLower(name)]; lexer != nil {
		candidates = append(candidates, lexer)
	}
	if lexer := Registry.byAlias[strings.ToLower(name)]; lexer != nil {
		candidates = append(candidates, lexer)
	}
	// Try file extension.
	if lexer := Match("filename." + name); lexer != nil {
		candidates = append(candidates, lexer)
	}
	// Try exact filename.
	if lexer := Match(name); lexer != nil {
		candidates = append(candidates, lexer)
	}
	if len(candidates) == 0 {
		return nil
	}
	sort.Sort(candidates)
	return candidates[0]
}

// MatchMimeType attempts to find a lexer for the given MIME type.
func MatchMimeType(mimeType string) chroma.Lexer {
	matched := chroma.PrioritisedLexers{}
	for _, l := range Registry.Lexers {
		for _, lmt := range l.Config().MimeTypes {
			if mimeType == lmt {
				matched = append(matched, l)
			}
		}
	}
	if len(matched) != 0 {
		sort.Sort(matched)
		return matched[0]
	}
	return nil
}

// Match returns the first lexer matching filename.
func Match(filename string) chroma.Lexer {
	filename = filepath.Base(filename)
	matched := chroma.PrioritisedLexers{}
	// First, try primary filename matches.
	for _, lexer := range Registry.Lexers {
		config := lexer.Config()
		for _, glob := range config.Filenames {
			if fnmatch.Match(glob, filename, 0) {
				matched = append(matched, lexer)
			}
		}
	}
	if len(matched) > 0 {
		sort.Sort(matched)
		return matched[0]
	}
	matched = nil
	// Next, try filename aliases.
	for _, lexer := range Registry.Lexers {
		config := lexer.Config()
		for _, glob := range config.AliasFilenames {
			if fnmatch.Match(glob, filename, 0) {
				matched = append(matched, lexer)
			}
		}
	}
	if len(matched) > 0 {
		sort.Sort(matched)
		return matched[0]
	}
	return nil
}

// Analyse text content and return the "best" lexer..
func Analyse(text string) chroma.Lexer {
	var picked chroma.Lexer
	highest := float32(0.0)
	for _, lexer := range Registry.Lexers {
		if analyser, ok := lexer.(chroma.Analyser); ok {
			weight := analyser.AnalyseText(text)
			if weight > highest {
				picked = lexer
				highest = weight
			}
		}
	}
	return picked
}

// Register a Lexer with the global registry.
func Register(lexer chroma.Lexer) chroma.Lexer {
	config := lexer.Config()
	Registry.byName[config.Name] = lexer
	Registry.byName[strings.ToLower(config.Name)] = lexer
	for _, alias := range config.Aliases {
		Registry.byAlias[alias] = lexer
		Registry.byAlias[strings.ToLower(alias)] = lexer
	}
	Registry.Lexers = append(Registry.Lexers, lexer)
	return lexer
}

// Used for the fallback lexer as well as the explicit plaintext lexer
var PlaintextRules = chroma.Rules{
	"root": []chroma.Rule{
		{`.+`, chroma.Text, nil},
		{`\n`, chroma.Text, nil},
	},
}

// Fallback lexer if no other is found.
var Fallback chroma.Lexer = chroma.MustNewLexer(&chroma.Config{
	Name:      "fallback",
	Filenames: []string{"*"},
}, PlaintextRules)
