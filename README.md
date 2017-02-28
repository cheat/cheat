[![PyPI](https://img.shields.io/pypi/v/cheat.svg)](https://pypi.python.org/pypi/cheat/)
[![PyPI](https://img.shields.io/pypi/dm/cheat.svg)](https://pypi.python.org/pypi/cheat/)

cheat
=====
`cheat` allows you to create and view interactive cheatsheets on the
command-line. It was designed to help remind \*nix system administrators of
options for commands that they use frequently, but not frequently enough to
remember.

![The obligatory xkcd](http://imgs.xkcd.com/comics/tar.png 'The obligatory xkcd')


Example
-------
The next time you're forced to disarm a nuclear weapon without consulting
Google, you may run:

```sh
cheat tar
```

You will be presented with a cheatsheet resembling:

```sh
# To extract an uncompressed archive:
tar -xvf '/path/to/foo.tar'

# To extract a .gz archive:
tar -xzvf '/path/to/foo.tgz'

# To create a .gz archive:
tar -czvf '/path/to/foo.tgz' '/path/to/foo/'

# To extract a .bz2 archive:
tar -xjvf '/path/to/foo.tgz'

# To create a .bz2 archive:
tar -cjvf '/path/to/foo.tgz' '/path/to/foo/'
```

To see what cheatsheets are available, run `cheat -l`.

Note that, while `cheat` was designed primarily for \*nix system administrators,
it is agnostic as to what content it stores. If you would like to use `cheat`
to store notes on your favorite cookie recipes, feel free.


Installing
----------
It is recommended to install `cheat` with `pip`:

```sh
[sudo] pip install cheat
```

[Other installation methods are available][installing].


Modifying Cheatsheets
---------------------
The value of `cheat` is that it allows you to create your own cheatsheets - the
defaults are meant to serve only as a starting point, and can and should be
modified.

Cheatsheets are stored in the `~/.cheat/` directory, and are named on a
per-keyphrase basis. In other words, the content for the `tar` cheatsheet lives
in the `~/.cheat/tar` file.

Provided that you have a `CHEAT_EDITOR`, `VISUAL`, or `EDITOR` environment
variable set, you may edit cheatsheets with:

```sh
cheat -e foo
```

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

```sh
export DEFAULT_CHEAT_DIR='/path/to/my/cheats'
```

### Setting a CHEATPATH ###
You can additionally instruct `cheat` to look for cheatsheets in other
directories by exporting a `CHEATPATH` environment variable:

```sh
export CHEATPATH='/path/to/my/cheats'
```

You may, of course, append multiple directories to your `CHEATPATH`:

```sh
export CHEATPATH="$CHEATPATH:/path/to/more/cheats"
```

You may view which directories are on your `CHEATPATH` with `cheat -d`.

### Enabling Syntax Highlighting ###
`cheat` can apply syntax highlighting to your cheatsheets if so desired. To
enable this feature, set a `CHEATCOLORS` environment variable:

```sh
export CHEATCOLORS=true
```

Development
-----------

Virtual environments and `setuptools` can be used to develop `cheat`:

Create a virtual environment with [virtualenv](http://docs.python-guide.org/en/latest/dev/virtualenvs/):

    $ virtualenv venv

Activate it:

    $ source venv/bin/activate

Install `cheat` in the venv in [development mode](http://setuptools.readthedocs.io/en/latest/setuptools.html#development-mode):

    $ python setup.py develop

Now `cheat` can be run from inside the venv:

    $ venv/bin/cheat

See Also:
---------
- [Enabling Command-line Autocompletion][autocompletion]
- [Related Projects][related-projects]


[autocompletion]:   https://github.com/chrisallenlane/cheat/wiki/Enabling-Command-line-Autocompletion
[dotfiles]:         http://dotfiles.github.io/
[installing]:       https://github.com/chrisallenlane/cheat/wiki/Installing
[related-projects]: https://github.com/chrisallenlane/cheat/wiki/Related-Projects
