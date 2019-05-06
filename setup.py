from distutils.core import setup
import os

# determine the directory in which to install system-wide cheatsheets
# KLUDGE: It would be better to read `/usr/share/cheat` from `config/cheat`
# rather than hard-coding it here
cheat_path = os.environ.get('CHEAT_PATH') or '/usr/share/cheat'

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
        'docopt >= 0.6.1',
        'pygments >= 1.6.0',
        'termcolor >= 1.1.0',
    ],
    data_files=[
        (cheat_path, cheat_files),
        ('/etc', ['config/cheat']),
    ],
)
