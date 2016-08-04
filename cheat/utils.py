from __future__ import print_function
import os
import sys

from appdirs import user_data_dir


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
    # Support Python 2 and 3 input
    # Default to Python 2's input()
    get_input = raw_input

    # If this is Python 3, use input()
    if sys.version_info[:2] >= (3, 0):
        get_input = input

    print(question)
    return get_input('[y/n] ') == 'y'


def warn(message):
    """ Prints a message to stderr """
    print((message), file=sys.stderr)


def get_default_data_dir():
    """
    Returns the full path to the directory containing the users data.

    Which directory is used, is determined the following way:

    1. If the `DEFAULT_CHEAT_DIR` environment variable is set, use it.

    2. If a `.cheat` directory exists in the home directory, use it.

    3. Use a `cheat` directory in the systems default directory for user data.

    """
    user_dir = os.environ.get("DEFAULT_CHEAT_DIR")
    if not user_dir:
        user_dir = os.path.expanduser(os.path.join("~", ".cheat"))
        if not os.path.exists(user_dir):
            user_dir = user_data_dir('cheat')
    return user_dir
