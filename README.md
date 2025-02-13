# greetings

A simple Go module for generating greeting messages.

## Features
- Generate a greeting message for a given name.
- Lightweight and easy to use.

## Installation

To install the module, run the following command:

```sh
go get github.com/agisnarevaldo/greetings
```

## Usage

Here's an example of how to use the `greetings` module:

```go
package main

import (
    "fmt"
    "github.com/username/greetings"
)

func main() {
    fmt.Println(greetings.Hello("Go"))
}
```

## Versioning

This project follows [Semantic Versioning](https://semver.org/).  
To upgrade to a specific version, use:

```sh
go get github.com/agisnarevaldo/greetings@v1.0.0
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.