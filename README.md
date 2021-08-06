<h1 align="center">CliCtx</h1>
<p align="center">
        <a href="https://github.com/fzerorubigd/clictx/releases"><img src="https://img.shields.io/github/v/tag/fzerorubigd/clictx.svg?color=brightgreen&label=version&sort=semver"></a>
        <a href="https://travis-ci.org/github/fzerorubigd/clictx"><img src="https://travis-ci.org/fzerorubigd/clictx.svg?branch=master"></a>
        <a href="https://goreportcard.com/report/github.com/fzerorubigd/clictx"><img src="https://goreportcard.com/badge/github.com/fzerorubigd/clictx"></a>
        <a href="https://codecov.io/gh/fzerorubigd/clictx"><img src="https://codecov.io/gh/fzerorubigd/clictx/branch/master/graph/badge.svg"/></a>
        <a href="https://godoc.org/github.com/fzerorubigd/clictx"><img src="https://img.shields.io/badge/godoc-reference-blue.svg?color=blue"></a>
        <a href="https://github.com/fzerorubigd/clictx/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-MIT-green"></a>
</p>

---
As of Go1.16 there is a <a href="https://pkg.go.dev/os/signal#NotifyContext">function in the standard library</a> for doing this, I will no longer maintain this library. 
This package helps you to create a context that cancels on specific signals. 

Usage: 

```go
package main 

import "github.com/fzerorubigd/clictx"

func main() {
    ctx := clictx.Context() /* Pass a specific signal to watch only for that 
    signal, otherwise, all signals */

    // Your application entry point, you need to use this context as base for all
    // context in your application
    // app(ctx)

    <- ctx.Done() 
}

```
