package main

import (
	"fmt"
	"github.com/bianavic/fullcycle_eventos/pkg/events"
)

// pkg/events/event_dispatcher.go:1
func main() {
	eventDispatcher := events.NewEventDispatcher()
	fmt.Println(eventDispatcher)
}
