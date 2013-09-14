#compdef cheat

declare -a cheats cheats_in_this_dir

for directory in $(cheat --cheat_directories); do
    cheats_in_this_dir=($directory/*(N:r:t))
    [[ ${#cheats_in_this_dir} > 0 ]] && cheats+=($cheats_in_this_dir)
done

_arguments "1:cheats:(${cheats})"

return 1
