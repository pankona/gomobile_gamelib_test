[![Circle CI](https://circleci.com/gh/pankona/gomo-simra/tree/master.svg?style=svg)](https://circleci.com/gh/pankona/gomo-simra/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/pankona/gomo-simra)](https://goreportcard.com/report/github.com/pankona/gomo-simra)
[![GoDoc](https://godoc.org/github.com/pankona/gomo-simra?status.svg)](https://godoc.org/github.com/pankona/gomo-simra)  
[ZenHub](https://app.zenhub.com/workspace/o/pankona/gomo-simra/)
<a href="https://app.zenhub.com/workspace/o/pankona/gomo-simra"><img src="https://raw.githubusercontent.com/ZenHubIO/support/master/zenhub-badge.png"></a>


# gomo-simra

GOMObile SIMple wRApper

# What is this

* simple wrapper for gomobile.
* purpose is ...
  * developers can focus only for scene creation.
    * wraps calculation of x, y coodinates using size.Event structure. 
    * wraps usage of f32.affine.
    * calculates scale to fit to any device's screen automatically.
    * provides easy scene transition
* everything is under construction.
* see `example` directory to know how to use.

# Build

* For Release build, `go build -tags=release` or `gomobile build -tags=release`
  * This efforts that logging for debug is disabled.

# License

MIT
