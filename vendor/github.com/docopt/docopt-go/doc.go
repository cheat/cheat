/*
Package docopt parses command-line arguments based on a help message.

Given a conventional command-line help message, docopt processes the arguments.
See https://github.com/docopt/docopt#help-message-format for a description of
the help message format.

This package exposes three different APIs, depending on the level of control
required. The first, simplest way to parse your docopt usage is to just call:

	docopt.ParseDoc(usage)

This will use os.Args[1:] as the argv slice, and use the default parser
options. If you want to provide your own version string and args, then use:

	docopt.ParseArgs(usage, argv, "1.2.3")

If the last parameter (version) is a non-empty string, it will be printed when
--version is given in the argv slice. Finally, we can instantiate our own
docopt.Parser which gives us control over how things like help messages are
printed and whether to exit after displaying usage messages, etc.

	parser := &docopt.Parser{
		HelpHandler: docopt.PrintHelpOnly,
		OptionsFirst: true,
	}
	opts, err := parser.ParseArgs(usage, argv, "")

In particular, setting your own custom HelpHandler function makes unit testing
your own docs with example command line invocations much more enjoyable.

All three of these return a map of option names to the values parsed from argv,
and an error or nil. You can get the values using the helpers, or just treat it
as a regular map:

	flag, _ := opts.Bool("--flag")
	secs, _ := opts.Int("<seconds>")

Additionally, you can `Bind` these to a struct, assigning option values to the
exported fields of that struct, all at once.

	var config struct {
		Command string `docopt:"<cmd>"`
		Tries   int    `docopt:"-n"`
		Force   bool   // Gets the value of --force
	}
	opts.Bind(&config)
*/
package docopt
