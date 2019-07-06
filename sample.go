// +build ignore

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
