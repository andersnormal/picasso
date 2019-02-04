# :art: Picasso

Picasso is a versatile task runner :running: and awesome build tool.

:see_no_evil: Contributions are welcome. 

## Features

* Universal task runner
* Template generation in task runs
* Task watcher (allows to watch on file system changes and re-runs tasks)

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
env GO111MODULE=on mkdir -p bin && go build -i -o bin/picasso && chmod +x bin/picasso
```

## License
[Apache 2.0](/LICENSE)
