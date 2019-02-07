# Running emacs

  GUI mode            $ emacs
  Terminal mode       $ emacs -nw

# Basic usage

  Indent              Select text then press TAB
  Cut                 C-w
  Copy                M-w
  Paste ("yank")      C-y
  Begin selection     C-SPACE
  Search/Find         C-s
  Replace             M-% (M-SHIFT-5)
  Save                C-x C-s
  Save as             C-x C-w
  Load/Open           C-x C-f
  Undo                C-x u
  Highlight all text  C-x h
  Directory listing   C-x d
  Cancel a command    C-g
  Font size bigger    C-x C-+
  Font size smaller   C-x C--

# Buffers

  Split screen vertically                         C-x 2
  Split screen vertically with 5 row height       C-u 5 C-x 2
  Split screen horizontally                       C-x 3
  Split screen horizontally with 24 column width  C-u 24 C-x 3
  Revert to single screen                         C-x 1
  Hide the current screen                         C-x 0
  Move to the next screen                         C-x o
  Kill the current buffer                         C-x k
  Select a buffer                                 C-x b
  Run command in the scratch buffer               C-x C-e

# Navigation ( backward / forward )
  
  Character-wise                                  C-b , C-f
  Word-wise                                       M-b  , M-f
  Line-wise                                       C-p , C-n
  Sentence-wise                                   M-a  , M-e
  Paragraph-wise                                  M-{ , M-}
  Function-wise                                   C-M-a , C-M-e
  Line beginning / end                            C-a , C-e

# Other stuff

  Open a shell         M-x eshell
  Goto a line number   M-x goto-line
  Word wrap            M-x toggle-word-wrap
  Spell checking       M-x flyspell-mode
  Line numbers         M-x linum-mode
  Toggle line wrap     M-x visual-line-mode
  Compile some code    M-x compile
  List packages        M-x package-list-packages

# Line numbers

  To add line numbers and enable moving to a line with C-l:

    (global-set-key "\C-l" 'goto-line)
    (add-hook 'find-file-hook (lambda () (linum-mode 1)))
