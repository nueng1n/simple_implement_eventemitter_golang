package main

import (
	"context"
	"log"
	"sync"
)

type MessageChannel struct {
	EventName string `json:"even_name"`
	Message   string `json:"message"`
}

type eventHandler struct {
	callback func(string)
}

type EventEmitter struct {
	ch     chan *MessageChannel
	events map[string]eventHandler
}

func NewEventEmitter(ctx context.Context, wg *sync.WaitGroup) *EventEmitter {

	eventEmitter := &EventEmitter{
		ch:     make(chan *MessageChannel),
		events: make(map[string]eventHandler),
	}

	go func() {

		for {

			select {
			case message := <-eventEmitter.ch:
				if handler, ok := eventEmitter.events[message.EventName]; ok {
					handler.callback(message.Message)

					wg.Done()
				}
			case <-ctx.Done():
				return

			}
		}

	}()

	return eventEmitter

}

func (emitter *EventEmitter) On(eventName string, callback func(string)) {
	emitter.events[eventName] = eventHandler{
		callback: callback,
	}

}

func (emitter *EventEmitter) Emit(eventName string, data string) {

	message := &MessageChannel{
		EventName: eventName,
		Message:   data,
	}

	emitter.ch <- message
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	emitter := NewEventEmitter(ctx, &wg)

	emitter.On("example1", func(str string) {
		log.Println("Received message  on `example1` event", str)
	})

	emitter.On("example2", func(str string) {
		log.Println("Received message  on `example2` event", str)
	})

	wg.Add(1)
	emitter.Emit("example1", "send a")

	wg.Add(1)
	emitter.Emit("example2", "send b")
	wg.Add(1)
	emitter.Emit("example1", "send c")

	wg.Wait()
}
