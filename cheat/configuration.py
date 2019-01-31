from cheat.utils import Utils
import json
import os

class Configuration:

    def __init__(self):
        # compute the location of the config files
        config_file_path_global = os.environ.get('CHEAT_GLOBAL_CONF_PATH') \
            or '/etc/cheat'
        config_file_path_local = (os.environ.get('CHEAT_LOCAL_CONF_PATH')  \
            or os.path.expanduser('~/.config/cheat/cheat'))

        # attempt to read the global config file
        config = {}
        try:
            config.update(self._read_config_file(config_file_path_global))
        except Exception as e:
            Utils.warn('Error while parsing global configuration: ' + e.message)

        # attempt to read the local config file
        try:
            config.update(self._read_config_file(config_file_path_local))
        except Exception as e:
            Utils.warn('Error while parsing local configuration: ' + e.message)

        # With config files read, now begin to apply envvar overrides and
        # default values

        # self.cheat_colors
        self.cheat_colors = self._select([
            self._boolify(os.environ.get('CHEAT_COLORS')),
            self._boolify(os.environ.get('CHEATCOLORS')),
            self._boolify(config.get('CHEAT_COLORS')),
            True,
        ])

        # self.cheat_default_dir
        self.cheat_default_dir = self._select([
            os.environ.get('CHEAT_DEFAULT_DIR'),
            os.environ.get('DEFAULT_CHEAT_DIR'),
            '~/.cheat',
        ])

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
        if (self.cheat_highlight.strip().lower() == "false"):
            self.cheat_highlight = False

        # self.cheat_path
        self.cheat_path = self._select([
            os.environ.get('CHEAT_PATH'),
            os.environ.get('CHEATPATH'),
            config.get('CHEAT_PATH'),
            '/usr/share/cheat',
        ])

    def _boolify(self, value):
        # if `value` is not a string, return it as-is
        if not isinstance(value, str):
            return value

        # otherwise, convert "true" and "false" to Boolean counterparts
        return value.strip().lower() == "true"

    def _read_config_file(self, path):
        # Reads configuration file and returns list of set variables
        config = {}
        if (os.path.isfile(path)):
            with open(path) as config_file:
                config.update(json.load(config_file))
        return config

    def _select(self, values):
        for v in values:
            if v is not None:
                return v

    def validate(self):
        """ Validate configuration parameters """

        # assert that cheat_highlight contains a valid value
        highlights = [
            'grey', 'red', 'green', 'yellow',
            'blue', 'magenta', 'cyan', 'white',
            False
        ]
        if (self.cheat_highlight not in highlights):
            Utils.die("%s %s" % ('CHEAT_HIGHLIGHT must be one of:', highlights))

        return True
