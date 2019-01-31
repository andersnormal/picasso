# :art: Picasso

Picasso is a versatile task runner :running:.

## Features

* Universal task runner
* Template generation in task runs
* Task watcher (allows to watch on file system changes and re-runs tasks)
* Scaffolding of projects

## Example

```yaml
version: 1
author: demo demo@example.com
project: demo
tasks:
  test:
    desc: run tests
    cmds:
      - go test -v ./...
  dev:
    desc: build and watch
    deps:
      - test
    vars:
      region:
        - test
    cmds:
      - go test -v ./...
      - go build
    watch:
      paths:
        - examples
      ignore:
        - .gitignore
        - .picasso.yml
    templates:
      - 
        file: ./examples/config.json.tpl
        output: ./config.json
        vars:
          foo: bar
```

## Development

The goal is that Picasso is build and maintained by itself. However, up until this very moment. There two quick steps to build it.

```
packr2 build && go build
```

## License
[Apache 2.0](/LICENSE)
