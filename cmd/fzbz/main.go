package main

import (
	"flag"
	"fmt"
	"github.com/a-skua/fzbz"
	"strings"
)

var (
	templ = flag.String("t", "{}", "template")
	num   = flag.Int("n", 15, "max number")
)

func main() {
	flag.Parse()

	for i := 0; i < *num; i++ {
		fmt.Println(strings.ReplaceAll(*templ, "{}", fzbz.FizzBuzz(i+1)))
	}
}
