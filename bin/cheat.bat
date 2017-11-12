@echo off

if not defined EDITOR (set EDITOR=write)

:: Retrieve the path to python executable.
for /f "delims=" %%A in ('where python') do set "PATHOFPYTHON=%%A"
%PATHOFPYTHON% %PATHOFPYTHON%\..\Scripts\cheat %*

:: Remove this variable to avoid polluting the environment.
set PATHOFPYTHON=
