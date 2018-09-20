#!/usr/bin/env python

"""cheat

Create and view cheatsheets on the command line.

Usage:
  cheat <cheatsheet>
  cheat -e <cheatsheet>
  cheat -s <keyword>
  cheat -l
  cheat -d
  cheat -v

Options:
  -d --directories  List directories on CHEATPATH
  -e --edit         Edit cheatsheet
  -l --list         List cheatsheets
  -s --search       Search cheatsheets for <keyword>
  -v --version      Print the version number

Examples:

  To view the `tar` cheatsheet:
    cheat tar

  To edit (or create) the `foo` cheatsheet:
    cheat -e foo

  To list all available cheatsheets:
    cheat -l

  To search for "ssh" among all cheatsheets:
    cheat -s ssh
"""

# require the dependencies
from cheat import sheets, sheet
from cheat.utils import colorize
from docopt import docopt


if __name__ == '__main__':
    # parse the command-line options
    options = docopt(__doc__, version='cheat 2.2.3')

    # list directories
    if options['--directories']:
        print("\n".join(sheets.paths()))

    # list cheatsheets
    elif options['--list']:
        print(sheets.list())

    # create/edit cheatsheet
    elif options['--edit']:
        sheet.create_or_edit(options['<cheatsheet>'])

    # search among the cheatsheets
    elif options['--search']:
        print(colorize(sheets.search(options['<keyword>'])))

    # print the cheatsheet
    else:
        print(colorize(sheet.read(options['<cheatsheet>'])))
