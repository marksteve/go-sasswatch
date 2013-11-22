go-sasswatch
============

Adds ability to compile SASS files on filesystem change within your Go applications

Example
-------

```go
package main

import (
  "github.com/marksteve/go-sasswatch"
  "github.com/marksteve/go-sasswatch/gosass"
)

func main() {
  sw := sasswatch.SassWatcher("/path/to/watch", gosass.Options{
    OutputStyle:  gosass.NESTED_STYLE,
    IncludePaths: []string{"foundation/scss"},
  })
  // Your app stuff...
  sw.Close()
}
```

Notes
-----

I'm using gosass (https://github.com/moovweb/gosass) to interface with libsass.
I included `gosass.go` directly instead of importing it from its repo because stupid
me can't figure out how some of the C stuff works :P
