cheat
=====
`cheat` allows you to create and view interactive cheatsheets on the
command-line. It was designed to help remind \*nix system administrators of
options for commands that they use frequently, but not frequently enough to
remember.

![The obligatory xkcd](http://imgs.xkcd.com/comics/tar.png 'The obligatory xkcd')

`cheat` depends only on python.


Examples
========
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
==========

### Installing for all users (requires root)

Clone this repository and `cd` into it, then run

    sudo python setup.py install

### Installing in your home directory

Clone this repository and `cd` into it, then run

    mkdir -p ~/bin
    cp cheat ~/bin
    mkdir ~/.cheat
    cp cheatsheets/* ~/.cheat

Modifying Cheatsheets
=====================
The value of `cheat` is that it allows you to create your own cheatsheets - the
defaults are meant to serve only as a starting point, and can and should be
modified.

Cheatsheets are stored in the `~/.cheat/` directory, and are named on a
per-keyphrase basis. In other words, the content for the `tar` cheatsheet lives
in the `~/.cheat/tar` file. To add a cheatsheet for a `foo` command, you would
create file `~/.cheat/foo`, whereby that file contained the cheatsheet content.

Note that `cheat` supports "subcommands" simply by naming files appropriately.
Thus, if you wanted to create a cheatsheet not only (for example) for `git` but
also for `git commit`, you could do so be creating cheatsheet files of the
appropriate names (`git` and `git commit`).

After you've customized your cheatsheets, I urge you to track `~/.cheat/` along
with your [dotfiles][].


Contributing
============
If you would like to contribute additional cheatsheets for basic \*nix
commands, please modify the `.cheat` file and send me a pull request.


[dotfiles]: http://dotfiles.github.io/
