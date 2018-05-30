from __future__ import print_function
import os
import subprocess
import sys

class Utils:


    def __init__(self,cheatcolors,editor_executable):
        self.displaycolors = cheatcolors
        self.editor_executable = editor_executable


    def highlight(self, needle, haystack):
        """ Highlights a search term matched within a line """

        # if a highlight color is not configured, exit early
        if not 'CHEAT_HIGHLIGHT' in os.environ:
            return haystack

        # otherwise, attempt to import the termcolor library
        try:
            from termcolor import colored

        # if the import fails, return uncolored text
        except ImportError:
            return haystack

        # if the import succeeds, colorize the needle in haystack
        return haystack.replace(needle, colored(needle, os.environ.get('CHEAT_HIGHLIGHT')))


    def colorize(self,sheet_content):
        """ Colorizes cheatsheet content if so configured """

        # only colorize if configured to do so, and if stdout is a tty
        if not self.displaycolors or not sys.stdout.isatty():
            return sheet_content

        # don't attempt to colorize an empty cheatsheet
        if not sheet_content.strip():
            return ""

        # otherwise, attempt to import the pygments library
        try:
            from pygments import highlight
            from pygments.lexers import get_lexer_by_name
            from pygments.formatters import TerminalFormatter

        # if the import fails, return uncolored text
        except ImportError:
            return sheet_content

        # otherwise, attempt to colorize
        first_line = sheet_content.splitlines()[0]
        lexer      = get_lexer_by_name('bash')

        # apply syntax-highlighting if the first line is a code-fence
        if first_line.startswith('```'):
            sheet_content = '\n'.join(sheet_content.split('\n')[1:-2])
            try:
                lexer = get_lexer_by_name(first_line[3:])
            except Exception:
                pass

        return highlight(sheet_content, lexer, TerminalFormatter())


    @staticmethod
    def die(message):
        """ Prints a message to stderr and then terminates """
        Utils.warn(message)
        exit(1)


    def editor(self):
        """ Determines the user's preferred editor """

        # assert that the editor is set
        if (not self.editor_executable):
            Utils.die(
                'You must set a CHEAT_EDITOR, VISUAL, or EDITOR environment '
                'variable or setting in order to create/edit a cheatsheet.'
            )

        return self.editor_executable


    def open_with_editor(self,filepath):
        """ Open `filepath` using the EDITOR specified by the environment variables """
        editor_cmd = self.editor().split()
        try:
            subprocess.call(editor_cmd + [filepath])
        except OSError:
            Utils.die('Could not launch ' + self.editor())


    @staticmethod
    def warn(message):
        """ Prints a message to stderr """
        print((message), file=sys.stderr)
