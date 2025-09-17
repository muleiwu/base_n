# base-n

A lightweight Go library for encoding and decoding values in arbitrary base-n numbering systems.

## Features

- Convert numbers to and from custom base-n representations
- Support for negative numbers
- Built-in presets for common bases:
  - Base62 (alphanumeric: 0-9, a-z, A-Z)
  - Base3 (0, 1, 2)
  - Special Base3 (0, o, O)
- Create custom alphabets for specialized encoding schemes

## Installation

```bash
go get github.com/muleiwu/base-n
```

## Usage

### Basic Usage

```go
package main

import (
    "fmt"
    BaseN "github.com/muleiwu/base-n"
)

func main() {
    // Create a Base62 encoder/decoder
    base62 := BaseN.NewBase62()
    
    // Encode an integer to Base62 string
    encoded := base62.Encode(12345)
    fmt.Println("Encoded:", encoded)
    
    // Decode a Base62 string back to integer
    decoded, err := base62.Decode(encoded)
    if err != nil {
        panic(err)
    }
    fmt.Println("Decoded:", decoded)
}
```

### Custom Alphabet

```go
package main

import (
    "fmt"
    BaseN "github.com/muleiwu/base-n"
)

func main() {
    // Create a custom Base with your own character set
    customBase := BaseN.NewBaseN([]byte("01"))  // Binary
    
    // Encode and decode using the custom base
    encoded := customBase.Encode(10)
    fmt.Println("Encoded in binary:", encoded)  // Outputs: 1010
    
    decoded, _ := customBase.Decode(encoded)
    fmt.Println("Decoded:", decoded)  // Outputs: 10
}
```

### Predefined Bases

```go
// Base3 (0, 1, 2)
base3 := BaseN.NewBase3Number()

// Special Base3 (0, o, O)
base3Special := BaseN.NewBase30oO()

// Base62 (0-9, a-z, A-Z)
base62 := BaseN.NewBase62()
```

## API Reference

### Constructor Functions

- `NewBaseN(chars []byte) *BaseN`: Create a new Base-N encoder/decoder with custom characters
- `NewBase3Number() *BaseN`: Create a Base3 encoder/decoder (0, 1, 2)
- `NewBase30oO() *BaseN`: Create a special Base3 encoder/decoder (0, o, O)
- `NewBase62() *BaseN`: Create a Base62 encoder/decoder (0-9, a-z, A-Z)

### Methods

- `Encode(n int64) string`: Convert an integer to a base-n string
- `Decode(s string) (int64, error)`: Convert a base-n string back to an integer

## License

MIT
