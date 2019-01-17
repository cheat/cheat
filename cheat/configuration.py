import os
from cheat.utils import Utils
import json


class Configuration:

    def __init__(self):
        self._get_global_conf_file_path()
        self._get_local_conf_file_path()
        self._saved_configuration = self._get_configuration()

    def _get_configuration(self):
        # get options from config files and environment vairables
        merged_config = {}

        try:
            merged_config.update(
                self._read_configuration_file(self.glob_config_path)
            )
        except Exception as e:
            Utils.warn('error while parsing global configuration Reason: '
                       + e.message
                       )

        try:
            merged_config.update(
                self._read_configuration_file(self.local_config_path)
            )
        except Exception as e:
            Utils.warn('error while parsing user configuration Reason: '
                       + e.message
                       )

        merged_config.update(self._read_env_vars_config())

        self._check_configuration(merged_config)

        return merged_config

    def _read_configuration_file(self, path):
        # Reads configuration file and returns list of set variables
        read_config = {}
        if (os.path.isfile(path)):
            with open(path) as config_file:
                read_config.update(json.load(config_file))
        return read_config

    def _read_env_vars_config(self):
        read_config = {}

        # NOTE: These variables are left here because of backwards
        # compatibility and are supported only as env vars but not in
        # configuration file

        if (os.environ.get('VISUAL')):
            read_config['EDITOR'] = os.environ.get('VISUAL')

        # variables supported both in environment and configuration file
        # NOTE: Variables without CHEAT_ prefix are legacy
        # key is variable name and value is its legacy_alias
        # if variable has no legacy alias then set to None
        variables = {'CHEAT_DEFAULT_DIR': 'DEFAULT_CHEAT_DIR',
                     'CHEAT_PATH': 'CHEATPATH',
                     'CHEAT_COLORS': 'CHEATCOLORS',
                     'CHEAT_EDITOR': 'EDITOR',
                     'CHEAT_HIGHLIGHT': None
                     }

        for (k, v) in variables.items():
            self._read_env_var(read_config, k, v)

        return read_config

    def _check_configuration(self, config):
        """ Check values in config and warn user or die """

        # validate CHEAT_HIGHLIGHT values if set
        colors = [
            'grey', 'red', 'green', 'yellow',
            'blue', 'magenta', 'cyan', 'white'
        ]
        if (
            config.get('CHEAT_HIGHLIGHT') and
            config.get('CHEAT_HIGHLIGHT') not in colors
        ):
            Utils.die("%s %s" % ('CHEAT_HIGHLIGHT must be one of:', colors))

    def _read_env_var(self, current_config, key, alias=None):
        if os.environ.get(key) is not None:
            current_config[key] = os.environ.get(key)
            return
        elif alias is not None and os.environ.get(alias) is not None:
            current_config[key] = os.environ.get(alias)
            return

    def _get_global_conf_file_path(self):
        self.glob_config_path = (os.environ.get('CHEAT_GLOBAL_CONF_PATH')
                                 or '/etc/cheat')

    def _get_local_conf_file_path(self):
        path = (os.environ.get('CHEAT_LOCAL_CONF_PATH')
                or os.path.expanduser('~/.config/cheat/cheat'))
        self.local_config_path = path

    def _choose_value(self, primary_value_name, secondary_value_name):
        """ Return primary or secondary value in saved_configuration

        If primary value is in configuration then return it. If it is not
        then return secondary. In the absence of both values return None
        """

        primary_value = self._saved_configuration.get(primary_value_name)
        secondary_value = self._saved_configuration.get(secondary_value_name)

        if primary_value is not None:
            return primary_value
        else:
            return secondary_value

    def get_default_cheat_dir(self):
        return self._choose_value('CHEAT_DEFAULT_DIR', 'DEFAULT_CHEAT_DIR')

    def get_cheatpath(self):
        return self._choose_value('CHEAT_PATH', 'CHEATPATH')

    def get_cheatcolors(self):
        return self._choose_value('CHEAT_COLORS', 'CHEATCOLORS')

    def get_editor(self):
        return self._choose_value('CHEAT_EDITOR', 'EDITOR')

    def get_highlight(self):
        return self._saved_configuration.get('CHEAT_HIGHLIGHT')
