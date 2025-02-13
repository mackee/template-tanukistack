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

# License

MIT
