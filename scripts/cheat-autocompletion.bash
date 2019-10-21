function _cheat_autocomplete {
    wdlist='cheat -l|awk "{print $1}"'
    COMPREPLY=(`compgen -W "$wdlist"`)
}

complete -F _cheat_autocomplete cheat
