package domain

import "reflect"

type UserCreatedEvent struct {
	UserId int
}

func (e *UserCreatedEvent) GetEventType() reflect.Type {
	return reflect.TypeOf(e)
}
