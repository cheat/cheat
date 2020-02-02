_cheat_cheatpaths()
{
  if [[ -z "${cur// }" ]]; then
    compgen -W "$(cheat -d | cut -d':' -f1)"
  else
    compgen -W "$(cheat -d | cut -d':' -f1 | grep $cur)"
  fi
}

_cheat_cheatsheets()
{
  if [[ "${CHEAT_USE_FZF:-0}" != 0 ]]; then
    FZF_COMPLETION_TRIGGER='' _fzf_complete "--no-multi" "$@" < <(cheat -l | tail -n +2 | cut -d' ' -f1)
  else
    COMPREPLY=( $(compgen -W "$(cheat -l $cur | tail -n +2 | cut -d' ' -f1)") )
  fi
}

_cheat_tags()
{
  if [[ -z "${cur// }" ]]; then
    compgen -W "$(cheat -T)"
  else
    compgen -W "$(cheat -T | grep $cur)"
  fi
}

_cheat()
{
  local cur prev words cword split
  _init_completion -s || return

  case $prev in
    --edit|-e)
      _cheat_cheatsheets
      ;;
    --path|-p)
      COMPREPLY=( $(_cheat_cheatpaths) )
      ;;
    --rm)
      _cheat_cheatsheets
      ;;
    --tags|-t)
      COMPREPLY=( $(_cheat_tags) )
      ;;
  esac

  $split && return

  if [[ ! "$cur" =~ ^-.*  ]]; then
      _cheat_cheatsheets
      return
  fi

  if [[ $cur == -* ]]; then
    COMPREPLY=( $(compgen -W '$(_parse_help "$1")' -- "$cur") )
    [[ $COMPREPLY == *= ]] && compopt -o nospace
    return
  fi

} &&
complete -F _cheat cheat

# ex: filetype=sh
