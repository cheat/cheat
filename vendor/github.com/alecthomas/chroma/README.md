# Chroma — A general purpose syntax highlighter in pure Go [![Golang Documentation](https://godoc.org/github.com/alecthomas/chroma?status.svg)](https://godoc.org/github.com/alecthomas/chroma) [![Build Status](https://travis-ci.org/alecthomas/chroma.svg)](https://travis-ci.org/alecthomas/chroma) [![Gitter chat](https://badges.gitter.im/alecthomas.svg)](https://gitter.im/alecthomas/Lobby)

> **NOTE:** As Chroma has just been released, its API is still in flux. That said, the high-level interface should not change significantly.

Chroma takes source code and other structured text and converts it into syntax
highlighted HTML, ANSI-coloured text, etc.

Chroma is based heavily on [Pygments](http://pygments.org/), and includes
translators for Pygments lexers and styles.

<a id="markdown-table-of-contents" name="table-of-contents"></a>
## Table of Contents

<!-- TOC -->

1. [Table of Contents](#table-of-contents)
2. [Supported languages](#supported-languages)
3. [Try it](#try-it)
4. [Using the library](#using-the-library)
    1. [Quick start](#quick-start)
    2. [Identifying the language](#identifying-the-language)
    3. [Formatting the output](#formatting-the-output)
    4. [The HTML formatter](#the-html-formatter)
5. [More detail](#more-detail)
    1. [Lexers](#lexers)
    2. [Formatters](#formatters)
    3. [Styles](#styles)
6. [Command-line interface](#command-line-interface)
7. [What's missing compared to Pygments?](#whats-missing-compared-to-pygments)

<!-- /TOC -->

<a id="markdown-supported-languages" name="supported-languages"></a>
## Supported languages

Prefix | Language
:----: | --------
A | ABAP, ABNF, ActionScript, ActionScript 3, Ada, Angular2, ANTLR, ApacheConf, APL, AppleScript, Arduino, Awk
B | Ballerina, Base Makefile, Bash, Batchfile, BlitzBasic, BNF, Brainfuck
C | C, C#, C++, Cap'n Proto, Cassandra CQL, Ceylon, CFEngine3, cfstatement, ChaiScript, Cheetah, Clojure, CMake, COBOL, CoffeeScript, Common Lisp, Coq, Crystal, CSS, Cython
D | D, Dart, Diff, Django/Jinja, Docker, DTD
E | EBNF, Elixir, Elm, EmacsLisp, Erlang
F | Factor, Fish, Forth, Fortran, FSharp
G | GAS, GDScript, Genshi, Genshi HTML, Genshi Text, GLSL, Gnuplot, Go, Go HTML Template, Go Text Template, GraphQL, Groovy
H | Handlebars, Haskell, Haxe, HCL, Hexdump, HTML, HTTP, Hy
I | Idris, INI, Io
J | J, Java, JavaScript, JSON, Julia, Jungle
K | Kotlin
L | Lighttpd configuration file, LLVM, Lua
M | Mako, markdown, Mason, Mathematica, Matlab, MiniZinc, MLIR, Modula-2, MonkeyC, MorrowindScript, Myghty, MySQL
N | NASM, Newspeak, Nginx configuration file, Nim, Nix
O | Objective-C, OCaml, Octave, OpenSCAD, Org Mode
P | PacmanConf, Perl, PHP, Pig, PkgConfig, PL/pgSQL, plaintext, PostgreSQL SQL dialect, PostScript, POVRay, PowerShell, Prolog, Protocol Buffer, Puppet, Python, Python 3
Q | QBasic
R | R, Racket, Ragel, react, reg, reStructuredText, Rexx, Ruby, Rust
S | Sass, Scala, Scheme, Scilab, SCSS, Smalltalk, Smarty, SML, Snobol, Solidity, SPARQL, SQL, SquidConf, Swift, SYSTEMD, systemverilog
T | TableGen, TASM, Tcl, Tcsh, Termcap, Terminfo, Terraform, TeX, Thrift, TOML, TradingView, Transact-SQL, Turing, Turtle, Twig, TypeScript, TypoScript, TypoScriptCssData, TypoScriptHtmlData
V | VB.net, verilog, VHDL, VimL, vue
W | WDTE
X | XML, Xorg
Y | YAML


_I will attempt to keep this section up to date, but an authoritative list can be
displayed with `chroma --list`._

<a id="markdown-try-it" name="try-it"></a>
## Try it

Try out various languages and styles on the [Chroma Playground](https://swapoff.org/chroma/playground/).

<a id="markdown-using-the-library" name="using-the-library"></a>
## Using the library

Chroma, like Pygments, has the concepts of
[lexers](https://github.com/alecthomas/chroma/tree/master/lexers),
[formatters](https://github.com/alecthomas/chroma/tree/master/formatters) and
[styles](https://github.com/alecthomas/chroma/tree/master/styles).

Lexers convert source text into a stream of tokens, styles specify how token
types are mapped to colours, and formatters convert tokens and styles into
formatted output.

A package exists for each of these, containing a global `Registry` variable
with all of the registered implementations. There are also helper functions
for using the registry in each package, such as looking up lexers by name or
matching filenames, etc.

In all cases, if a lexer, formatter or style can not be determined, `nil` will
be returned. In this situation you may want to default to the `Fallback`
value in each respective package, which provides sane defaults.

<a id="markdown-quick-start" name="quick-start"></a>
### Quick start

A convenience function exists that can be used to simply format some source
text, without any effort:

```go
err := quick.Highlight(os.Stdout, someSourceCode, "go", "html", "monokai")
```

<a id="markdown-identifying-the-language" name="identifying-the-language"></a>
### Identifying the language

To highlight code, you'll first have to identify what language the code is
written in. There are three primary ways to do that:

1. Detect the language from its filename.

    ```go
    lexer := lexers.Match("foo.go")
    ```

3. Explicitly specify the language by its Chroma syntax ID (a full list is available from `lexers.Names()`).

    ```go
    lexer := lexers.Get("go")
    ```

3. Detect the language from its content.

    ```go
    lexer := lexers.Analyse("package main\n\nfunc main()\n{\n}\n")
    ```

In all cases, `nil` will be returned if the language can not be identified.

```go
if lexer == nil {
  lexer = lexers.Fallback
}
```

At this point, it should be noted that some lexers can be extremely chatty. To
mitigate this, you can use the coalescing lexer to coalesce runs of identical
token types into a single token:

```go
lexer = chroma.Coalesce(lexer)
```

<a id="markdown-formatting-the-output" name="formatting-the-output"></a>
### Formatting the output

Once a language is identified you will need to pick a formatter and a style (theme).

```go
style := styles.Get("swapoff")
if style == nil {
  style = styles.Fallback
}
formatter := formatters.Get("html")
if formatter == nil {
  formatter = formatters.Fallback
}
```

Then obtain an iterator over the tokens:

```go
contents, err := ioutil.ReadAll(r)
iterator, err := lexer.Tokenise(nil, string(contents))
```

And finally, format the tokens from the iterator:

```go
err := formatter.Format(w, style, iterator)
```

<a id="markdown-the-html-formatter" name="the-html-formatter"></a>
### The HTML formatter

By default the `html` registered formatter generates standalone HTML with
embedded CSS. More flexibility is available through the `formatters/html` package.

Firstly, the output generated by the formatter can be customised with the
following constructor options:

- `Standalone()` - generate standalone HTML with embedded CSS.
- `WithClasses()` - use classes rather than inlined style attributes.
- `ClassPrefix(prefix)` - prefix each generated CSS class.
- `TabWidth(width)` - Set the rendered tab width, in characters.
- `WithLineNumbers()` - Render line numbers (style with `LineNumbers`).
- `LinkableLineNumbers()` - Make the line numbers linkable.
- `HighlightLines(ranges)` - Highlight lines in these ranges (style with `LineHighlight`).
- `LineNumbersInTable()` - Use a table for formatting line numbers and code, rather than spans.

If `WithClasses()` is used, the corresponding CSS can be obtained from the formatter with:

```go
formatter := html.New(html.WithClasses())
err := formatter.WriteCSS(w, style)
```

<a id="markdown-more-detail" name="more-detail"></a>
## More detail

<a id="markdown-lexers" name="lexers"></a>
### Lexers

See the [Pygments documentation](http://pygments.org/docs/lexerdevelopment/)
for details on implementing lexers. Most concepts apply directly to Chroma,
but see existing lexer implementations for real examples.

In many cases lexers can be automatically converted directly from Pygments by
using the included Python 3 script `pygments2chroma.py`. I use something like
the following:

```sh
python3 ~/Projects/chroma/_tools/pygments2chroma.py \
  pygments.lexers.jvm.KotlinLexer \
  > ~/Projects/chroma/lexers/kotlin.go \
  && gofmt -s -w ~/Projects/chroma/lexers/*.go
```

See notes in [pygments-lexers.go](https://github.com/alecthomas/chroma/blob/master/pygments-lexers.txt)
for a list of lexers, and notes on some of the issues importing them.

<a id="markdown-formatters" name="formatters"></a>
### Formatters

Chroma supports HTML output, as well as terminal output in 8 colour, 256 colour, and true-colour.

A `noop` formatter is included that outputs the token text only, and a `tokens`
formatter outputs raw tokens. The latter is useful for debugging lexers.

<a id="markdown-styles" name="styles"></a>
### Styles

Chroma styles use the [same syntax](http://pygments.org/docs/styles/) as Pygments.

All Pygments styles have been converted to Chroma using the `_tools/style.py` script.

When you work with one of [Chroma's styles](https://github.com/alecthomas/chroma/tree/master/styles), know that the `chroma.Background` token type provides the default style for tokens. It does so by defining a foreground color and background color.

For example, this gives each token name not defined in the style a default color of `#f8f8f8` and uses `#000000` for the highlighted code block's background:

~~~go
chroma.Background: "#f8f8f2 bg:#000000",
~~~

Also, token types in a style file are hierarchical. For instance, when `CommentSpecial` is not defined, Chroma uses the token style from `Comment`. So when several comment tokens use the same color, you'll only need to define `Comment` and override the one that has a different color.

For a quick overview of the available styles and how they look, check out the [Chroma Style Gallery](https://xyproto.github.io/splash/docs/).

<a id="markdown-command-line-interface" name="command-line-interface"></a>
## Command-line interface

A command-line interface to Chroma is included. It can be installed with:

```sh
go get -u github.com/alecthomas/chroma/cmd/chroma
```

<a id="markdown-whats-missing-compared-to-pygments" name="whats-missing-compared-to-pygments"></a>
## What's missing compared to Pygments?

- Quite a few lexers, for various reasons (pull-requests welcome):
    - Pygments lexers for complex languages often include custom code to
      handle certain aspects, such as Perl6's ability to nest code inside
      regular expressions. These require time and effort to convert.
    - I mostly only converted languages I had heard of, to reduce the porting cost.
- Some more esoteric features of Pygments are omitted for simplicity.
- Though the Chroma API supports content detection, very few languages support them.
  I have plans to implement a statistical analyser at some point, but not enough time.
