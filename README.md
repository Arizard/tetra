Tetra
========

Tetra is a csv processing library for the Tesseract ecosystem.

Usage example:

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
    // In this case, we exclude the first row.
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
Features
--------

- Define sequences of CSV transformations.

Installation
------------

Install Tetra by running:

    go get github.com/arizard/tetra

Contribute
----------

- Issue Tracker: https://github.com/Arizard/tetra/issues
- Source Code: https://github.com/Arizard/tetra

Support
-------

If you are having issues, please let us know using the issue tracker.

License
-------

The project is licensed under the BSD license.