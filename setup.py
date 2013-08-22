#!/usr/bin/env python

from distutils.core import setup
import os

setup(name='cheat',
      version='1.0',
      description='Create and view interactive cheatsheets on the command-line',
      author='Chris Lane',
      author_email='chris@chris-allen-lane.com',
      url='https://github.com/chrisallenlane/cheat',
      packages=['cheatlib'],
      package_data={'cheatlib': ['cheatsheets/*']},
      scripts=['cheat'],
     )
