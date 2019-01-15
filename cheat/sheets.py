import os

from cheat.utils import Utils

class Sheets:


    def __init__(self, config):
        self._default_cheat_dir = config.get_default_cheat_dir()
        self._cheatpath = config.get_cheatpath()
        self._utils = Utils(config)

    def default_path(self):
        """ Returns the default cheatsheet path """

        # determine the default cheatsheet dir
        default_sheets_dir = self._default_cheat_dir or os.path.join('~', '.cheat')
        default_sheets_dir = os.path.expanduser(os.path.expandvars(default_sheets_dir))

        # create the DEFAULT_CHEAT_DIR if it does not exist
        if not os.path.isdir(default_sheets_dir):
            try:
                # @kludge: unclear on why this is necessary
                os.umask(0000)
                os.mkdir(default_sheets_dir)

            except OSError:
                Utils.die('Could not create DEFAULT_CHEAT_DIR')

        # assert that the DEFAULT_CHEAT_DIR is readable and writable
        if not os.access(default_sheets_dir, os.R_OK):
            Utils.die('The DEFAULT_CHEAT_DIR (' + default_sheets_dir +') is not readable.')
        if not os.access(default_sheets_dir, os.W_OK):
            Utils.die('The DEFAULT_CHEAT_DIR (' + default_sheets_dir +') is not writable.')

        # return the default dir
        return default_sheets_dir


    def get(self):
        """ Assembles a dictionary of cheatsheets as name => file-path """
        cheats = {}

        # otherwise, scan the filesystem
        for cheat_dir in reversed(self.paths()):
            cheats.update(
                dict([
                    (cheat, os.path.join(cheat_dir, cheat))
                    for cheat in os.listdir(cheat_dir)
                    if not cheat.startswith('.')
                    and not cheat.startswith('__')
                ])
            )

        return cheats


    def paths(self):
        """ Assembles a list of directories containing cheatsheets """
        sheet_paths = [
            self.default_path(),
            '/usr/share/cheat',
        ]

        # merge the CHEATPATH paths into the sheet_paths
        if self._cheatpath:
            for path in self._cheatpath.split(os.pathsep):
                if os.path.isdir(path):
                    sheet_paths.append(path)

        if not sheet_paths:
            Utils.die('The DEFAULT_CHEAT_DIR dir does not exist or the CHEATPATH is not set.')

        return sheet_paths


    def list(self):
        """ Lists the available cheatsheets """
        sheet_list = ''
        pad_length = max([len(x) for x in self.get().keys()]) + 4
        for sheet in sorted(self.get().items()):
            sheet_list += sheet[0].ljust(pad_length) + sheet[1] + "\n"
        return sheet_list


    def search(self,term):
        """ Searches all cheatsheets for the specified term """
        result = ''

        for cheatsheet in sorted(self.get().items()):
            match = ''
            for line in open(cheatsheet[1]):
                if term in line:
                    match += '  ' + self._utils.highlight(term, line)

            if match != '':
                result += cheatsheet[0] + ":\n" + match + "\n"

        return result
