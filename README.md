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


Advanced Features
=================

Setting a CHEATPATH
----------------------
By default, all cheatsheets are installed to `~/.cheat/`, but you can instruct
`cheat` to look for cheatsheets in other directories by exporting a `CHEATPATH`
environment variable:

```bash
export CHEATPATH=/path/to/my/cheats
```

You may, of course, append multiple directories to your `CHEATPATH`:

```bash
export CHEATPATH=$CHEATPATH:/path/to/more/cheats
```

Enabling Syntax Highlighting
----------------------------
`cheat` can apply syntax highlighting to your cheatsheets if so desired. To
enable this feature, set a `$CHEATCOLORS` environment variable:

```bash
export CHEATCOLORS=true
```


Contributing
============
If you would like to contribute cheetsheets or program functionality, please
fork this repository, make your chanages, and send me a pull request.

Related Projects
================

- [lucaswerkmeister/cheats][1]: An implementation of this concept in pure bash
  that also allows not only for numerical indexing of subcomands but also
  supports running commands interactively.

- [jahendrie/cheat][2]: A bash-only implmentation that additionally allows for
  cheatsheets to be created and `grep` searched from the command-line.
  ([jahendrie][] contributed key ideas to this project as well.)

- [`cheat` RubyGem][3]: A clever gem from 2006 that clearly had similar
  motivations. It is unclear whether or not it is currently maintained.


[dotfiles]:  http://dotfiles.github.io/
[jahendrie]: https://github.com/jahendrie
[1]:         https://github.com/lucaswerkmeister/cheats   
[2]:         https://github.com/jahendrie/cheat
[3]:         http://errtheblog.com/posts/21-cheat
