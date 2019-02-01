from cheat.colorize import Colorize
from cheat.utils import Utils
import io
import os


class Sheets:

    def __init__(self, config):
        self._config = config
        self._colorize = Colorize(config)

        # Assembles a dictionary of cheatsheets as name => file-path
        self._sheets = {}
        sheet_paths = [
            config.cheat_user_dir
        ]

        # merge the CHEAT_PATH paths into the sheet_paths
        if config.cheat_path:
            for path in config.cheat_path.split(os.pathsep):
                if os.path.isdir(path):
                    sheet_paths.append(path)

        if not sheet_paths:
            Utils.die('The CHEAT_USER_DIR dir does not exist '
                      + 'or the CHEAT_PATH is not set.')

        # otherwise, scan the filesystem
        for cheat_dir in reversed(sheet_paths):
            self._sheets.update(
                dict([
                    (cheat, os.path.join(cheat_dir, cheat))
                    for cheat in os.listdir(cheat_dir)
                    if not cheat.startswith('.')
                    and not cheat.startswith('__')
                ])
            )

    def directories(self):
        """ Assembles a list of directories containing cheatsheets """
        sheet_paths = [
            self._config.cheat_user_dir,
        ]

        # merge the CHEATPATH paths into the sheet_paths
        for path in self._config.cheat_path.split(os.pathsep):
            sheet_paths.append(path)

        return sheet_paths

    def get(self):
        """ Returns a dictionary of cheatsheets as name => file-path """
        return self._sheets

    def list(self):
        """ Lists the available cheatsheets """
        sheet_list = ''
        pad_length = max([len(x) for x in self.get().keys()]) + 4
        for sheet in sorted(self.get().items()):
            sheet_list += sheet[0].ljust(pad_length) + sheet[1] + "\n"
        return sheet_list

    def search(self, term):
        """ Searches all cheatsheets for the specified term """
        result = ''

        for cheatsheet in sorted(self.get().items()):
            match = ''
            for line in io.open(cheatsheet[1], encoding='utf-8'):
                if term in line:
                    match += '  ' + self._colorize.search(term, line)

            if match != '':
                result += cheatsheet[0] + ":\n" + match + "\n"

        return result
