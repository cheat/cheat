"""cheat
    ~~~~~~~~

    cheat allows you to create and view interactive cheatsheets on the
    command-line. It was designed to help remind *nix system administrators of
    options for commands that they use frequently, but not frequently enough
    to remember.
    :license: GPL3
"""

from setuptools import setup, find_packages

setup(
    name         = 'cheat',
    version      = '2.1.16',
    author       = 'Chris Lane',
    author_email = 'chris@chris-allen-lane.com',
    license      = 'GPL3',
    description  = 'cheat allows you to create and view interactive cheatsheets on the command-line',
    long_description = __doc__,
    url          = 'https://github.com/chrisallenlane/cheat',
    packages     = find_packages(),
    package_data = {
        'cheat.cheatsheets': ['*'],
    },
    entry_points = {
        'console_scripts': [
            'cheat = cheat.app:main',
        ],
    },
    install_requires = [
        'docopt >= 0.6.1',
        'pygments >= 1.6.0',
    ]
)
