package main

import (
	"github.com/adamknbrv/forum/internal/app"
	"github.com/adamknbrv/forum/pkg/log"
)

func main() {
	log.RunLogger()
	app.Run()
}
