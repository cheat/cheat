from distutils.core import setup
import os

# install appdirs if it cannot be imported
try:
    import appdirs
except ImportError:
    import pip
    pip.main(['install', 'appdirs'])
    import appdirs

# determine the path in which to install the cheatsheets included with the
# package
cheat_path = os.environ.get('CHEAT_PATH') or \
                appdirs.user_data_dir('cheat', 'cheat')

# determine the path in which to install the config file
config_path = os.environ.get('CHEAT_GLOBAL_CONF_PATH') or \
                os.environ.get('CHEAT_LOCAL_CONF_PATH') or \
                appdirs.user_config_dir('cheat', 'cheat')

# aggregate the systme-wide cheatsheets
cheat_files = []
for f in os.listdir('cheat/cheatsheets/'):
    cheat_files.append(os.path.join('cheat/cheatsheets/', f))

# specify build params
setup(
    name='cheat',
    version='2.5.1',
    author='Chris Lane',
    author_email='chris@chris-allen-lane.com',
    license='GPL3',
    description='cheat allows you to create and view interactive cheatsheets '
    'on the command-line. It was designed to help remind *nix system '
    'administrators of options for commands that they use frequently, but not '
    'frequently enough to remember.',
    url='https://github.com/chrisallenlane/cheat',
    packages=[
        'cheat',
        'cheat.test',
    ],
    scripts=['bin/cheat'],
    install_requires=[
        'appdirs >= 1.4.3',
        'docopt >= 0.6.1',
        'pygments >= 1.6.0',
        'termcolor >= 1.1.0',
    ],
    data_files=[
        (cheat_path, cheat_files),
        (config_path, ['config/cheat']),
    ],
)
