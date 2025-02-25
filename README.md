# UTM Parameter Serializer

A library for serializing parameters into a UTM formatted string with generics.
Currently standard types used by the API are supported (`uint`, `int`, `string` `[]string`)

---

## Features

- **Generic Option Handling**: Supports multiple types (`uint`, `int`, `string`, `[]string`) using Go generics.
- **Customizable Output**: Generates formatted strings with parameter names and values.
- **Error Handling**: Gracefully handles unsupported types and provides meaningful error messages.
- **Extensible**: Easily add support for new types by implementing the `Option` interface.

---

## Installation

Quick and simple, import as a dependency wherever you need to use it

```go
import "t25.tokyo/utm"
```

---

## Usage

### Basic Example

```go
package main

import (
	"log"
	"t25.tokyo/utm"
)

func main() {
	// Create options
	options := []utm.Option{
		utm.Uint("myUintParam", 42, 0),
		utm.String("myStringParam", "example", ""),
		utm.Int("myIntParam", 10, 0),
		utm.StringArray{Param: "myArrayParam", Values: []string{"one", "two", "three"}},
	}

	// Process options and get the result
	result := ProcessOptions(options)
	log.Println(result)
}
```

### Output

```
uintParam 42 stringParam example intParam 10 arrayParam one arrayParam two arrayParam three
```

---

### Supported Types

| Type       | Example Usage                          |
|------------|----------------------------------------|
| `uint`     | `NewUintOption("param", 42, 0)`        |
| `int`      | `NewIntOption("param", 10, 0)`         |
| `string`   | `NewStringOption("param", "value", "")`|
| `[]string` | `StringArrayOption{Param: "param", Values: []string{"one", "two"}}` |

---

### Error Handling

If an unsupported type is provided, the library will return an error:

```go
opt, err := routeType("floatParam", 3.14, 0.0)
if err != nil {
    log.Printf("Error: %v", err)
}
```

