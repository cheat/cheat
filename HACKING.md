Hacking
=======
The following is a quickstart guide for developing `cheat`.

## 1. Install system dependencies
Before you begin, you must install a handful of system dependencies. The
following are required, and must be available on your `PATH`:

- `git`
- `go` (>= 1.17 is recommended)
- `make`

The following dependencies are optional:
- `docker`
- `pandoc` (necessary to generate a `man` page)

## 2. Install utility applications
Run `make setup` to install `scc` and `revive`, which are used by various
`make` targets.

## 3. Development workflow
After your environment has been configured, your development workflow will
resemble the following:

1. Make changes to the `cheat` source code.
2. Run `make test` to run unit-tests.
3. Fix compiler errors and failing tests as necessary.
4. Run `make`. A `cheat` executable will be written to the `dist` directory.
5. Use the new executable by running `dist/cheat <command>`.
6. Run `make install` to install `cheat` to your `PATH`.
7. Run `make build-release` to build cross-platform binaries in `dist`.
8. Run `make clean` to clean the `dist` directory when desired.

You may run `make help` to see a list of available `make` commands.

### Developing with docker
It may be useful to test your changes within a pristine environment. An
Alpine-based docker container has been provided for that purpose.

If you would like to build the docker container, run:
```sh
make docker-setup
```

To shell into the container, run:
```sh
make docker-sh
```

The `cheat` source code will be mounted at `/app` within the container.

If you would like to destroy this container, you may run:
```sh
make distclean
```

[go]: https://go.dev/
