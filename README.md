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
