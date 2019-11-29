![Imgur](https://imgur.com/To5zr4X.jpg)

![](https://github.com/gocomu/comu/workflows/CI/badge.svg?branch=master) [![codecov](https://codecov.io/gh/gocomu/comu/branch/master/graph/badge.svg)](https://codecov.io/gh/gocomu/comu) [<img src="https://img.shields.io/badge/slack-gocomu/gophers-blue.svg?logo=slack">](https://app.slack.com/client/T029RQSE6/CQE31A4E5) [<img src="https://img.shields.io/badge/slack-get/invite-green.svg?logo=slack">](https://invite.slack.golangbridge.org/)

# comu
 
comu (computer music) is an open source music library for creative coding in Go.

# Getting Started

Bellow are instruction on how to install `comu` and run a few [examples](). 


Full documentaion lives in this repo's [wiki](https://github.com/gocomu/comu/wiki).

## CLI

To get started faster you can use [`gocomu`](https://github.com/gocomu/cli), a command line interface helper for `comu`.

To install `gocomu` run 
```
go get github.com/gocomu/cli/cmd/gocomu
```

For information on how to use it you can 

1. read `gocomu`'s [documentation](https://github.com/gocomu/cli/blob/master/README.md) online
2. or run `gocomu -help`

# Prerequisites

## Go

If you don't have `go` installed already, here is the [official documenation](https://golang.org/doc/install).

## PortAudio

At the moment the moment the only way to make real time sound is using port-audio. 

The download link is [here](http://www.portaudio.com/download.html) and instructions for each platform [here](http://portaudio.com/docs/v19-doxydocs/tutorial_start.html).

# Installation

to clone the library locally along the examples

``` 
go get github.com/gocomu/comu
```

# Use

## Examples

### comuGen

`cd $GOPATH/src/github.com/gocomu/comu/examples/comuGen`

`go run main.go`


# roadmap to v0.1.0
 - [x] TempoClock
 - [ ] Time functions
 - [ ] Patterns
 - [ ] Timeline
 - [ ] Pan (support for 2..n channels)
 - [ ] Utilities implementation
    - [ ] embedder
    - [ ] renderer
 - [ ] IO wrappers
    - [x] OpenSoundControl
    - [ ] Audio
       - [ ] In
       - [x] Out
    - [ ] Midi
- [x] README
- [ ] wiki
- [ ] Pan implementation (support for 2-4-8-n channels)
- [ ] CLI
- [ ] wrap Oto (as alternative to port-audio)