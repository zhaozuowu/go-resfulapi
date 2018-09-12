package main

import "os"

var user = os.Getenv("USER")

func init() {
	if user == "" {
		panic("no value for $USER")
	}
}

func main()  {

	if user == "" {
		panic("no value for $USER")
	}
}