package main

import (
	"flag"
	"fmt"
)

func main() {
	greeting := flag.String("greeting", "hello", "word to greet with")
	flag.Parse()
	name := "world"
	if flag.NArg() > 0 {
		name = flag.Arg(0)
	}
	fmt.Printf("%s, %s\n", *greeting, name)
}
