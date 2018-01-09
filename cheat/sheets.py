import os

from cheat import cheatsheets
from cheat.utils import die

def default_path():
    """ Returns the default cheatsheet path """

    # determine the default cheatsheet dir
    default_sheets_dir = os.environ.get('DEFAULT_CHEAT_DIR') or os.path.join('~', '.cheat')
    default_sheets_dir = os.path.expanduser(os.path.expandvars(default_sheets_dir))

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
        die('The DEFAULT_CHEAT_DIR (' + default_sheets_dir +') is not writable.')

    # return the default dir
    return default_sheets_dir


def get():
    """ Assembles a dictionary of cheatsheets as name => file-path """
    cheats = {}

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

# GMFTBY add : 
# try to return the subdir(s) under the sheet_paths created by the `paths` function below
def get_subdir(current_dir):
    '''
        Input  : the path of the current_dir
        Output : the list contain the subdirs

        P.S : maybe need to add the limitation about the search depth ?
    '''
    result = []
    for f_name in os.listdir(current_dir):
        if f_name.startswith('__') or f_name.startswith('.') : continue 
        whole_path = os.path.join(current_dir, f_name)
        if os.path.isdir(whole_path):
            result.append(whole_path)
            result.extend(get_subdir(whole_path))    # GMFTBY add : recursion, function
    return result

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

    # GMFTBY add : add the subdir under the current_dir into the sheet_paths
    for_extend = []
    for path in sheet_paths:
        subdirs = get_subdir(path)
        if subdirs:
            for_extend.extend(subdirs)
    sheet_paths.extend(for_extend)

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
        # GMFTBY add : decide not to open the dir or the illegal file
        if os.path.isdir(cheatsheet[1]) == True \
                or os.access(cheatsheet[1], os.R_OK) == False: continue
        match = ''
        # GMFTBY add : add the line number into the match message
        for index, line in enumerate(open(cheatsheet[1])):
            if term in line:
                match += '[%d]\t\t' % index + line.strip() + '\n'

        if match != '':
            result += cheatsheet[0] + ":\n" + match + "\n"

    return result
