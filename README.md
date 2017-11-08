# :hammer: dig [![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov] [![Go Report Card][report-card-img]][report-card]

A reflection based dependency injection toolkit for Go.

Making incompatible changes from the original code at https://github.com/uber-go/dig.

### Good for:

* Powering an application framework, e.g. [Fx](https://github.com/uber-go/fx).
* Resolving the object graph during process startup.

### Bad for:

* Using in place of an application framework, e.g. [Fx](https://github.com/uber-go/fx).
* Resolving dependencies after the process has already started.
* Exposing to user-land code as a [Service Locator](https://martinfowler.com/articles/injection.html#UsingAServiceLocator).

## Installation

We recommend locking to [SemVer](http://semver.org/) range `^1` using [Glide](https://github.com/Masterminds/glide):

```
glide get 'github.com/anuvu/dig#^1'
```

## Stability

This library is `v1` and follows [SemVer](http://semver.org/) strictly.

No breaking changes will be made to exported APIs before `v2.0.0`.

[doc-img]: http://img.shields.io/badge/GoDoc-Reference-blue.svg
[doc]: https://godoc.org/github.com/anuvu/dig

[ci-img]: https://img.shields.io/travis/anuvu/dig/master.svg
[ci]: https://travis-ci.org/anuvu/dig/branches

[cov-img]: https://codecov.io/gh/anuvu/dig/branch/master/graph/badge.svg
[cov]: https://codecov.io/gh/anuvu/dig/branch/master

[report-card-img]: https://goreportcard.com/badge/github.com/anuvu/dig
[report-card]: https://goreportcard.com/report/github.com/anuvu/dig
