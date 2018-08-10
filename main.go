package main

import (
	"log"

	"github.com/docker/docker/pkg/broadcaster"
	events "github.com/docker/go-events"
)

func main() {
	log.Println(events.Channel{})
	log.Println(broadcaster.Unbuffered{})
}
