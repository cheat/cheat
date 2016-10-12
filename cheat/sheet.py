import os
import shutil

from cheat import sheets
from cheat.utils import die, open_with_editor

def copy(current_sheet_path, new_sheet_path):
    """ Copies a sheet to a new path """

    # attempt to copy the sheet to DEFAULT_CHEAT_DIR
    try:
        shutil.copy(current_sheet_path, new_sheet_path)

    # fail gracefully if the cheatsheet cannot be copied. This can happen if
    # DEFAULT_CHEAT_DIR does not exist
    except IOError:
        die('Could not copy cheatsheet for editing.')


def create_or_edit(sheet):
    """ Creates or edits a cheatsheet """

    # if the cheatsheet does not exist
    if not exists(sheet):
        create(sheet)

    # if the cheatsheet exists but not in the default_path, copy it to the
    # default path before editing
    elif exists(sheet) and not exists_in_default_path(sheet):
        copy(path(sheet), os.path.join(sheets.default_path(), sheet))
        edit(sheet)

    # if it exists and is in the default path, then just open it
    else:
        edit(sheet)


def create(sheet):
    """ Creates a cheatsheet """
    new_sheet_path = os.path.join(sheets.default_path(), sheet)
    open_with_editor(new_sheet_path)


def edit(sheet):
    """ Opens a cheatsheet for editing """
    open_with_editor(path(sheet))


def exists(sheet):
    """ Predicate that returns true if the sheet exists """
    return sheet in sheets.get() and os.access(path(sheet), os.R_OK)


def exists_in_default_path(sheet):
    """ Predicate that returns true if the sheet exists in default_path"""
    default_path_sheet = os.path.join(sheets.default_path(), sheet)
    return sheet in sheets.get() and os.access(default_path_sheet, os.R_OK)


def is_writable(sheet):
    """ Predicate that returns true if the sheet is writeable """
    return sheet in sheets.get() and os.access(path(sheet), os.W_OK)


def path(sheet):
    """ Returns a sheet's filesystem path """
    return sheets.get()[sheet]


def read(sheet):
    """ Returns the contents of the cheatsheet as a String """
    if not exists(sheet):
        die('No cheatsheet found for ' + sheet)

    with open(path(sheet)) as cheatfile:
        return cheatfile.read()
