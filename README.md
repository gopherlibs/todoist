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

- [x] Implement the minimum functionality needed to support the Todoist module in [wtfutil/wtf](https://github.com/wtfutil/wtf).
- [ ] All task endpoints supported.
- [ ] Sync endpoint supported.
- [ ] Other endpoints supported.


## Requirements

- The minimum Go version supported is v1.25.
- A Todoist account will be needed.


## Usage

Gopherlibs Todoist is a Go module so the best way to use it is to add an import in the file you want to use it in and then run `go mod tidy` to get it downloaded.
For example:

```go
import(
	"github.com/gopherlibs/todoist/api"
)
```

Alternatively, you can run `go get github.com/gopherlibs/todoist/api` in your project directory.


## Usage

In short, you create a new instance of the `client`, authenticate with it, and then call methods that return the objects you'd like.
Here's an example of getting a slice of tasks.

```go
package main

import (
	"fmt"

	"github.com/gopherlibs/todoist/api"
)

func main() {

	c := api.New(os.Getenv("TODOIST_TOKEN"))

	tasks, err := c.Tasks(args[0])
	if err != nil {
		fmt.Printf("Running 'Tasks' failed. Err: %s\n", err.Error())
		return err
	}

	for _, t := range tasks.Results {
		fmt.Printf("- %s - %s\n", t.Content, t.ID)
	}
}
```

## Development

This library is written and tested with Go v1.25+ in mind.
`go fmt` is your friend.
Please feel free to open Issues and PRs are you see fit.
Any PR that requires a good amount of work or is a significant change, it would be best to open an Issue to discuss the change first.


## Credits

This module was written by Ricardo N Feliciano (FelicianoTech).
This repository is licensed under the MIT license.
This repo's license can be found [here](./LICENSE).
