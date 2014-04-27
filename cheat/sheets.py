from cheat import cheatsheets
from cheat.utils import *
import os


def default_path():
    """ Returns the default cheatsheet path """

    # the default path becomes confused when cheat is run as root, so fail
    # under those circumstances. (There is no good reason to need to run cheat
    # as root.)
    if os.geteuid() == 0:
        die('Please do not run this application as root.');

    return os.environ.get('DEFAULT_CHEAT_DIR') or os.path.join(os.path.expanduser('~'), '.cheat')


# @todo: memoize result
def get():
    """ Assembles a dictionary of cheatsheets as name => file-path """
    cheats = {}
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
