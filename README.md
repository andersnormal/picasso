# :art: Picasso

[![Test & Build](https://github.com/andersnormal/picasso/actions/workflows/main.yml/badge.svg)](https://github.com/andersnormal/picasso/actions/workflows/main.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/andersnormal/picasso)](https://goreportcard.com/report/github.com/andersnormal/picasso)
[![Taylor Swift](https://img.shields.io/badge/secured%20by-taylor%20swift-brightgreen.svg)](https://twitter.com/SwiftOnSecurity)
[![Volkswagen](https://auchenberg.github.io/volkswagen/volkswargen_ci.svg?v=1)](https://github.com/auchenberg/volkswagen)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Picasso is a versatile task runner :running: with a powerful plugin system.

:point_right: [Documentation](https://andersnormal.github.io/picasso/)

:warning: The task runner is under active development and the specification may change.

## Features

* Universal task runner :hammer_and_wrench:
* Extensible via plugins :partying_face:
* Template generation in task runs
* Task watcher (allows to watch on file system changes and re-runs tasks)no 

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
plugins:
  -
    id: picasso-plugin-init
    path: picasso-plugin-init
tasks:
  test:
    disabled: false
    desc: test
    vars:
      region: test
    env:
      REGION: test
    steps:
      - cmd: echo {{.OS}}
    watch:
      paths:
        - pkg/config
    template:
      - 
        file: ./examples/config.json.tpl
        out: ./config.json
        vars:
          foo: bar    
```

## Plugins

Plugins are implemented using the [go plugin over RPC](https://github.com/hashicorp/go-plugin) package.

Plugins are triggered using the `--plugin` flag. They are executed with the `vars` set with the `--var` flag and the global `vars` property in the `.picasso.yml` file.

### Example: 

See [`plugin.go`](/examples/plugin.go) for an example.

## Development

The goal is that Picasso is build and maintained by itself. However, up until this very moment. There two quick steps to build it.

The development is intended to be run with [Codespaces](https://github.com/features/codespaces) the blazing fast cloud developer environment.

```bash
env GO111MODULE=on goreleaser release --snapshot --rm-dist
```

## License

[Apache 2.0](/LICENSE)
