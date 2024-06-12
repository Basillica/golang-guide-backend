package main

import (
	"fmt"

	"github.com/Basillica/golang-guide/handler/api"
)

var (
	Version = "v1"
	Date    = ""
)

func main() {
	fmt.Println("app was built on ", Date)
	api.New(Version)
}
