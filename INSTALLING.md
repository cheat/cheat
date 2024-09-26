# Installing

`cheat` has no runtime dependencies. As such, installing it is generally
straightforward. There are a few methods available:

## Install manually
### Unix-like
On Unix-like systems, you may simply paste the following snippet into your terminal:

```sh
cd /tmp \
  && wget https://github.com/cheat/cheat/releases/download/4.4.2/cheat-linux-amd64.gz \
  && gunzip cheat-linux-amd64.gz \
  && chmod +x cheat-linux-amd64 \
  && sudo mv cheat-linux-amd64 /usr/local/bin/cheat
```

You may need to need to change the version number (`4.4.2`) and the archive
(`cheat-linux-amd64.gz`) depending on your platform.

See the [releases page][releases] for a list of supported platforms.

<br>

### Windows
Head over to the [Releases](https://github.com/cheat/cheat/releases/latest) page and download `cheat-windows-amd64.exe.zip`.

1. **Extract the archive**:  
   Extract the executable into your local appdata directory:
   ```powershell
   Expand-Archive .\cheat-windows-amd64.exe.zip -DestinationPath $env:LOCALAPPDATA\cheat
   ```
   **Note**: You can install `cheat` to any other location if preferred, simply substitute your chosen installation path in steps 2 and 3 below.

2. **(Optional) Create a symbolic link**:  
   To facilitate invoking `cheat` via `cheat.exe`, you can create a symbolic link:
   ```powershell
   saps -v runas cmd -args "/c mklink %LOCALAPPDATA%\cheat\cheat.exe %LOCALAPPDATA%\cheat\cheat-windows-amd64.exe"
   ```

3. **Add `cheat` to your PATH**:  
   Check if `cheat` is already in your PATH, and if not, append it:
   ```powershell
   if (-not (("$env:PATH" -cmatch "\\cheat") -or "$(where.exe cheat.exe 2>$null)")) {
     $NEWPATH = "$env:PATH;$env:LOCALAPPDATA\cheat";
     write-host $NEWPATH;
     saps -v runas cmd -args "setx /M PATH `"$NEWPATH`""
   } else {
     write-host 'Cheat is already found on your PATH, skipping...'
   }
   ```

4. **Restart your PC or shell**:  
   To apply the PATH changes, either restart your computer, or close and reopen any terminal/PowerShell sessions.

5. **Install Verification**:  
   In your __freshly opened__ `powershell` or `CMD` terminal, invoke `cheat`:
   ```powershell
   cheat
   ```
   If the install was successful, you should see the initial run dialog. This will present you with the option to generate a new config file and download community cheatsheets.

  **Troubleshooting**

  If `cheat` isn’t running after installation:

- **Verify the PATH was applied correctly**:
  Verify your `cheat` path is present in the system environment variables. If it’s missing, try step 3 again to add it to your PATH.

- **Loosen ExecutionPolicy**:
  If at any point you run into issues with PowerShell’s script execution policy, you can temporarily relax it:
  ```powershell
  Set-ExecutionPolicy -Scope Process -ExecutionPolicy RemoteSigned
  ```
  This will let you run scripts for the current shell session without making permanent changes to the system policy.

- **Try another installation method**: If you're still having problems after trying the above, you can try installing `cheat` via `docker` or `go install`.

<br>

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

<br>

## Configuring
Three things must be done before you can use `cheat`:
1. A config file must be generated
2. [`cheatpaths`][cheatpaths] must be configured
3. [Community cheatsheets][community] must be downloaded

On first run, `cheat` will run an installer that will do all of the above
automatically. After the installer is complete, it is strongly advised that you
view the configuration file that was generated, as you may want to change some
of its default values (to enable colorization, change the paginator, etc).

<br>

## conf.yml
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
