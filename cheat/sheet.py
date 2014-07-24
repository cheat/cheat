from cheat import sheets 
from cheat import utils
from cheat.utils import *
import os
import shutil
import subprocess


def copy(current_sheet_path, new_sheet_path):
    """ Copies a sheet to a new path """

    # attempt to copy the sheet to DEFAULT_CHEAT_DIR
    try:
        shutil.copy(current_sheet_path, new_sheet_path)

    # fail gracefully if the cheatsheet cannot be copied. This can happen if
    # DEFAULT_CHEAT_DIR does not exist
    except IOError:
        die ('Could not copy cheatsheet for editing.')


def create_or_edit(sheet):
    """ Creates or edits a cheatsheet """

    # if the cheatsheet does not exist
    if not exists(sheet):
        create(sheet)

    # if the cheatsheet exists and is writeable...
    elif exists(sheet) and is_writable(sheet):
        edit(sheet)

    # if the cheatsheet exists but is not writable...
    elif exists(sheet) and not is_writable(sheet):
        # ... ask the user if we should copy the cheatsheet to her home directory for editing
        yes = prompt_yes_or_no(
          'The ' + sheet + ' sheet is not editable. Do you want to copy it to '
          'your user cheatsheets directory before editing? Keep in mind that '
          'your sheet will always be used before system-wide one.'
        )

        # if yes, copy the cheatsheet to the home directory before editing
        if yes:
            new_sheet_path = os.path.join(sheets.default_path(), sheet)
            copy(path(sheet), new_sheet_path)
            # don't call edit() since the cheatsheets cache would return the previous location of the sheet
            try:
                subprocess.call([editor(), new_sheet_path])

            except OSError:
                die('Could not launch ' + editor())

        # if no, just abort
        else:
            die('Aborting.')


def create(sheet):
    """ Creates a cheatsheet """
    new_sheet_path = os.path.join(sheets.default_path(), sheet)

    try:
        subprocess.call([editor(), new_sheet_path])

    except OSError:
        die('Could not launch ' + editor())


def edit(sheet):
    """ Opens a cheatsheet for editing """

    try:
        subprocess.call([editor(), path(sheet)])

    except OSError:
        die('Could not launch ' + editor())


def exists(sheet):
    """ Predicate that returns true if the sheet exists """
    return sheet in sheets.get() and os.access(path(sheet), os.R_OK)


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

    with open (path(sheet)) as cheatfile:
          return cheatfile.read()
