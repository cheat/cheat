# Running emacs

  GUI mode            $ emacs
  Terminal mode       $ emacs -nw

# Basic usage

  Indent              Select text then press TAB
  Cut                 CTRL-w
  Copy                ALT-w
  Paste ("yank")      CTRL-y
  Begin selection     CTRL-SPACE
  Search/Find         CTRL-s
  Replace             ALT-% (ALT-SHIFT-5)
  Save                CTRL-x CTRL-s
  Save as             CTRL-x CTRL-w
  Load/Open           CTRL-x CTRL-f
  Undo                CTRL-x u
  Highlight all text  CTRL-x h
  Directory listing   CTRL-x d
  Cancel a command    CTRL-g
  Font size bigger    CTRL-x CTRL-+
  Font size smaller   CTRL-x CTRL--

# Buffers

  Split screen vertically                         CTRL-x 2
  Split screen vertically with 5 row height       CTRL-u 5 CTRL-x 2
  Split screen horizontally                       CTRL-x 3
  Split screen horizontally with 24 column width  CTRL-u 24 CTRL-x 3
  Revert to single screen                         CTRL-x 1
  Hide the current screen                         CTRL-x 0
  Move to the next screen                         CTRL-x o
  Kill the current buffer                         CTRL-x k
  Select a buffer                                 CTRL-x b
  Run command in the scratch buffer               CTRL-x CTRL-e

# Navigation ( backward / forward )
  
  Character-wise                                  CTRL-b , CTRL-f
  Word-wise                                       ALT-b  , ALT-f
  Line-wise                                       CTRL-p , CTRL-n
  Sentence-wise                                   ALT-a  , ALT-e
  Paragraph-wise                                  ALT-{ , ALT-}
  Function-wise                                   CTRL-ALT-a , CTRL-ALT-e
  Line beginning / end                            CTRL-a , CTRL-e

# Other stuff

  Open a shell         ALT-x eshell
  Goto a line number   ALT-x goto-line
  Word wrap            ALT-x toggle-word-wrap
  Spell checking       ALT-x flyspell-mode
  Line numbers         ALT-x linum-mode
  Toggle line wrap     ALT-x visual-line-mode
  Compile some code    ALT-x compile
  List packages        ALT-x package-list-packages

# Line numbers

  To add line numbers and enable moving to a line with CTRL-l:

    (global-set-key "\C-l" 'goto-line)
    (add-hook 'find-file-hook (lambda () (linum-mode 1)))
