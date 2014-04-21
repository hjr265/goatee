# Goatee

A tool to generate `.go` files from `.got` templates. Can be used to embed abitrary blobs of data from other files in Go source files.

## Installation

Install Goatee using the `go get` command:

     $ go get github.com/hjr265/goatee

The only dependency is the Go distribution itself.

## Usage

Create a `.got` template, and, may be, name it `hello.got`

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println({{string "greeting.txt"}})
}
```

Create a file named `greeting.txt` in the same directory

    Hello, world!

Execute Goatee

    $ goatee hello.got

This will generate the following as `hello.go`

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
}
```

Right now, there is the template function `string`, which reads a file and echos the content as a string literal; and then there is `bytes`, which reads the file and echos the content as an array of bytes.

## Documentation

- [Reference](http://godoc.org/github.com/hjr265/goatee)

## Contributing

Contributions are welcome.

## License

Goatee is available under the [BSD (3-Clause) License](http://opensource.org/licenses/BSD-3-Clause).

