from mock import patch
import os
from tempfile import mkdtemp
from unittest import TestCase

from cheat import cheatsheets
from cheat.sheets import (
    default_path, paths, get,
)


class TestSheets(TestCase):
    def setUp(self):
        self.tmp_dir = mkdtemp()

    def _make_sheets(self, sheets):
        for sheet in sheets:
            dir_, file_ = os.path.split(sheet)
            if not os.path.exists(dir_):
                os.mkdir(dir_)
            open(sheet, 'w').close()

    @patch('cheat.sheets.os.environ.get')
    def test_default_path_via_env(self, mock_environ_get):
        mock_environ_get.return_value = self.tmp_dir

        path = default_path()

        self.assertEqual(path, self.tmp_dir)

    @patch('cheat.sheets.os.environ.get')
    def test_default_path_via_env_unexists(self, mock_environ_get):
        os.rmdir(self.tmp_dir)
        mock_environ_get.return_value = self.tmp_dir

        path = default_path()

        self.assertEqual(path, self.tmp_dir)
        os.path.isdir(self.tmp_dir)

    @patch('cheat.sheets.os.path.expanduser')
    @patch('cheat.sheets.os.environ.get')
    def test_default_path_from_home(self, mock_environ_get, mock_expanduser):
        mock_environ_get.return_value = None
        mock_expanduser.return_value = self.tmp_dir
        os.mkdir(os.path.join(self.tmp_dir, '.cheat'))

        path = default_path()

        self.assertEqual(path, os.path.join(self.tmp_dir, '.cheat'))

    @patch('cheat.sheets.os.path.expanduser')
    @patch('cheat.sheets.os.environ.get')
    def test_default_path_from_home_not_exists(
            self, mock_environ_get, mock_expanduser):
        mock_environ_get.return_value = None
        mock_expanduser.return_value = self.tmp_dir

        path = default_path()

        self.assertEqual(path, os.path.join(self.tmp_dir, '.cheat'))
        os.path.isdir(os.path.join(self.tmp_dir, '.cheat'))

    @patch('cheat.sheets.die')
    @patch('cheat.sheets.os.environ.get')
    def test_default_path_unreadable(self, mock_environ_get, mock_die):
        mock_environ_get.return_value = self.tmp_dir
        os.chmod(self.tmp_dir, int('333', 8))

        default_path()

        mock_die.assert_called_with(
            'The DEFAULT_CHEAT_DIR (%s) is not readable.' % self.tmp_dir)

    @patch('cheat.sheets.die')
    @patch('cheat.sheets.os.environ.get')
    def test_default_path_untouchable(self, mock_environ_get, mock_die):
        mock_environ_get.return_value = self.tmp_dir
        os.chmod(self.tmp_dir, int('555', 8))

        default_path()

        mock_die.assert_called_with(
            'The DEFAULT_CHEAT_DIR (%s) is not writable.' % self.tmp_dir)

    @patch('cheat.sheets.default_path')
    @patch('cheat.sheets.os.environ.get')
    def test_paths_only_default(self, mock_environ_get, mock_default_path):
        mock_environ_get.return_value = None
        mock_default_path.return_value = self.tmp_dir

        result = paths()

        self.assertEqual(result, [self.tmp_dir, cheatsheets.sheets_dir()[0]])

    @patch('cheat.sheets.default_path')
    @patch('cheat.sheets.os.environ.get')
    def test_paths_with_additonal_definition(
            self, mock_environ_get, mock_default_path):
        additonal_paths = [
            os.path.join(self.tmp_dir, 'path1'),
            os.path.join(self.tmp_dir, 'path2'),
        ]
        for path in additonal_paths:
            os.mkdir(path)
        mock_environ_get.return_value = os.pathsep.join(additonal_paths)
        mock_default_path.return_value = self.tmp_dir

        result = paths()

        self.assertEqual(
                result,
                [self.tmp_dir, cheatsheets.sheets_dir()[0]] +
                additonal_paths)

    @patch('cheat.sheets.paths')
    def test_get(self, mock_paths):
        self._make_sheets([
            os.path.join(self.tmp_dir, 'cheatdir1', 'curl'),
            os.path.join(self.tmp_dir, 'cheatdir1', 'http'),
            os.path.join(self.tmp_dir, 'cheatdir2', 'tar'),
            os.path.join(self.tmp_dir, 'cheatdir2', 'zip'),
        ])
        mock_paths.return_value = [
            os.path.join(self.tmp_dir, 'cheatdir1'),
            os.path.join(self.tmp_dir, 'cheatdir2'),
        ]

        cheats = get()
        self.assertEqual(
            cheats,
            {
                'curl': os.path.join(self.tmp_dir, 'cheatdir1', 'curl'),
                'http': os.path.join(self.tmp_dir, 'cheatdir1', 'http'),
                'tar': os.path.join(self.tmp_dir, 'cheatdir2', 'tar'),
                'zip': os.path.join(self.tmp_dir, 'cheatdir2', 'zip'),
            }
        )

    @patch('cheat.sheets.paths')
    def test_get_with_duplicate_cheat(self, mock_paths):
        self._make_sheets([
            os.path.join(self.tmp_dir, 'cheatdir1', 'curl'),
            os.path.join(self.tmp_dir, 'cheatdir1', 'http'),
            os.path.join(self.tmp_dir, 'cheatdir2', 'tar'),
            os.path.join(self.tmp_dir, 'cheatdir2', 'zip'),
            os.path.join(self.tmp_dir, 'cheatdir2', 'http'),
        ])
        mock_paths.return_value = [
            os.path.join(self.tmp_dir, 'cheatdir1'),
            os.path.join(self.tmp_dir, 'cheatdir2'),
        ]

        cheats = get()
        self.assertEqual(
            cheats,
            {
                'curl': os.path.join(self.tmp_dir, 'cheatdir1', 'curl'),
                'http': os.path.join(self.tmp_dir, 'cheatdir1', 'http'),
                'tar': os.path.join(self.tmp_dir, 'cheatdir2', 'tar'),
                'zip': os.path.join(self.tmp_dir, 'cheatdir2', 'zip'),
            }
        )

    @patch('cheat.sheets.paths')
    def test_get_with_exclude_pattern(self, mock_paths):
        self._make_sheets([
            os.path.join(self.tmp_dir, 'cheatdir1', 'curl'),
            os.path.join(self.tmp_dir, 'cheatdir1', 'http'),
            os.path.join(self.tmp_dir, 'cheatdir1', '__init__.py'),
            os.path.join(self.tmp_dir, 'cheatdir2', 'tar'),
            os.path.join(self.tmp_dir, 'cheatdir2', 'zip'),
            os.path.join(self.tmp_dir, 'cheatdir2', '.git'),
        ])
        mock_paths.return_value = [
            os.path.join(self.tmp_dir, 'cheatdir1'),
            os.path.join(self.tmp_dir, 'cheatdir2'),
        ]

        cheats = get()
        self.assertEqual(
            cheats,
            {
                'curl': os.path.join(self.tmp_dir, 'cheatdir1', 'curl'),
                'http': os.path.join(self.tmp_dir, 'cheatdir1', 'http'),
                'tar': os.path.join(self.tmp_dir, 'cheatdir2', 'tar'),
                'zip': os.path.join(self.tmp_dir, 'cheatdir2', 'zip'),
            }
        )
