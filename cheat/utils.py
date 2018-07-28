from __future__ import print_function
import os
import sys
import subprocess
import re

def colorize(sheet_content):
    """ Colorizes cheatsheet content if so configured """

    # only colorize if so configured
    if not 'CHEATCOLORS' in os.environ:
        return sheet_content

    try:
        from pygments import highlight
        from pygments.lexers import get_lexer_by_name
        from pygments.formatters import TerminalFormatter

    # if pygments can't load, just return the uncolorized text
    except ImportError:
        return sheet_content

    results = re.finditer(r"^(([ \t]*`{3})(?P<lexer>[^\n]*)(?P<code>[\s\S]+?)(^[ \t]*\2))", sheet_content, re.M)

    sheet_no_code = []  # list of slices of sheet_content without code blocks
    sheet_code = []  # list of slices of sheet_content with code blocks

    current = 0
    for result in results:
        try:
            lexer = get_lexer_by_name(result.group('lexer'))
            code = result.group('code')
            sheet_code.append(highlight(code, lexer, TerminalFormatter()).strip())
        except Exception:
            sheet_code.append(sheet_content[result.start():result.end()])

        sheet_no_code.append(sheet_content[current:result.start()])
        current = result.end()

    # Edge case, there aren't any blocks of code
    if current == 0:
        return sheet_content

    if current < len(sheet_content):
        sheet_no_code.append(sheet_content[current:])

    # Interleave the two list of slices together to rebuild the original document, but this time highlighted
    return "".join([val for pair in zip(sheet_no_code, sheet_code) for val in pair])


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
