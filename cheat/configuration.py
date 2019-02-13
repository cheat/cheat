from cheat.utils import Utils
import json
import os


class Configuration:

    def __init__(self):
        # compute the location of the config files
        config_file_path_global = self._select([
            os.environ.get('CHEAT_GLOBAL_CONF_PATH'),
            '/etc/cheat',
        ])
        config_file_path_local = self._select([
            os.environ.get('CHEAT_LOCAL_CONF_PATH'),
            os.path.expanduser('~/.config/cheat/cheat'),
        ])

        # attempt to read the global config file
        config = {}
        try:
            config.update(self._read_config_file(config_file_path_global))
        except Exception as e:
            Utils.warn('Error while parsing global configuration: '
                       + e.message)

        # attempt to read the local config file
        try:
            config.update(self._read_config_file(config_file_path_local))
        except Exception as e:
            Utils.warn('Error while parsing local configuration: ' + e.message)

        # With config files read, now begin to apply envvar overrides and
        # default values

        # self.cheat_colors
        self.cheat_colors = self._select([
            Utils.boolify(os.environ.get('CHEAT_COLORS')),
            Utils.boolify(os.environ.get('CHEATCOLORS')),
            Utils.boolify(config.get('CHEAT_COLORS')),
            True,
        ])

        # self.cheat_colorscheme
        self.cheat_colorscheme = self._select([
            os.environ.get('CHEAT_COLORSCHEME'),
            config.get('CHEAT_COLORSCHEME'),
            'light',
        ]).strip().lower()

        # self.cheat_user_dir
        self.cheat_user_dir = self._select(
            map(os.path.expanduser,
                filter(None,
                    [os.environ.get('CHEAT_USER_DIR'),
                     os.environ.get('CHEAT_DEFAULT_DIR'),
                     os.environ.get('DEFAULT_CHEAT_DIR'),
                     # TODO: XDG home?
                     os.path.join('~', '.cheat')])))

        # self.cheat_editor
        self.cheat_editor = self._select([
            os.environ.get('CHEAT_EDITOR'),
            os.environ.get('EDITOR'),
            os.environ.get('VISUAL'),
            config.get('CHEAT_EDITOR'),
            'vi',
        ])

        # self.cheat_highlight
        self.cheat_highlight = self._select([
            os.environ.get('CHEAT_HIGHLIGHT'),
            config.get('CHEAT_HIGHLIGHT'),
            False,
        ])
        if isinstance(self.cheat_highlight, str):
            Utils.boolify(self.cheat_highlight)

        # self.cheat_path
        self.cheat_path = self._select([
            os.environ.get('CHEAT_PATH'),
            os.environ.get('CHEATPATH'),
            config.get('CHEAT_PATH'),
            '/usr/share/cheat',
        ])

    def _read_config_file(self, path):
        """ Reads configuration file and returns list of set variables """
        config = {}
        if os.path.isfile(path):
            with open(path) as config_file:
                config.update(json.load(config_file))
        return config

    def _select(self, values):
        for v in values:
            if v is not None:
                return v

    def validate(self):
        """ Validates configuration parameters """

        # assert that cheat_highlight contains a valid value
        highlights = [
            'grey', 'red', 'green', 'yellow',
            'blue', 'magenta', 'cyan', 'white',
            False
        ]
        if self.cheat_highlight not in highlights:
            Utils.die("%s %s" %
                      ('CHEAT_HIGHLIGHT must be one of:', highlights))

        # assert that the color scheme is valid
        colorschemes = ['light', 'dark']
        if self.cheat_colorscheme not in colorschemes:
            Utils.die("%s %s" %
                      ('CHEAT_COLORSCHEME must be one of:', colorschemes))

        return True
