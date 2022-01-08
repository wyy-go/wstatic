# wstatic

![GitHub Repo stars](https://img.shields.io/github/stars/wyy-go/wstatic?style=social)
![GitHub](https://img.shields.io/github/license/wyy-go/wstatic)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/wyy-go/wstatic)
![GitHub CI Status](https://img.shields.io/github/workflow/status/wyy-go/wstatic/ci?label=CI)
[![Go Report Card](https://goreportcard.com/badge/github.com/wyy-go/wstatic)](https://goreportcard.com/report/github.com/wyy-go/wstatic)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/wyy-go/wstatic?tab=doc)
[![codecov](https://codecov.io/gh/wyy-go/wstatic/branch/main/graph/badge.svg)](https://codecov.io/gh/wyy-go/wstatic)



Static for gin middleware

## Usage

### Start using it

Download and install it:

```sh
go get github.com/wyy-go/wstatic
```

Import it in your code:

```go
import "github.com/wyy-go/wstatic"
```

### Canonical example

See the [_example](_example)

```go
package main

import (
  "github.com/wyy-go/wstatic"
  "github.com/gin-gonic/gin"
  "log"
)

func main() {
  r := gin.Default()
  
  r.Use(wstatic.New(wstatic.WithUrlPrefix("/"),
	  wstatic.WithRoot("./form"),
	  wstatic.WithIndexes(false)))
  
  r.GET("/ping", func(c *gin.Context) {
    c.String(200, "test")
  })
  
  // Listen and Server in 0.0.0.0:8080
  if err := r.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```
