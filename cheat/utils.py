import os
import sys


def colorize(sheet_content):
    """ Colorizes cheatsheet content if so configured """

    # only colorize if so configured
    if not 'CHEATCOLORS' in os.environ:
        return sheet_content

    try:
        from pygments import highlight
        from pygments.lexers import BashLexer
        from pygments.formatters import TerminalFormatter

    # if pygments can't load, just return the uncolorized text
    except ImportError:
        return sheet_content

    return highlight(sheet_content, BashLexer(), TerminalFormatter())


def die(message):
    """ Prints a message to stderr and then terminates """
    warn(message)
    exit(1)


def editor():
    """ Determines the user's preferred editor """
    if 'EDITOR' not in os.environ:
        die(
            'In order to create/edit a cheatsheet you must set your EDITOR '
            'environment variable to your editor\'s path.'
        )

    elif os.environ['EDITOR'] == "":
        die(
          'Your EDITOR environment variable is set to an empty string. It must '
          'be set to your editor\'s path.'
        )

    else:
        return os.environ['EDITOR']


def prompt_yes_or_no(question):
    """ Prompts the user with a yes-or-no question """
    print(question)
    return raw_input('[y/n] ') == 'y'


def warn(message):
    """ Prints a message to stderr """
    print >> sys.stderr, (message)
