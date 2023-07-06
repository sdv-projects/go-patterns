package domain_events

import "reflect"

type DomainEvent interface {
	GetEventType() reflect.Type
}
