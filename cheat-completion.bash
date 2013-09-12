#
# Autocomplete for cheat.py
# Copyright (c) 2013, Jean-Philippee "Orax" Roemer
#

function _cheat_autocomplete {
    sheets=$(cheat | tail -n +17 | cut -d' ' -f1)
    COMPREPLY=()
    if [ $COMP_CWORD = 1 ]; then
	COMPREPLY=(`compgen -W "$sheets" -- $2`)
    fi
}

complete -F _cheat_autocomplete cheat
