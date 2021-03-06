![Imgur](https://imgur.com/To5zr4X.jpg)

![](https://github.com/gocomu/comu/workflows/build/badge.svg?branch=master) [![codecov](https://codecov.io/gh/gocomu/comu/branch/master/graph/badge.svg)](https://codecov.io/gh/gocomu/comu) [<img src="https://img.shields.io/badge/slack-gocomu/gophers-blue.svg?logo=slack">](https://app.slack.com/client/T029RQSE6/CQE31A4E5) [<img src="https://img.shields.io/badge/slack-get/invite-green.svg?logo=slack">](https://invite.slack.golangbridge.org/)

# comu
 
comu _(computer music)_ is an open source music library for creative coding in Go.

### Features

* Real-time audio
* Multi OS support (Linux/Mac/Win)
* Command Line helper
* Build stand-alone applications (CLI or GUI)
* Benefit from existing Go libraries and tools

# Getting Started

Bellow are instruction on how to install `comu` and run a few [examples](https://github.com/gocomu/comu/tree/master/examples). 


Full documentaion exists in this repo's [wiki](https://github.com/gocomu/comu/wiki).

# Prerequisites

### Go

If you don't have Go installed already, here is the [official documenation](https://golang.org/doc/install).

### PortAudio

At the moment the only supported way to make real time sound is PortAudio. 

The download link is [here](http://www.portaudio.com/download.html) and instructions for each platform [here](http://portaudio.com/docs/v19-doxydocs/tutorial_start.html).

#### Ubuntu Package

`apt install portaudio19-dev`

#### Homebrew 

`brew install portaudio`

## Command Line helper

To get started faster you can use [`gocomu`](https://github.com/gocomu/cli), a command line interface designed to make working with `comu` a simpler experience. Especially recommended to newcomers.

Install `gocomu` by running 
```
go get github.com/gocomu/cli/cmd/gocomu
```

For information on how to use it you can 

1. read `gocomu`'s [documentation](https://github.com/gocomu/cli/blob/master/README.md) online
2. or run `gocomu -help`


# Installation

to clone the library locally along the examples run:

``` 
go get github.com/gocomu/comu
```

# Use

### Examples

### comuGen

`cd $GOPATH/src/github.com/gocomu/comu/examples/comuGen`

`go run main.go`

### comuOSC


# Roadmap to v0.1.0
 - [x] TempoClock
 - [ ] Time functions
 - [ ] Patterns
 - [ ] Timeline
 - [ ] Pan (support for 2..n channels)
 - [ ] Utilities
    - [ ] embedder
    - [ ] renderer
    - [ ] recorder
 - [x] IO wrappers
    - [x] OpenSoundControl
    - [x] Audio
       - [ ] In
       - [x] Out
    - [ ] Midi
- [x] README
- [ ] wiki
- [x] CLI
- [ ] replace PortAudio with customised Oto
