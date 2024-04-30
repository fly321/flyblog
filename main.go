package main

import (
	"flyblog/src/config"
	"fmt"
)

func main() {
	c := config.DbConfig
	fmt.Printf("%+v\n", c.ListDb())
}
