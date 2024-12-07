# maple

Map operations that are missing from the GoLang core.

## Installation

```sh
go get github.com/briandamaged/maple
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/briandamaged/maple"
)

func main() {
	m := map[string]int{
		"pizza":  5,
		"salad":  5,
		"soup":   4,
		"cheese": 6,
	}

	fmt.Println("m                : ", m)

	// Extract the keys from `m`
	fmt.Println("maple.Keys(m)    : ", maple.Keys(m))

	// Extract the values from `m`
	fmt.Println("maple.Values(m)  : ", maple.Values(m))

	// Create a new map that inverts the keys and values of `m`
	fmt.Println("maple.Inverse(m) : ", maple.Invert(m))
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
