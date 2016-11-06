package main

import (
	"github.com/useada/gopher"
)

func main() {
	go gopher.RssRefresh()
	gopher.StartServer()
}
