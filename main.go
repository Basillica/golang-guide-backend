package main

import (
	"fmt"

	"github.com/Basillica/golang-guide/api"
)

var (
	Version = "v1"
	Date    = ""
)

func main() {
	fmt.Println("app was built on ", Date)
	a := api.New(Version)
	a.Start(a.Engine)
}
