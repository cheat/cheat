#compdef cheat

local cheats taglist pathlist

_cheat_complete_personal_cheatsheets()
{
  cheats=("${(f)$(cheat -l -t personal | tail -n +2 | cut -d' ' -f1)}")
  _describe -t cheats 'cheats' cheats
}

_cheat_complete_full_cheatsheets()
{
  cheats=("${(f)$(cheat -l | tail -n +2 | cut -d' ' -f1)}")
  _describe -t cheats 'cheats' cheats
}

_cheat_complete_tags()
{
  taglist=("${(f)$(cheat -T)}")
  _describe -t taglist 'taglist' taglist
}

_cheat_complete_paths()
{
  pathlist=("${(f)$(cheat -d | cut -d':' -f1)}")
  _describe -t pathlist 'pathlist' pathlist
}

_cheat() {

  _arguments -C \
    '(--init)--init[Write a default config file to stdout]: :->none' \
    '(-c --colorize)'{-c,--colorize}'[Colorize output]: :->none' \
    '(-d --directories)'{-d,--directories}'[List cheatsheet directories]: :->none' \
    '(-e --edit)'{-e,--edit}'[Edit <sheet>]: :->personal' \
    '(-l --list)'{-l,--list}'[List cheatsheets]: :->full' \
    '(-p --path)'{-p,--path}'[Return only sheets found on path <name>]: :->pathlist' \
    '(-r --regex)'{-r,--regex}'[Treat search <phrase> as a regex]: :->none' \
    '(-s --search)'{-s,--search}'[Search cheatsheets for <phrase>]: :->none' \
    '(-t --tag)'{-t,--tag}'[Return only sheets matching <tag>]: :->taglist' \
    '(-T --tags)'{-T,--tags}'[List all tags in use]: :->none' \
    '(-v --version)'{-v,--version}'[Print the version number]: :->none' \
    '(--rm)--rm[Remove (delete) <sheet>]: :->personal' 

  case $state in
    (none)
      ;;
    (full)
      _cheat_complete_full_cheatsheets
      ;;
    (personal)
      _cheat_complete_personal_cheatsheets
      ;;
    (taglist)
      _cheat_complete_tags
      ;;
    (pathlist)
      _cheat_complete_paths
      ;;
    (*)
      ;;
  esac
}

compdef _cheat cheat
