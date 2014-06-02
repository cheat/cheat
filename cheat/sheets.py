from cheat import cheatsheets
from cheat.utils import *
from xdg.BaseDirectory import save_data_path
import os

# @kludge: it breaks the functional paradigm to a degree, but declaring this
# var here (versus within get()) gives us a "poor man's" memoization on the
# call to get(). This, in turn, spares us from having to call out to the
# filesystem more than once.
cheats = {}


def default_path():
    """ Returns the default cheatsheet path """

    # the default path becomes confused when cheat is run as root, so fail
    # under those circumstances. (There is no good reason to need to run cheat
    # as root.)
    if os.name != 'nt' and os.geteuid() == 0:
        die('Please do not run this application as root.');

    # determine the default cheatsheet dir
    if os.environ.get('DEFAULT_CHEAT_DIR'):
        # if the environment variable is set, use it in any case
        default_sheets_dir = os.environ.get('DEFAULT_CHEAT_DIR')
    else:
        # use `.cheat` in the users HOME *if* it exists
        default_sheets_dir = os.path.join(os.path.expanduser('~'), '.cheat')
        if not os.path.isdir(default_sheets_dir):
            # or else use `cheat` in the systems prefered data-directory
            # the directory is created if it doesn't already exists
            default_sheets_dir = save_data_path('cheat')

    # create the DEFAULT_CHEAT_DIR if it does not exist
    if not os.path.isdir(default_sheets_dir):
        try:
            # @kludge: unclear on why this is necessary
            os.umask(0000)
            os.mkdir(default_sheets_dir)

        except OSError:
            die('Could not create DEFAULT_CHEAT_DIR')

    # assert that the DEFAULT_CHEAT_DIR is readable and writable
    if not os.access(default_sheets_dir, os.R_OK):
        die('The DEFAULT_CHEAT_DIR (' + default_sheets_dir +') is not readable.')
    if not os.access(default_sheets_dir, os.W_OK):
        die('The DEFAULT_CHEAT_DIR (' + default_sheets_dir +') is not writeable.')

    # return the default dir
    return default_sheets_dir


def get():
    """ Assembles a dictionary of cheatsheets as name => file-path """

    # if we've already reached out to the filesystem, just return the result
    # from memory
    if cheats:
        return cheats

    # otherwise, scan the filesystem
    for cheat_dir in reversed(paths()):
        cheats.update(
            dict([
                (cheat, os.path.join(cheat_dir, cheat))
                for cheat in os.listdir(cheat_dir)
                if not cheat.startswith('.')
                and not cheat.startswith('__')
            ])
        )

    return cheats


def paths():
    """ Assembles a list of directories containing cheatsheets """
    sheet_paths = [
        default_path(),
        cheatsheets.sheets_dir()[0],
    ]

    # merge the CHEATPATH paths into the sheet_paths
    if 'CHEATPATH' in os.environ and os.environ['CHEATPATH']:
        for path in os.environ['CHEATPATH'].split(os.pathsep):
            if os.path.isdir(path):
                sheet_paths.append(path)

    if not sheet_paths:
        die('The DEFAULT_CHEAT_DIR dir does not exist or the CHEATPATH is not set.')

    return sheet_paths


def list():
    """ Lists the available cheatsheets """
    sheet_list = ''
    pad_length = max([len(x) for x in get().keys()]) + 4
    for sheet in sorted(get().items()):
        sheet_list += sheet[0].ljust(pad_length) + sheet[1] + "\n"
    return sheet_list


def search(term):
    """ Searches all cheatsheets for the specified term """
    result = ''

    for cheatsheet in sorted(get().items()):
        match = ''
        for line in open(cheatsheet[1]):
             if term in line:
                  match += '  ' + line

        if not match == '':
            result += cheatsheet[0] + ":\n" + match + "\n"

    return result
