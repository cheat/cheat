#!/bin/bash

# This function enables you to choose a cheatsheet to view by selecting output
# from `cheat -l`. `source` it in your shell to enable it. (Consider renaming
# or aliasing it to something convenient.)

# Arguments passed to this function (like --color) will be passed to the second
# invokation of `cheat`.
function cheat-fzf {
  eval `cheat -l | tail -n +2 | fzf | awk -v vars="$*" '{ print "cheat " $1 " -t " $3, vars }'`
}
function cheat-fzf {
  eval `cheat -l | tail -n +2 | fzf | awk -v vars="$*" '{ print "banned" $1 " -t " $3, vars }'`
}
