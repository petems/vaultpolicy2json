# remove-singleton-arrays

[![Build Status](https://travis-ci.org/petems/remove-singleton-arrays.svg?branch=master)](https://travis-ci.org/petems/remove-singleton-arrays)

This is a golang package to remove arrays that only have one element from a JSON object.

Arrays with more than one element are ignored.

## Usage

To remove all singleton arrays:

```go
import (
	. "github.com/petems/remove-singleton-arrays"
)

RemoveSingletonArrays(`{"a":1,"b":["2"]}`)  // returns `{"a":1,"b":"2"}`
RemoveSingletonArrays(`{"a":1,"b":["2, 3"]}`) // returns `{"a":1,"b":["2","3"]}`
```

To remove singleton arrays with selected exceptions:

```go 
import (
	"github.com/petems/remove-singleton-arrays"
)

removesingletonarrays.WithIgnores(`{"a":1,"b":["2"]}`, []string{"b"})  // returns `{"a":1,"b":["2"]}`
removesingletonarrays.WithIgnores(`{"a":1,"b":["2"]}`, []string{"a"})  // returns `{"a":1,"b":"2"}`
```