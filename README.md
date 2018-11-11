git-update-commit-status
===

A simple tool to update commit status of GitHub


Installation
====

Requirements: Go 1.11 or higher

```
$ export GO111MODULE=on
$ go get -u github.com/pocke/git-update-commit-status
```

Usage
===

`git update-commit-status STATUS [revision]`

```
$ git update-commit-status success
$ git update-commit-status failure some-branch-name
```

License
-------

These codes are licensed under CC0.

[![CC0](http://i.creativecommons.org/p/zero/1.0/88x31.png "CC0")](http://creativecommons.org/publicdomain/zero/1.0/deed.en)
