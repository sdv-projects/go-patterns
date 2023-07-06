package domain_events

import (
	"context"
	"fmt"
	"reflect"
	"sync"
)

type DomainEventChannel struct {
	handlers map[reflect.Type]map[DomainEventHandler]interface{}
	mu       sync.Mutex
}

func (c *DomainEventChannel) Publish(ctx context.Context, events ...DomainEvent) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	riseMap := make(map[DomainEventHandler][]DomainEvent)

	for _, e := range events {
		if handlers, ok := c.handlers[e.GetEventType()]; ok {
			for h := range handlers {
				if _, ok := riseMap[h]; ok {
					riseMap[h] = make([]DomainEvent, 0)
				}

				riseMap[h] = append(riseMap[h], e)
			}
		}
	}

	for h, evs := range riseMap {
		h.Handle(ctx, evs)
	}

	return nil
}

func (c *DomainEventChannel) Subscribe(ctx context.Context, eventType reflect.Type, h DomainEventHandler) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.handlers[eventType]; !ok {
		c.handlers[eventType] = make(map[DomainEventHandler]interface{})
	}

	if _, ok := c.handlers[eventType][h]; ok {
		return fmt.Errorf("handler already registred")
	}

	c.handlers[eventType][h] = h

	return nil
}

func (c *DomainEventChannel) Unsubscribe(ctx context.Context, eventType reflect.Type, h DomainEventHandler) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.handlers[eventType]; ok {
		delete(c.handlers[eventType], h)

		if len(c.handlers[eventType]) == 0 {
			delete(c.handlers, eventType)
		}
	}

	return nil
}

func NewDomainEventChannel() *DomainEventChannel {
	return &DomainEventChannel{
		handlers: make(map[reflect.Type]map[DomainEventHandler]interface{}),
		mu:       sync.Mutex{},
	}
}
