#!/usr/bin/env python

from distutils.core import setup
import os

setup(name='cheat',
      version='1.0',
      summary='Create and view interactive cheatsheets on the command-line',
      homepage='',
      author='Chris Lane',
      author_email='chris@chris-allen-lane.com',
      license='GPL3',
      description='',
      platform='Posix',
      url='https://github.com/chrisallenlane/cheat',
      packages=['cheatsheets'],
      package_data={'cheatsheets': [f for f in os.listdir('cheatsheets')
                                    if '.' not in f]},
      scripts=['cheat'],
      data_files=[('/usr/share/zsh/site-functions', ['_cheat'])]
     )
