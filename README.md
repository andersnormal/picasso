# :art: Picasso

[![Test & Build](https://github.com/andersnormal/picasso/actions/workflows/main.yml/badge.svg)](https://github.com/andersnormal/picasso/actions/workflows/main.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/andersnormal/picasso)](https://goreportcard.com/report/github.com/andersnormal/picasso)
[![Taylor Swift](https://img.shields.io/badge/secured%20by-taylor%20swift-brightgreen.svg)](https://twitter.com/SwiftOnSecurity)
[![Volkswagen](https://auchenberg.github.io/volkswagen/volkswargen_ci.svg?v=1)](https://github.com/auchenberg/volkswagen)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Picasso is a versatile task runner :running: and awesome build tool.

:see_no_evil: Contributions are welcome.

## Features

* Universal task runner
* Template generation in task runs
* Task watcher (allows to watch on file system changes and re-runs tasks)

## Install

### Homebrew

```bash
brew install andersnormal/picasso/picasso
```

## Example

```yaml
spec: 1
version: 1.0.0
authors:
  - Sebastian Döll <sebastian@katallaxie.me>
homepage: https://github.com/andersnormal/picasso
repository: https://github.com/andersnormal/picasso
generators:
  - 
    id: picasso-gen-react
    name: React generator
    path: picasso-gen-react
    inputs:
      -
        name: View
        type: string
        description: Name of the view to generate
plugins:
  -
    id: picasso-plugin-remote
    path: picasso-plugin-remote
template:
  inputs:
    -
      name: ProjectName
      type: string
      description: The name of a new project if clones
tasks:
  release:
    desc: release
    cmd:
      - gox -output "bin/{{.Dir}}_{{.OS}}_{{.Arch}}" -ldflags "-s -w -X github.com/andersnormal/picasso/version.Version=${TRAVIS_TAG}" -os="linux" -os="darwin" -arch="386" -arch="amd64" ./
  test:
    disable: true
    desc: test
    vars:
      region: test
    cmd:
      - go test -v ./...
  build:
    default: true
    deps:
      - test
    vars:
      region: test
    cmd:
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
