package main

import (
	"fmt"
	"flag"
)

var name = flag.String("name", "everyone", "The greeting object ")

func init() {

}

func main() {

	flag.Parse()

	fmt.Printf("this name is :%v", *name);
}
