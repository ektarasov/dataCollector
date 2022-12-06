package main

import (
	"diplom/pkg"
	"flag"
)

func main() {
	flag.Parse()
	pkg.ListenAndServeHTTP()
}
