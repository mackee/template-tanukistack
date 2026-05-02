# Template of tanukistack

## What is tanukistack?

Tanukistack is a template for the development of web application. This contains the following components:

* [tanukirpc](https://github.com/mackee/tanukirpc): Web Framework
* [sqlla](https://github.com/mackee/go-sqlla): ORMapper
* [genddl](https://github.com/mackee/go-genddl): DDL Generator

## How to use

### Requirements

* `gonew` command
```sh
$ go install golang.org/x/tools/cmd/gonew@latest
```
* [aqua](https://github.com/aquaproj/aqua)
* PostgreSQL

### Create a new project

```sh
$ gonew github.com/mackee/template-tanukistack path/to/your/module ./your_directory
$ cd path/to/your/project
$ aqua i
```

### Settings

`.env` file is required to run the application. Create a `.env` file in the root directory of the project and write the following contents.

```sh
$ echo 'DATABASE_DSN="postgres://user:password@localhost:5432/dbname?sslmode=disable"' > .env
```

### Run the application

```sh
$ task dev
```

### Run the test

```sh
$ task test
```

## Use with devcontainer

A devcontainer setup is provided so you can develop without installing Go, task, or PostgreSQL on the host.

### Requirements

* Docker
* [devcontainer CLI](https://github.com/devcontainers/cli)
  ```sh
  $ brew install --cask devcontainer
  # or: npm install -g @devcontainers/cli
  ```

### Usage

```sh
$ task dc:up                       # build & start the container (first run takes a while)
$ task dc:exec -- task test        # run tests against the postgres sidecar
$ task dc:exec -- task dev         # run the dev server (forwarded to host :8080)
$ task dc:shell                    # open a shell inside the container
$ task dc:down                     # stop & remove containers (volumes are kept)
```

### Notes

* `DATABASE_DSN` is set in `containerEnv` to point at the bundled postgres sidecar, so a host `.env` is not required.
* aqua-managed tools (`go`, `task`) are installed on `postCreate` and cached in a named volume.
* The postgres data volume (`pgdata`) is preserved across `dc:down` / `dc:up`.

# License

MIT
