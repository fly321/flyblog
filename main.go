package main

import (
	"flyblog/src/config"
	"flyblog/src/core"
	"fmt"
)

func main() {
	c := config.DbConfig
	fmt.Printf("%+v\n", c.ListDb())
	db := core.Db
	fmt.Printf("%+v\n", db)
}
