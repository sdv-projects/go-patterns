package handlers

import (
	"context"
	"fmt"
	domain_events "sdv-projects/go-patterns/domain/events"
	domain "sdv-projects/go-patterns/exapmle/event_bus/domain"
)

type UserCreatedEventHandler struct {
}

func (h *UserCreatedEventHandler) Handle(ctx context.Context, events []domain_events.DomainEvent) {
	for _, e := range events {
		switch event := e.(type) {
		case *domain.UserCreatedEvent:
			fmt.Printf("User created: %v", event.UserId)

		default:
			fmt.Printf("Other event: %s", event.GetEventType().Name())
		}
	}
}
