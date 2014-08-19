import os
import sys

from xdg.BaseDirectory import save_data_path


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
            user_dir = get_userdata_dir()
    return user_dir


def get_userdata_dir():
    """
    Returns the full path to a `cheat` directory in the platform specific
    default data directory for the current user.

    .. note:: The directory is created, if it's not already present.

    """
    return save_data_path('cheat')
