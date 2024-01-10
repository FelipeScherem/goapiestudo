package main

import (
	"fmt"
	"github.com/FelipeScherem/goapiestudo.git/config"
	"github.com/FelipeScherem/goapiestudo.git/router"
)

func main() {
	// Initialize config
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Initialize o router
	router.Initialize()
}
