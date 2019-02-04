from __future__ import print_function
import sys


class Colorize:

    def __init__(self, config):
        self._config = config

    def search(self, needle, haystack):
        """ Colorizes search results matched within a line """

        # if a highlight color is not configured, exit early
        if not self._config.cheat_highlight:
            return haystack

        # otherwise, attempt to import the termcolor library
        try:
            from termcolor import colored

        # if the import fails, return uncolored text
        except ImportError:
            return haystack

        # if the import succeeds, colorize the needle in haystack
        return haystack.replace(needle,
                                colored(needle, self._config.cheat_highlight))

    def syntax(self, sheet_content):
        """ Applies syntax highlighting """

        # only colorize if cheat_colors is true, and stdout is a tty
        if self._config.cheat_colors is False or not sys.stdout.isatty():
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
        lexer = get_lexer_by_name('bash')

        # apply syntax-highlighting if the first line is a code-fence
        if first_line.startswith('```'):
            sheet_content = '\n'.join(sheet_content.split('\n')[1:-2])
            try:
                lexer = get_lexer_by_name(first_line[3:])
            except Exception:
                pass

        return highlight(
               sheet_content,
               lexer,
               TerminalFormatter(bg=self._config.cheat_colorscheme))
