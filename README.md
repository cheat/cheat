cheat
=====
`cheat` allows you to create and view interactive cheatsheets on the
command-line. It was designed to help remind \*nix system administrators of
options for commands that they use frequently, but not frequently enough to
remember.

![The obligatory xkcd](http://imgs.xkcd.com/comics/tar.png 'The obligatory xkcd')

`cheat` depends only on `python` and `pip`.

PyPI status:

[![Latest Version](https://pypip.in/version/cheat/badge.png)](https://pypi.python.org/pypi/cheat/)
[![Downloads](https://pypip.in/download/cheat/badge.png)](https://pypi.python.org/pypi/cheat/)

Example
-------
The next time you're forced to disarm a nuclear weapon without consulting
Google, you may run:

    cheat tar

You will be presented with a cheatsheet resembling:

```
# To extract an uncompressed archive: 
tar -xvf /path/to/foo.tar

# To extract a .gz archive:
tar -xzvf /path/to/foo.tgz

# To create a .gz archive:
tar -czvf /path/to/foo.tgz /path/to/foo/

# To extract a .bz2 archive:
tar -xjvf /path/to/foo.tgz

# To create a .bz2 archive:
tar -cjvf /path/to/foo.tgz /path/to/foo/
```

To see what cheatsheets are availble, run `cheat -l`.

Note that, while `cheat` was designed primarily for *nix system administrators,
it is agnostic as to what content it stores. If you would like to use `cheat`
to store notes on your favorite cookie recipes, feel free.


Installing
----------
    
### Using pip ###

    sudo pip install cheat

### Using homebrew ###

    brew install cheat
    
### Manually ###

First install the required python dependencies with:

    sudo pip install docopt pygments

Then, clone this repository, `cd` into it, and run:

    sudo python setup.py install


Modifying Cheatsheets
---------------------
The value of `cheat` is that it allows you to create your own cheatsheets - the
defaults are meant to serve only as a starting point, and can and should be
modified.

Cheatsheets are stored in the `~/.cheat/` directory, and are named on a
per-keyphrase basis. In other words, the content for the `tar` cheatsheet lives
in the `~/.cheat/tar` file.

Provided that you have an `EDITOR` environment variable set, you may edit
cheatsheets with:

    cheat -e foo

If the 'foo' cheatsheet already exists, it will be opened for editing.
Otherwise, it will be created automatically.

After you've customized your cheatsheets, I urge you to track `~/.cheat/` along
with your [dotfiles][].


Configuring
-----------

### Setting a DEFAULT_CHEAT_DIR ###
Personal cheatsheets are saved in the `~/.cheat` directory by default, but you
can specify a different default by exporting a `DEFAULT_CHEAT_DIR` environment
variable:

    export DEFAULT_CHEAT_DIR=/path/to/my/cheats

### Setting a CHEATPATH ###
You can additionally instruct `cheat` to look for cheatsheets in other
directories by exporting a `CHEATPATH` environment variable:

    export CHEATPATH=/path/to/my/cheats

You may, of course, append multiple directories to your `CHEATPATH`:

    export CHEATPATH=$CHEATPATH:/path/to/more/cheats

You may view which directories are on your `CHEATPATH` with `cheat -d`.

### Enabling Syntax Highlighting ###
`cheat` can apply syntax highlighting to your cheatsheets if so desired. To
enable this feature, set a `CHEATCOLORS` environment variable:

    export CHEATCOLORS=true

### Enabling Command-line Autocompletion ###
The `cheat/autocompletion` directory contains scripts to enable command-line
autocompletion for various shells. To activate autocompletion, simply copy the
appropriate script to the appropriate path on your system. (The "appropriate
path" will vary on a per-platform basis, so this documentation shall not
speculate as to where that may be.)


Related Projects
----------------

- [lucaswerkmeister/cheats][1]: An implementation of this concept in pure bash
  that also allows not only for numerical indexing of subcomands but also
  supports running commands interactively.

- [jahendrie/cheat][2]: A bash-only implementation that additionally allows for
  cheatsheets to be created and `grep` searched from the command-line.
  ([jahendrie][] contributed key ideas to this project as well.)

- [`cheat` RubyGem][3]: A clever gem from 2006 that clearly had similar
  motivations. It is unclear whether or not it is currently maintained.

- [`tldr`][tldr]: "Simplified and community-driven man pages".


[dotfiles]:  http://dotfiles.github.io/
[jahendrie]: https://github.com/jahendrie
[1]:         https://github.com/lucaswerkmeister/cheats   
[2]:         https://github.com/jahendrie/cheat
[3]:         http://errtheblog.com/posts/21-cheat
[4]:         https://github.com/chrisallenlane/cheat/pull/77
[tldr]:      https://github.com/tldr-pages/tldr
