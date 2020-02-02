# cheat(1) completion                                      -*- shell-script -*-

# generate cheatsheet completions, optionally using `fzf`
_cheat_complete_cheatsheets()
{
  if [[ "$CHEAT_USE_FZF" = true ]]; then
    FZF_COMPLETION_TRIGGER='' _fzf_complete "--no-multi" "$@" < <(
      cheat -l | tail -n +2 | cut -d' ' -f1
    )
  else
    COMPREPLY=( $(compgen -W "$(cheat -l | tail -n +2 | cut -d' ' -f1)" -- "$cur") )
  fi
}

# generate tag completions, optionally using `fzf`
_cheat_complete_tags()
{
  if [ "$CHEAT_USE_FZF" = true ]; then
    FZF_COMPLETION_TRIGGER='' _fzf_complete "--no-multi" "$@" < <(cheat -T)
  else
    COMPREPLY=( $(compgen -W "$(cheat -T)" -- "$cur") )
  fi
}

# implement the `cheat` autocompletions
_cheat()
{
  local cur prev words cword split
  _init_completion -s || return

  # complete options that are currently being typed: `--col` => `--colorize`
  if [[ $cur == -* ]]; then
    COMPREPLY=( $(compgen -W '$(_parse_help "$1" | sed "s/=//g")' -- "$cur") )
    [[ $COMPREPLY == *= ]] && compopt -o nospace
    return
  fi

  # implement completions
  case $prev in
    --colorize|-c|\
    --directories|-d|\
    --init|\
    --regex|-r|\
    --search|-s|\
    --tags|-T|\
    --version|-v)
      # noop the above, which should implement no completions
      ;;
    --edit|-e)
      _cheat_complete_cheatsheets
      ;;
    --list|-l)
      _cheat_complete_cheatsheets
      ;;
    --path|-p)
      COMPREPLY=( $(compgen -W "$(cheat -d | cut -d':' -f1)" -- "$cur") )  
      ;;
    --rm)
      _cheat_complete_cheatsheets
      ;;
    --tag|-t)
      _cheat_complete_tags
      ;;
    *)
      _cheat_complete_cheatsheets
      ;;
  esac

  $split && return

} &&
complete -F _cheat cheat

# ex: filetype=sh
