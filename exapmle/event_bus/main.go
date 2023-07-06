package main

import (
	"context"
	"reflect"
	domain_events "sdv-projects/go-patterns/domain/events"
	"sdv-projects/go-patterns/exapmle/event_bus/domain"
	"sdv-projects/go-patterns/exapmle/event_bus/handlers"
	"time"
)

var (
	userCreatedEvent *domain.UserCreatedEvent = nil
)

func main() {
	ctx := context.Background()

	var eventBus domain_events.DomainEventBus = domain_events.NewDomainEventChannel()

	userCreatedHandler := &handlers.UserCreatedEventHandler{}
	eventBus.Subscribe(ctx, reflect.TypeOf(userCreatedEvent), userCreatedHandler)

	userCreatedEvent := domain.UserCreatedEvent{
		UserId: 123,
	}

	eventBus.Publish(ctx, &userCreatedEvent)

	time.Sleep(5 * time.Second)
}
