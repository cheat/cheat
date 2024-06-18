Installing
==========
`cheat` has no runtime dependencies. As such, installing it is generally
straightforward. There are a few methods available:

### Install manually
#### Unix-like
On Unix-like systems, you may simply paste the following snippet into your terminal:

```sh
cd /tmp \
https://github.com/cheat/cheat/releases/latest/download/cheat-linux-amd64.gz \
  && gunzip cheat-linux-amd64.gz \
  && chmod +x cheat-linux-amd64 \
  && sudo mv cheat-linux-amd64 /usr/local/bin/cheat
```

You may need to need to change the archive
(`cheat-linux-amd64.gz`) depending on your platform.

See the [releases page][releases] for a list of supported platforms.

#### Windows
You can download the latest windows archive in the releases page. Now extract the archive and put "cheat-windows-amd64.exe" in the directory you wish to install it in then rename it to "cheat.exe". Finally, edit the PATH enviroment variable for your user account or the entire system to include the directory where "cheat.exe" is located.

### Install via `winget`
If you have Windows 10 or later and winget on your system, you can install `cheat` via `winget`:
```pwsh
winget install ChrisAllenLane.cheat
```

### Install via `go install`
If you have `go` version `>=1.17` available on your `PATH`, you can install
`cheat` via `go install`:

```sh
go install github.com/cheat/cheat/cmd/cheat@latest
```

### Install via package manager
Several community-maintained packages are also available:

Package manager  | Package(s)
---------------- | -----------
aur              | [cheat][pkg-aur-cheat], [cheat-bin][pkg-aur-cheat-bin]
brew             | [cheat][pkg-brew]
docker           | [docker-cheat][pkg-docker]
nix              | [nixos.cheat][pkg-nix]
snap             | [cheat][pkg-snap]

<!--[pacman][]       |-->

## Configuring
Three things must be done before you can use `cheat`:
1. A config file must be generated
2. [`cheatpaths`][cheatpaths] must be configured
3. [Community cheatsheets][community] must be downloaded

On first run, `cheat` will run an installer that will do all of the above
automatically. After the installer is complete, it is strongly advised that you
view the configuration file that was generated, as you may want to change some
of its default values (to enable colorization, change the paginator, etc).

### conf.yml ###
`cheat` is configured by a YAML file that will be auto-generated on first run.

By default, the config file is assumed to exist on an XDG-compliant
configuration path like `~/.config/cheat/conf.yml`. If you would like to store
it elsewhere, you may export a `CHEAT_CONFIG_PATH` environment variable that
specifies its path:

```sh
export CHEAT_CONFIG_PATH="~/.dotfiles/cheat/conf.yml"
```

[cheatpaths]:        README.md#cheatpaths
[community]:         https://github.com/cheat/cheatsheets/
[pkg-aur-cheat-bin]: https://aur.archlinux.org/packages/cheat-bin
[pkg-aur-cheat]:     https://aur.archlinux.org/packages/cheat
[pkg-brew]:          https://formulae.brew.sh/formula/cheat 
[pkg-docker]:        https://github.com/bannmann/docker-cheat
[pkg-nix]:           https://search.nixos.org/packages?channel=unstable&show=cheat&from=0&size=50&sort=relevance&type=packages&query=cheat 
[pkg-snap]:          https://snapcraft.io/cheat
[releases]:          https://github.com/cheat/cheat/releases
