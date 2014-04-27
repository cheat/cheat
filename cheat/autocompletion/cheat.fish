#completion for cheat
complete -c cheat -s h -l help -f -x --description "Display help and exit"
complete -c cheat -l edit -f -x --description "Edit <cheatsheet>"
complete -c cheat -s e -f -x --description "Edit <cheatsheet>"
complete -c cheat -s l -l list -f -x --description "List all available cheatsheets"
complete -c cheat -s d -l cheat-directories -f -x --description "List all current cheat dirs"
complete -c cheat --authoritative -f
for cheatsheet in (cheat -l | cut -d' ' -f1)
    complete -c cheat -a "$cheatsheet"
    complete -c cheat -o e -a "$cheatsheet" 
    complete -c cheat -o '-edit' -a "$cheatsheet"
end
