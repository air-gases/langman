# Langman

[![GoDoc](https://godoc.org/github.com/air-gases/langman?status.svg)](https://godoc.org/github.com/air-gases/langman)

A useful gas that used to manage the Accept-Language header for the web
applications built using [Air](https://github.com/aofei/air).

## Installation

Open your terminal and execute

```bash
$ go get github.com/air-gases/langman
```

done.

> The only requirement is the [Go](https://golang.org), at least v1.12.

## Usage

Create a file named `main.go`

```go
package main

import (
	"github.com/air-gases/langman"
	"github.com/aofei/air"
)

func main() {
	a := air.Default
	a.DebugMode = true
	a.GET("/", func(req *air.Request, res *air.Response) error {
		return res.WriteString(req.Header.Get("Accept-Language"))
	}, langman.Gas(langman.GasConfig{}))
	a.Serve()
}
```

and run it

```bash
$ go run main.go
```

then visit `http://localhost:8080/?accept-language=en-US`.

## Community

If you want to discuss Langman, or ask questions about it, simply post questions
or ideas [here](https://github.com/air-gases/langman/issues).

## Contributing

If you want to help build Langman, simply follow
[this](https://github.com/air-gases/langman/wiki/Contributing) to send pull
requests [here](https://github.com/air-gases/langman/pulls).

## License

This project is licensed under the Unlicense.

License can be found [here](LICENSE).
