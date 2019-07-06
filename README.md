# Sysloc

[![GoDoc](https://godoc.org/github.com/RadhiFadlillah/sysloc?status.png)](https://godoc.org/github.com/RadhiFadlillah/sysloc)

Sysloc is a Go package for getting OS locale. It's inspired by [os-locale](https://github.com/sindresorhus/os-locale) by [Sindre Sorhus](https://github.com/sindresorhus), except here for Mac OS we simply use `$LANG` environment variable.

## Usage Examples

Says we want to get the locale that used by system :

```go
package main

import (
	"fmt"

	"github.com/RadhiFadlillah/sysloc"
	"golang.org/x/text/language/display"
)

func main() {
	locale, err := sysloc.GetLocale()
	if err != nil {
		panic(err)
	}

	fmt.Printf("System locale is %s (%s)\n",
		display.Self.Name(locale),
		locale,
	)
}
```

Which will give us following results :

```
System locale is American English (en-US)
```

## License

Sysloc is distributed using [MIT](http://choosealicense.com/licenses/mit/) license.
