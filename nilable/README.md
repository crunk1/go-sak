# nilable
`import "github.com/crunk1/go-sak/nilable"`

nilable provides type-wrappers around all Go primitives. These types represent
the primitive data types, but also allow for nil values.
 
Example:
```go
package main

import (
    "fmt"

    "github.com/crunk1/go-sak/nilable"
)

func main() {
	var nilS nilable.String

	// Check against nil.
	fmt.Println(nilS == nil)                         // true
	fmt.Println(nilable.NewString("foo") == nil)     // false

	// Get the underlying value, but default to zero value.
	fmt.Println(nilS.VOrZero())                      // ""
	fmt.Println(nilable.NewString("foo").VOrZero())  // "foo"

	// Get the underlying value.
	// nilS.V()                                      // panics because nilS is nil.
	fmt.Println(nilable.NewString("foo").V())        // "foo"
}
```