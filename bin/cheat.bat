@echo off

if not defined CHEAT_EDITOR (set CHEAT_EDITOR=notepad)

REM %~dp0 is black magic for getting directory of script
python %~dp0\cheat %*
