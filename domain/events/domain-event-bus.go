package domain_events

import (
	"context"
	"reflect"
)

type DomainEventBus interface {
	Publish(ctx context.Context, events ...DomainEvent) error
	Subscribe(ctx context.Context, eventType reflect.Type, h DomainEventHandler) error
	Unsubscribe(ctx context.Context, eventType reflect.Type, h DomainEventHandler) error
}
