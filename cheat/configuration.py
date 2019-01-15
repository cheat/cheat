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
            merged_config.update(self._read_configuration_file(self.glob_config_path))
        except Exception as e:
            Utils.warn('error while parsing global configuration Reason: ' + e.message)

        try:
            merged_config.update(self._read_configuration_file(self.local_config_path))
        except Exception as e:
            Utils.warn('error while parsing user configuration Reason: ' + e.message)

        merged_config.update(self._read_env_vars_config())



        return merged_config


    def _read_configuration_file(self,path):
        # Reads configuration file and returns list of set variables
        read_config = {}
        if (os.path.isfile(path)):
            with open(path) as config_file:
                read_config.update(json.load(config_file))
        return read_config


    def _read_env_vars_config(self):
        read_config = {}

        # All these variables are left here because of backwards compatibility

        if (os.environ.get('CHEAT_EDITOR')):
            read_config['EDITOR'] = os.environ.get('CHEAT_EDITOR')

        if (os.environ.get('VISUAL')):
            read_config['EDITOR'] = os.environ.get('VISUAL')

        keys = ['DEFAULT_CHEAT_DIR',
                'CHEATPATH',
                'CHEATCOLORS',
                'EDITOR',
                'CHEAT_HIGHLIGHT'
               ]

        for k in keys:
            self._read_env_var(read_config,k)

        return read_config


    def _read_env_var(self,current_config,key):
        if (os.environ.get(key)):
            current_config[key] = os.environ.get(key)


    def _get_global_conf_file_path(self):
        self.glob_config_path = os.environ.get('CHEAT_GLOBAL_CONF_PATH') \
        or '/etc/cheat'


    def _get_local_conf_file_path(self):
        self.local_config_path = os.environ.get('CHEAT_LOCAL_CONF_PATH') \
        or os.path.expanduser('~/.config/cheat/cheat')


    def get_default_cheat_dir(self):
        return self._saved_configuration.get('DEFAULT_CHEAT_DIR')


    def get_cheatpath(self):
        return self._saved_configuration.get('CHEATPATH')


    def get_cheatcolors(self):
        return self._saved_configuration.get('CHEATCOLORS')


    def get_editor(self):
        return self._saved_configuration.get('EDITOR')

    def get_highlight(self):
        return self._saved_configuration.get('CHEAT_HIGHLIGHT')
