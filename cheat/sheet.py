from cheat.editor import Editor
from cheat.utils import Utils
import io
import os
import shutil


class Sheet:

    def __init__(self, config, sheets):
        self._config = config
        self._editor = Editor(config)
        self._sheets = sheets

    def _exists(self, sheet):
        """ Predicate that returns true if the sheet exists """
        return (sheet in self._sheets.get() and
                os.access(self._path(sheet), os.R_OK))

    def _exists_in_default_path(self, sheet):
        """ Predicate that returns true if the sheet exists in default_path"""
        default_path = os.path.join(self._config.cheat_user_dir, sheet)
        return (sheet in self._sheets.get() and
                os.access(default_path, os.R_OK))

    def _path(self, sheet):
        """ Returns a sheet's filesystem path """
        return self._sheets.get()[sheet]

    def edit(self, sheet):
        """ Creates or edits a cheatsheet """

        # if the cheatsheet does not exist
        if not self._exists(sheet):
            new_path = os.path.join(self._config.cheat_user_dir, sheet)
            self._editor.open(new_path)

        # if the cheatsheet exists but not in the default_path, copy it to the
        # default path before editing
        elif self._exists(sheet) and not self._exists_in_default_path(sheet):
            try:
                shutil.copy(
                            self._path(sheet),
                            os.path.join(self._config.cheat_user_dir, sheet)
                           )

            # fail gracefully if the cheatsheet cannot be copied. This can
            # happen if CHEAT_USER_DIR does not exist
            except IOError:
                Utils.die('Could not copy cheatsheet for editing.')

            self._editor.open(self._path(sheet))

        # if it exists and is in the default path, then just open it
        else:
            self._editor.open(self._path(sheet))

    def read(self, sheet):
        """ Returns the contents of the cheatsheet as a String """
        if not self._exists(sheet):
            Utils.die('No cheatsheet found for ' + sheet)

        with io.open(self._path(sheet), encoding='utf-8') as cheatfile:
            return cheatfile.read()
