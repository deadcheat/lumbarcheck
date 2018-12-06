package main

import (
	"fmt"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	dsn = kingpin.Arg("dsn", "Data Source Name you'll try to connect to").Required().String()
)

func main() {
	kingpin.Parse()

	fmt.Println("dsn: ", *dsn)
}
