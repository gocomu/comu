![Imgur](https://imgur.com/To5zr4X.jpg)

![](https://github.com/gocomu/comu/workflows/CI/badge.svg?branch=master) [![codecov](https://codecov.io/gh/gocomu/comu/branch/master/graph/badge.svg)](https://codecov.io/gh/gocomu/comu) [<img src="https://img.shields.io/badge/slack-gocomu/gophers-blue.svg?logo=slack">](https://app.slack.com/client/T029RQSE6/CQE31A4E5) [<img src="https://img.shields.io/badge/slack-get/invite-green.svg?logo=slack">](https://invite.slack.golangbridge.org/)

# comu
 
comu (computer music) is an open source library, aiming to assists electronic music creation using Go.

comu uses [go-audio](https://github.com/go-audio).

At the moment we are at the early stages with few things working.

The roadmap to v0.1.0 includes:
 - [x] TempoClock
 - [ ] Time functions
 - [ ] Pattern system design
 - [ ] Utilities implementation
    - [ ] embedder
    - [ ] renderer
    - [ ] recorder
 - [ ] IO wrappers
    - [x] OpenSoundControl
    - [X] Audio
    - [ ] Midi
- [ ] Pan implementation (support for 2-4-8-n channels)
- [ ] CLI helper https://github.com/gocomu/cli
- [ ] wrapper around Oto (as alternative to port-audio)
- [ ] RTaudio?

# installation

to install the library 

``` 
go get github.com/gocomu/comu
```

## port-audio

# usage

## examples

`go run ./examples/generator`
