package domain_events

import "context"

type DomainEventHandler interface {
	Handle(ctx context.Context, events []DomainEvent)
}
