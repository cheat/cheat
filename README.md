![Workflow status](https://github.com/cheat/cheat/workflows/build/badge.svg)

cheat
=====

`cheat` allows you to create and view interactive cheatsheets on the
command-line. It was designed to help remind \*nix system administrators of
options for commands that they use frequently, but not frequently enough to
remember.

![The obligatory xkcd](http://imgs.xkcd.com/comics/tar.png 'The obligatory xkcd')

Use `cheat` with [cheatsheets][].


Example
-------
The next time you're forced to disarm a nuclear weapon without consulting
Google, you may run:

```sh
cheat tar
```

You will be presented with a cheatsheet resembling the following:

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


Installing
----------
`cheat` has no dependencies. To install it, download the executable from the
[releases][] page and place it on your `PATH`.

Alternatively, if you have [go][] installed, you may install `cheat` using `go
get`:

```sh
go get -u github.com/cheat/cheat/cmd/cheat
```

Configuring
-----------
### conf.yml ###
`cheat` is configured by a YAML file that will be auto-generated on first run.

By default, the config file is assumed to exist on an XDG-compliant
configuration path like `~/.config/cheat/conf.yml`. If you would like to store
it elsewhere, you may export a `CHEAT_CONFIG_PATH` environment variable that
specifies its path:

```sh
export CHEAT_CONFIG_PATH="~/.dotfiles/cheat/conf.yml"
```

Cheatsheets
-----------
Cheatsheets are plain-text files with no file extension, and are named
according to the command used to view them:

```sh
cheat tar     # file is named "tar"
cheat foo/bar # file is named "bar", in a "foo" subdirectory
```

Cheatsheet text may optionally be preceeded by a YAML frontmatter header that
assigns tags and specifies syntax:

```
---
syntax: javascript
tags: [ array, map ]
---
// To map over an array:
const squares = [1, 2, 3, 4].map(x => x * x);
```

The `cheat` executable includes no cheatsheets, but [community-sourced
cheatsheets are available][cheatsheets]. You will be asked if you would like to
install the community-sourced cheatsheets the first time you run `cheat`.


Cheatpaths
----------
Cheatsheets are stored on "cheatpaths", which are directories that contain
cheatsheets. Cheatpaths are specified in the `conf.yml` file.

It can be useful to configure `cheat` against multiple cheatpaths. A common
pattern is to store cheatsheets from multiple repositories on individual
cheatpaths:

```yaml
# conf.yml:
# ...
cheatpaths:
  - name: community                   # a name for the cheatpath
    path: ~/documents/cheat/community # the path's location on the filesystem
    tags: [ community ]               # these tags will be applied to all sheets on the path
    readonly: true                    # if true, `cheat` will not create new cheatsheets here

  - name: personal
    path: ~/documents/cheat/personal  # this is a separate directory and repository than above
    tags: [ personal ]
    readonly: false                   # new sheets may be written here
# ...
```

The `readonly` option instructs `cheat` not to edit (or create) any cheatsheets
on the path. This is useful to prevent merge-conflicts from arising on upstream
cheatsheet repositories.

If a user attempts to edit a cheatsheet on a read-only cheatpath, `cheat` will
transparently copy that sheet to a writeable directory before opening it for
editing.

### Directory-scoped Cheatpaths ###
At times, it can be useful to closely associate cheatsheets with a directory on
your filesystem. `cheat` facilitates this by searching for a `.cheat` folder in
the current working directory. If found, the `.cheat` directory will
(temporarily) be added to the cheatpaths.

Usage
-----
To view a cheatsheet:

```sh
cheat tar      # a "top-level" cheatsheet
cheat foo/bar  # a "nested" cheatsheet
```

To edit a cheatsheet:

```sh
cheat -e tar     # opens the "tar" cheatsheet for editing, or creates it if it does not exist
cheat -e foo/bar # nested cheatsheets are accessed like this
```

To view the configured cheatpaths:

```sh
cheat -d
```

To list all available cheatsheets:

```sh
cheat -l
```

To list all cheatsheets that are tagged with "networking":

```sh
cheat -l -t networking
```

To list all cheatsheets on the "personal" path:

```sh
cheat -l -p personal
```

To search for the phrase "ssh" among cheatsheets:

```sh
cheat -s ssh
```

To search (by regex) for cheatsheets that contain an IP address:

```sh
cheat -r -s '(?:[0-9]{1,3}\.){3}[0-9]{1,3}'
```

Flags may be combined in intuitive ways. Example: to search sheets on the
"personal" cheatpath that are tagged with "networking" and match a regex:

```sh
cheat -p personal -t networking --regex -s '(?:[0-9]{1,3}\.){3}[0-9]{1,3}'
```


Advanced Usage
--------------
Shell autocompletion is currently available for `bash`, `fish`, and `zsh`. Copy
the relevant [completion script][completions] into the appropriate directory on
your filesystem to enable autocompletion. (This directory will vary depending
on operating system and shell specifics.)

Additionally, `cheat` supports enhanced autocompletion via integration with
[fzf][]. To enable `fzf` integration:

1. Ensure that `fzf` is available on your `$PATH`
2. Set an envvar: `export CHEAT_USE_FZF=true`

[Releases]:    https://github.com/cheat/cheat/releases
[cheatsheets]: https://github.com/cheat/cheatsheets
[completions]: https://github.com/cheat/cheat/tree/master/scripts
[fzf]:         https://github.com/junegunn/fzf
[go]:          https://golang.org
