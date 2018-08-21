import unittest2
import os
import shutil
from cheat.configuration import Configuration


def _set_loc_conf(key, value):
    _path = os.path.dirname(os.path.abspath(__file__)) + '/home/.config/cheat/cheat'
    if value == None:
        os.remove(_path)
    else:
        if not os.path.exists(os.path.dirname(_path)):
            os.makedirs(os.path.dirname(_path))
        f = open(_path,"w+")
        f.write('{"'+ key +'":"'+ value +'"}')
        f.close()


def _set_glob_conf(key, value):
    _path = os.path.dirname(os.path.abspath(__file__))+ "/etc/cheat"
    if value == None:
        os.remove(_path)
    else:
        if not os.path.exists(os.path.dirname(_path)):
            os.mkdir(os.path.dirname(_path))
        f = open(_path,"w+")
        f.write('{"'+ key +'":"'+ value +'"}' )
        f.close()


def _set_env_var(key, value):
    if value == None:
        del os.environ[key]
    else:
        os.environ[key] = value


def _configuration_key_test(TestConfiguration, key,values, conf_get_method):
    for glob_conf in values:
            _set_glob_conf(key,glob_conf)
            for loc_conf in values:
                _set_loc_conf(key,loc_conf)
                for env_conf in values:
                    _set_env_var(key,env_conf)
                    if env_conf:
                        TestConfiguration.assertEqual(conf_get_method(Configuration()),env_conf)
                    elif loc_conf:
                        TestConfiguration.assertEqual(conf_get_method(Configuration()),loc_conf)
                    elif glob_conf:
                        TestConfiguration.assertEqual(conf_get_method(Configuration()),glob_conf)
                    else:
                        TestConfiguration.assertEqual(conf_get_method(Configuration()),None)


class ConfigurationTestCase(unittest2.TestCase):


    def setUp(self):
        os.environ['CHEAT_GLOBAL_CONF_PATH'] = os.path.dirname(os.path.abspath(__file__)) \
        + '/etc/cheat'
        os.environ['CHEAT_LOCAL_CONF_PATH'] = os.path.dirname(os.path.abspath(__file__)) \
        + '/home/.config/cheat/cheat'


    def test_get_editor(self):
        _configuration_key_test(self,"EDITOR",["nano","vim","gedit",None],
        Configuration.get_editor)


    def test_get_cheatcolors(self):
        _configuration_key_test(self,"CHEATCOLORS",["true",None],
        Configuration.get_cheatcolors)


    def test_get_cheatpath(self):
        _configuration_key_test(self,"CHEATPATH",["/etc/myglobalcheats",
        "/etc/anotherglobalcheats","/rootcheats",None],Configuration.get_cheatpath)


    def test_get_defaultcheatdir(self):
        _configuration_key_test(self,"DEFAULT_CHEAT_DIR",["/etc/myglobalcheats",
        "/etc/anotherglobalcheats","/rootcheats",None],Configuration.get_default_cheat_dir)

    def tearDown(self):
        shutil.rmtree(os.path.dirname(os.path.abspath(__file__)) +'/etc')
        shutil.rmtree(os.path.dirname(os.path.abspath(__file__)) +'/home')
