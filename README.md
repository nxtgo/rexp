# rexp

chainable regexp builder inspired by [magic-regexp](https://regexp.dev/).

## usage

```go
package main

import (
	"fmt"
	"regexp"

	"github.com/nxtgo/rexp"
)

func main() {
	reg := rexp.Create(
		rexp.Digit().OneOrMore().As("num"),
	)

	text := "version 123 released"
	m := reg.FindStringSubmatch(text)
	if m != nil {
		names := reg.SubexpNames()
		for i, name := range names {
			if i == 0 || name == "" {
				continue
			}
			fmt.Printf("%s = %s\n", name, m[i])
		}
	}
}
```

# license

CC0 1.0 (public domain) + ip waiver.
