package domain_events

import "reflect"

type DomainEvent interface {
	GetType() reflect.Type
}
