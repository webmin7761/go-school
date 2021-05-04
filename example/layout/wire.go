package main

import "github.com/google/wire"

type Message string

type Greeter struct {
	Message Message
}

type Event struct {
	Greeter Greeter
}

//go get github.com/google/wire/cmd/wire
//wire

func NewMessage() Message {
	return Message("Hi there!")
}

func NewGreeter(m Message) (Greeter, error) {
	return Greeter{Message: m}, nil
}

func NewEvent(g Greeter) (Event, error) {
	return Event{Greeter: g}, nil
}

func InitializeEvent() (Event, error) {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}, nil
}
