package main

import (
	"log"

	"github.com/docker/docker/container"
	events "github.com/docker/go-events"
)

func main() {
	log.Println(events.Channel{})
	log.Println(container.Health{})
}
