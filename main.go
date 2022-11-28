package main

import (
	"fmt"

	"genpwd/cryptosource"
	"genpwd/generator"
)

func main() {
	src := cryptosource.New()
	g := generator.New(src)
	fmt.Println(g.GenPassword(3, 5))
}
