#!/usr/bin/env python

from distutils.core import setup
import os

setup(name='cheat',
      version='1.0',
      author='Chris Lane',
      author_email='chris@chris-allen-lane.com',
      license='GPL3',
      description='cheat allows you to create and view interactive cheatsheets\
      on the command-line. It was designed to help remind *nix system\
      administrators of options for commands that they use frequently, but not\
      frequently enough to remember.',
      url='https://github.com/chrisallenlane/cheat',
      packages=['cheatsheets'],
      package_data={'cheatsheets': [f for f in os.listdir('cheatsheets')
                                    if '.' not in f]},
      scripts=['cheat'],
      data_files=[('/usr/share/zsh/site-functions', ['autocompletion/_cheat.zsh']),
                  ('/etc/bash_completion.d'       , ['autocompletion/cheat.bash']),
                  ('/usr/share/fish/completions'  , ['autocompletion/cheat.fish'])
              ]
      )
