from __future__ import print_function
import os
import sys
import subprocess


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

    # determine which editor to use
    editor = os.environ.get('CHEAT_EDITOR') \
        or os.environ.get('VISUAL')         \
        or os.environ.get('EDITOR')         \
        or False

    # assert that the editor is set
    if editor == False:
        die(
            'You must set a CHEAT_EDITOR, VISUAL, or EDITOR environment '
            'variable in order to create/edit a cheatsheet.'
        )

    return editor


def open_with_editor(filepath):
    """ Open `filepath` using the EDITOR specified by the environment variables """
    editor_cmd = editor().split()
    try:
        subprocess.call(editor_cmd + [filepath])
    except OSError:
        die('Could not launch ' + editor())


def warn(message):
    """ Prints a message to stderr """
    print((message), file=sys.stderr)
