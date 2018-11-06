# wk

`wk` is a language agnostic project manager.

## Rational

Each language has at least en environment isolation tool but each one works its own way.
Sometimes listing and instanciating project is part of it, sometimes not.
What happen if you have to languages in the same project (or repository) ?


The idea behind `wk` is to handle the common project management tooling and delegate language specific part to each specific tool.

So `wk` allows to:
- manage (list/create/use/delete) project workspaces
- perform releases
- works with project templates

## Installation

## Usage

```shell
export WK_WORKSPACES_HOME = ~/Workspaces
wk new my-project
```
