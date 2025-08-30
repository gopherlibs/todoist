This repo is being setup. Nothing is ready yet, not even for contributions.

# GopherLibs => Todoist [![Go Reference](https://pkg.go.dev/badge/github.com/gopherlibs/todoist.svg)](https://pkg.go.dev/github.com/gopherlibs/todoist) [![Go Report Card](https://goreportcard.com/badge/github.com/gopherlibs/todoist)](https://goreportcard.com/report/github.com/gopherlibs/todoist) [![Software License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/gopherlibs/todoist/trunk/LICENSE)

*This project is really early and figuring out its identity. Don't expect a stable (Go) API at the moment.*

`Todoist` is a Go (Golang) wrapper and SDK for the official Todoist API.
A basic CLI is also included.

The CLI is and will remain fairly basic.
It is meant to test this library, not really for end users.


## Table of Contents

- [Goals](#goals)
- [Requirements](#requirements)
- [Usage](#usage)
- [Development](#development)
- [Credits](#credits)


## Goals

- The first and only goal right now is to get this module usable to return the current MLB standings. This will then be used in the [wtfutil/wtf](https://github.com/wtfutil/wtf) project as a widget/module.


## Requirements

The minimum Go version supported is v1.24.x.


## Usage

`Big League Stats` is a Go module so the best way to use it is to add an import in the file you want to use it in and then run `go mod tidy` to get it downloaded.

```go
import(
	"github.com/gopherlibs/big-league-stats/mlb"
)
```

Alternatively, you can run `go get github.com/gopherlibs/big-league-stats/mlb` in your project directory.


## Usage

```go
package main

import (
	"fmt"

	"github.com/gopherlibs/big-league-stats/mlb"
)

func main() {

	img, err := gpic.NewImage("Ricardo@Feliciano.Tech")
	if err != nil {
		fmt.Println(err)
		return
	}

	imgURL, err := img.URL()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(imgURL.String())
}
```

## Development

This library is written and tested with Go v1.24+ in mind.
`go fmt` is your friend.
Please feel free to open Issues and PRs are you see fit.
Any PR that requires a good amount of work or is a significant change, it would be best to open an Issue to discuss the change first.


## Credits

This module was written by Ricardo N Feliciano (FelicianoTech).
This repository is licensed under the MIT license.
This repo's license can be found [here](./LICENSE).
