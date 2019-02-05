@echo OFF

if not defined CHEAT_EDITOR if not defined EDITOR if not defined VISUAL (
    set CHEAT_EDITOR=write
)

REM %~dp0 is black magic for getting directory of script
python %~dp0cheat %*
