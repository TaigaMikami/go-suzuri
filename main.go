package main

import (
	"github.com/TaigaMikami/go-suzuri/suzuri"
	"github.com/k0kubun/pp"
	"github.com/labstack/gommon/log"
)

func main() {
	c := suzuri.New("apiKey")

	items, err := c.GetItems()
	if err != nil {
		log.Error(err)
	}

	pp.Print(items)
}
