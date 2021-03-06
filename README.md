<p align="center">
    <a target="_blank" rel="noopener noreferrer" href="https://github.com/wkhub/wk">
        <img src="https://raw.githubusercontent.com/wkhub/wk/master/logo/logo-text-128.png"
            alt="WK Logo"
            title="WK"
            style="max-width:100%;"
        />
    </a>
</p>

[![Build Status](https://travis-ci.org/wkhub/wk.svg?branch=master)](https://travis-ci.org/wkhub/wk)

`wk` is a language agnostic project manager.

## Rational

Each language has at least en environment isolation tool but each one works its own way.
Sometimes listing and instanciating project is part of it, sometimes not.
What happen if you have to languages in the same project (or repository) ?


The idea behind `wk` is to handle the common project management tooling and delegate language specific part to each specific tool.

So `wk` allows to:
- manage (list/create/use/delete) project workspaces
- detect and activate environment if present (virtualenv, pipenv, nvm, rbenv, .env...)
- perform releases
- works with project templates
- handle global, local and shared configuration

## Installation

## Usage

```shell
export WK_PROJECTS = ~/Workspaces
wk new my-project
```
