from distutils.core import setup
import os

cheat_files = []
for f in os.listdir('cheat/cheatsheets/'):
    cheat_files.append(os.path.join('cheat/cheatsheets/',f))

setup(
    name         = 'cheat',
    version      = '2.4.0',
    author       = 'Chris Lane',
    author_email = 'chris@chris-allen-lane.com',
    license      = 'GPL3',
    description  = 'cheat allows you to create and view interactive cheatsheets '
    'on the command-line. It was designed to help remind *nix system '
    'administrators of options for commands that they use frequently, but not '
    'frequently enough to remember.',
    url          = 'https://github.com/chrisallenlane/cheat',
    packages     = [
        'cheat',
        'cheat.test',
    ],
    scripts          = ['bin/cheat'],
    install_requires = [
        'docopt >= 0.6.1',
        'pygments >= 1.6.0',
    ],
    data_files = [
        ('/usr/share/cheat', cheat_files),
        ('/etc', ['config/cheat']),
    ],
)
