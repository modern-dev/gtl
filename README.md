
# Go Template Library

<img src="http://modern-dev.com/images/gtl/gtl_logo.png" width="180" align="right" />

A set of extensions and helpers to Go's standard library.

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/modern-dev/gtl/Build%20&%20Test%20workflow?style=flat) ![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/modern-dev/gtl?style=flat)  ![GitHub](https://img.shields.io/github/license/modern-dev/gtl?style=flat)  [![Gitter chat](https://badges.gitter.im/gitterHQ/gitter.png)](https://gitter.im/modern-dev/gtl)

## :pushpin: Overview


GTL is a set of extension to the Go's standard library providing some of the most widely used data structures and algorithms written in [Go 2](https://go.googlesource.com/go/+/refs/heads/dev.go2go/README.go2go.md).

The package hierarchy is separated into several standalone components offering different set of tools, just like in [C++](https://en.wikipedia.org/wiki/Standard_Template_Library):
- The `containers` submodule offers the objects that store data. It includes *ordered collections* ([`vector`](https://en.wikipedia.org/wiki/Vector_(STL)), [`singly linked list`](https://en.wikipedia.org/wiki/Linked_list), [`doubly linked list`](https://en.wikipedia.org/wiki/Doubly_linked_list)), *container adaptors* ([`queue`](https://en.wikipedia.org/wiki/Queue_(data_structure)), [`priority queue`](https://en.wikipedia.org/wiki/Priority_queue), `heap`, [`stack`](https://en.wikipedia.org/wiki/Stack_(data_structure))), *associative containers* ([`set`](https://en.wikipedia.org/wiki/Set_(computer_science))), *simple containers* (`pair`) and others (`bitset`);
- The `algo` submodule provides components that perform algorithmic operations on containers and other sequences;
- The `funcs` submodule is reach in many helper functions.

The main motivation of this project is to provide Go's ecosystem with an ultimate toolbox for comfortable development process. That is why we decided to use Go 2 alongside with Go's new [generics feature](https://blog.golang.org/generics-next-step) to achieve our main goal. Generics enabled us to write elegant, concise and future proof code that can be used as a foundation for many applications of your own.

<img src="http://modern-dev.com/images/gtl/GOPHER_LAPTOP.png" width="180" align="right" />

## :minidisc: Build and Test

You must first build `go2go` toolchain locally in order to run code from `go2` sources. It is available in a branch of the master Go repo. Follow the [instructions on installing Go from source](https://golang.org/doc/install/source). Where those instructions direct you to check out the latest release tag, instead run `git checkout dev.go2go`. Then build the Go toolchain as directed.

You can then use the new `go2go` tool to build (`go tool go2go run .`) and test (`go tool go2go test .`) the code.

GTL may also be used alongside with our [docker-go2go](https://github.com/modern-dev/docker-go2go) image that provides precompiled `go2go` source out of the box. You may then use `go2` command inside the container.

<img src="http://modern-dev.com/images/gtl/GOPHER_SHARE.png" width="220" align="right" />

## :bulb: Contributing to GTL

For a complete guide to contributing to GTL, see the [Contribution Guide](https://github.com/modern-dev/gtl/blob/master/CONTRIBUTING.md).

## :green_book: License

Copyright 2020. The GTL Authors. All rights reserved.

Licensed and the [MIT license](https://raw.githubusercontent.com/modern-dev/gtl/master/LICENSE).

[Gopher artwork by Ashley McNamara](https://github.com/ashleymcnamara/gophers) licensed under the [Creative Commons Attribution-NonCommercial-ShareAlike 4.0 License](https://creativecommons.org/licenses/by-nc-sa/4.0/)

---

> [modern-dev.com](http://modern-dev.com) &nbsp;&middot;&nbsp;
> GitHub [@modern-dev](https://github.com/modern-dev)