package main

import (
	"fmt"

	"github.com/bianavic/utils_secret/pkg/events"
)

func main() {
	eventDispatcher := events.NewEventDispatcher()
	fmt.Println(eventDispatcher)
}
