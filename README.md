# Tetra

A CSV pre-processing library for Tesseract. Provide a string csv and a list of row/cell/column transformations, and get the transformed string csv back.

## Usage

```go

package main

import (
	"fmt"

	"github.com/arizard/tetra"
)

func main() {
    // A config is required. Takes same params as csv.Reader
    config := tetra.Config{
        Comma:           ',',
        FieldsPerRecord: -1,
    }

    // Add transforms in the same order as you want them to be applied.
    config.AddTransform(
        "slice_rows",
        map[string]interface{}{
            "start": 1,
            "end":   -1,
        },
    )
    config.AddTransform(
        "reverse_rows",
        map[string]interface{}{},
    )

    // Input data is a string
    csvString1 := "1,2,3,\na,b,c,\n4,5,6,\n"
    csvString2 := "x,y,z,\n0,0,0,\nm,n,o,\n"

    // Output data is a string
    output1 := tetra.TransformCSV(config, csvString1)
    output2 := tetra.TransformCSV(config, csvString2)

    fmt.Println(output1)
    fmt.Println(output2)
}


```

## Transformations

### Slice Rows

Get a slice of rows from the csv. `start` is the first row of the slice. `end` is the last row of the slice. `start` and `end` can be negative values, referring to the nth row from the end.

Example:

```go
config.AddTransform(
    "slice_rows",
    map[string]interface{}{
        "start": 1, // 2nd row
        "end":   -2, // 2nd last row
    },
)
```

### Reverse Rows

Reverse all rows of the csv. Takes no keyword arguments.

Example:

```go
config.AddTransform(
    "reverse_rows",
    map[string]interface{}{},
)
```