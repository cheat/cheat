from mock import patch
import os
from tempfile import mkdtemp
from unittest import TestCase

from cheat import sheet


class TestSheet(TestCase):
    def setUp(self):
        self.tmp_dir = mkdtemp()

        self.editor = patch('cheat.sheet.editor').start()
        self.editor.return_value = 'vim'

        self.die = patch('cheat.sheet.die').start()

        self.builtin_path = os.path.join(self.tmp_dir, 'builtin')
        os.mkdir(self.builtin_path)

        self.default_path = os.path.join(self.tmp_dir, 'default')
        os.mkdir(self.default_path)
        self.sheets_default_path = \
            patch('cheat.sheet.sheets.default_path').start()
        self.sheets_default_path.return_value = self.default_path

        self.sheets_paths = patch('cheat.sheet.sheets.paths').start()
        self.sheets_paths.return_value = [
            self.default_path, self.builtin_path,
        ]

        self.subprocess = patch('cheat.sheet.subprocess').start()

    def test_create_or_edit_when_sheet_inexist(self):
        sheet.create_or_edit('new')

        self.subprocess.call.assert_called_with([
            'vim', os.path.join(self.default_path, 'new')])

    def test_create_or_edit_when_sheet_inexist_at_default_path(self):
        cheat_path = os.path.join(self.builtin_path, 'cheat')
        with open(cheat_path, 'w') as f:
            f.write((
                'cheatsheet for cheat\n'
                'cheat cheat\n'
            ))

        sheet.create_or_edit('cheat')

        new_cheat_path = os.path.join(self.default_path, 'cheat')
        self.assertTrue(os.path.exists(new_cheat_path))
        with open(new_cheat_path, 'r') as f:
            self.assertEqual(
                f.read(),
                ('cheatsheet for cheat\n'
                 'cheat cheat\n')
            )

        self.subprocess.call.assert_called_with([
            'vim', new_cheat_path])

    def test_create_or_edit_when_sheet_exist_at_default_path(self):
        cheat_path = os.path.join(self.default_path, 'cheat')
        open(cheat_path, 'w').close()

        sheet.create_or_edit('cheat')
        self.subprocess.call.assert_called_with([
            'vim', cheat_path])

    def test_read(self):
        cheat_path = os.path.join(self.builtin_path, 'cheat')
        with open(cheat_path, 'w') as f:
            f.write((
                'cheatsheet for cheat\n'
                'cheat cheat\n'
            ))

        self.assertEqual(
            sheet.read('cheat'),
            ('cheatsheet for cheat\n'
             'cheat cheat\n')
        )
