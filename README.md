cheat
=====
`cheat` allows you to create and view interactive cheatsheets on the
command-line. It was designed to help \*nix system administrators remember
options for commands that they use frequently, but not frequently enough to
remember.

![The obligatory xkcd](http://imgs.xkcd.com/comics/tar.png 'The obligatory xkcd')

`cheat` depends only on python.

Examples
--------
The next time you're forced to disarm a nuclear weapon without consulting
Google, you may run:

```sh
cheat tar
```

You will be presented with a cheatsheet resembling:

```text
To extract an uncompressed archive: 
tar -xvf /path/to/foo.tar

To extract a .gz archive:
tar -xzvf /path/to/foo.tgz

To create a .gz archive:
tar -czvf /path/to/foo.tgz /path/to/foo/

To extract a .bz2 archive:
tar -xjvf /path/to/foo.tgz

To create a .bz2 archive:
tar -cjvf /path/to/foo.tgz /path/to/foo/
```

To see what cheatsheets are availble, run `cheat` with no arguments.

Note that, while `cheat` was designed primarily for *nix system administrators,
it is agnostic as to what content it stores. If you would like to use `cheat`
to store notes on your favorite cookie recipes, feel free.


Installing
----------
Do the following to install `cheat`:

1. Clone this repository and `cd` into it
2. Run `sudo ./install`

The `install` script will copy a python file into `/usr/local/bin/`, and will
also create a hidden file (containing the cheatsheet content) in your home
directory.


Modifying Cheatsheets
---------------------
The value of `cheat` is that it allows you to create your own cheatsheets - the
defaults are meant to serve only as a starting point, and can and should be
modified.

To modify your cheatsheets, edit the `~/.cheat` file, which simply contains a
python dictionary. To add new cheatsheets, you need only append new key/value
pairs to the dictionary.

Note that `cheat` supports subcommands, such that (for example) `git` and `git
commit` may each be assigned their own cheatsheets.

After you've customized your cheatsheets, I urge you to track `.cheat` along
with your [dotfiles][].


Contributing
------------
If you would like to contribute additional cheatsheets for basic \*nix
commands, please modify the `.cheat` file and send me a pull request.


[dotfiles]: http://dotfiles.github.io/
