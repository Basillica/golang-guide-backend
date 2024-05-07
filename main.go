package main

import (
	"github.com/Basillica/golang-guide/handler/api"
)

const Version = "v1"

func main() {

	api.New(Version)
}
