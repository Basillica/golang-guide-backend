package main

import (
	"github.com/Basillica/golang-guide/handler/api"
)

const Version = "v2"

func main() {

	api.New(Version)
}
