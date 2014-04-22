import os
from StringIO import StringIO
import sys

from mock import patch
from nose.tools import assert_equals, assert_true, assert_false, assert_raises


class TestCheatSheets(object):

    def setUp(self):
        self.cheatpath = os.path.join(
            os.path.dirname(__file__), '.cheats')
        os.environ['CHEATPATH'] = self.cheatpath
        os.environ['DEFAULT_CHEAT_DIR'] = ''
        import imp
        self.cheat = imp.load_source('cheat', './cheat')
        self.cheatsheets = self.cheat.CheatSheets()
        # Change stdout for testing.
        # It would be preferable to make CheatSheets more testable.
        self.out = StringIO()
        sys.stdout = self.out
        assert_true(self.cheatpath in self.cheatsheets.dirs)

    def test_display_foobar_cheat_sheet(self):
        """The display_sheet method on a test cheatsheet should dump the known
        content to stdout."""
        # There should be a testing cheatsheet called '-test-foobar'.
        cheatfile = '-test-foobar'
        expected = 'binbat'
        cheatfile_path = os.path.join(self.cheatpath, cheatfile)
        assert_true(os.path.exists(cheatfile_path))
        # The content of '-test-foobar' should start with 'binbat'.
        # CheatSheets.display_sheet() may add extra "pretty" formatting, so
        # for this test, we only care that only the first set of characters
        # match.
        assert_equals(None, self.cheatsheets.display_sheet(cheatfile))
        assert_equals(expected, self.out.getvalue()[0:len(expected)])
        assert_true(cheatfile in self.cheatsheets.sheets)

    def test_display_missing_cheat_sheet_file(self):
        """Attempting to display a non-existant cheatsheet file directly with
        display_sheet() raises a KeyError."""
        cheatfile = '-test-missing_file'
        cheatfile_path = os.path.join(self.cheatpath, cheatfile)
        # Verify that test file is indeed missing.
        assert_false(os.path.exists(cheatfile_path))
        # Verify that missing cheatfile is not in our cheatsheet.sheets.
        assert_false(cheatfile in self.cheatsheets.sheets)
        with assert_raises(KeyError):
            self.cheatsheets.display_sheet(cheatfile)

    @patch('cheat.CheatSheets.vim_view', return_value=None)
    def test_vim_view_call(self, vim_view):
        """The display_sheet() method will invoke a vim_view call when trying
        to open a Vim Encrypted file."""
        cheatfile = '-test-vim_crypted'
        cheatfile_path = os.path.join(self.cheatpath, cheatfile)
        assert_true(os.path.exists(cheatfile_path))
        assert_true(self.cheatsheets.is_vim_crypted(cheatfile_path))
        assert_false(vim_view.called)
        self.cheatsheets.display_sheet(cheatfile)
        vim_view.assert_called_with(cheatfile_path)

    def test_is_vim_crypted(self):
        """is_vim_crypted() helper method detects a Vim Encrypted file, vs a
        non-encrypted file."""
        # This test file is clear text.
        clear_cheatfile = '-test-foobar'
        clear_cheatfile_path = os.path.join(self.cheatpath, clear_cheatfile)
        assert_false(self.cheatsheets.is_vim_crypted(clear_cheatfile_path))

        # This file is "Vim Encrypted" with the pass phrase "passphrase".
        vimcrypted_cheatfile = '-test-vim_crypted'
        vimcrypted_cheatfile_path = os.path.join(
            self.cheatpath, vimcrypted_cheatfile)
        assert_true(self.cheatsheets.is_vim_crypted(vimcrypted_cheatfile_path))

    @patch('subprocess.call', return_value=None)
    def test_vim_view_cmd(self, subprocess_call):
        """vim_view() uses subprocess.call() correctly."""
        vimcrypted_cheatfile = '-test-vim_crypted'
        vimcrypted_cheatfile_path = os.path.join(
            self.cheatpath, vimcrypted_cheatfile)
        assert_false(subprocess_call.called)
        assert_true(self.cheatsheets.vim_view(vimcrypted_cheatfile_path))
        assert_true(subprocess_call.called)
        assert_equals(
            ['/usr/bin/env', 'vim', '-R', vimcrypted_cheatfile_path],
            subprocess_call.call_args[0][0]
        )
