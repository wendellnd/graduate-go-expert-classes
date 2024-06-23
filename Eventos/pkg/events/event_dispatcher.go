package events

import (
	"errors"
	"sync"
)

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (d *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if _, ok := d.handlers[eventName]; ok {
		for _, h := range d.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}

	d.handlers[eventName] = append(d.handlers[eventName], handler)
	return nil
}

func (d *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	if _, ok := d.handlers[eventName]; ok {
		for i, h := range d.handlers[eventName] {
			if h == handler {
				d.handlers[eventName] = append(d.handlers[eventName][:i], d.handlers[eventName][i+1:]...)
				return nil
			}
		}
	}

	return nil
}

func (d *EventDispatcher) Clear() error {
	d.handlers = make(map[string][]EventHandlerInterface)
	return nil
}

func (d *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	if _, ok := d.handlers[eventName]; ok {
		for _, h := range d.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}

	return false
}

func (d *EventDispatcher) Dispatch(event EventInterface) error {
	if _, ok := d.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}
		for _, handler := range d.handlers[event.GetName()] {
			wg.Add(1)
			go handler.Handle(event, wg)
		}
		wg.Wait()
	}

	return nil
}
