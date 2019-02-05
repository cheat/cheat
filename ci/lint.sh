#!/bin/sh

# Resolve the app root
SCRIPT=`realpath $0`
SCRIPTPATH=`dirname $SCRIPT`
APPROOT=`realpath "$SCRIPTPATH/.."`

flake8 $APPROOT/setup.py
flake8 $APPROOT/bin/cheat
flake8 $APPROOT/cheat/*.py
