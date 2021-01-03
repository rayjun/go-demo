package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var c = flag.String("c", " ", "conetxt")
var a = flag.String("a", "sha256", "algroithm")

func main() {
	flag.Parse()

	if *a == "sha256" {
		fmt.Printf("%x", sha256.Sum256([]byte(*c)))
	} else if *a == "sha384" {
		fmt.Printf("%x", sha512.Sum384([]byte(*c)))
	} else if *a == "sha512" {
		fmt.Printf("%x", sha512.Sum512([]byte(*c)))
	} else {
		fmt.Printf("Error algo %s\n", *c)
		os.Exit(1)
	}
}
