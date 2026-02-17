package infrastructure

import "fmt"

// Simple event bus example (replace with Redis Streams or NATS in production)
type SimpleEventBus struct{}

func (b *SimpleEventBus) Publish(event interface{}) error {
	fmt.Printf("Event publishe: %#v\n", event)
	return nil
}
