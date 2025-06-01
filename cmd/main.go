package main

import (
	"github.com/Defenestrationq/forum-api/internal/app"
	"github.com/Defenestrationq/forum-api/pkg/log"
)

func main() {
	log.RunLogger()
	app.Run()
}
