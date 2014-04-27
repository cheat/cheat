#compdef cheat

declare -a cheats
cheats=$(cheat -l | cut -d' ' -f1)
_arguments "1:cheats:(${cheats})" && return 0
