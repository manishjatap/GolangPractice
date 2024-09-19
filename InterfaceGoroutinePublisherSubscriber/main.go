package main

import (
	"fmt"
	"sync"
)

type MessageQueue interface {
	publisher()
	subscriber()
}

type Message struct {
	ID         int
	MessgeData string
}

type MessageQueueImplementation struct {
	Messages []Message
	wg       sync.WaitGroup
	ch       chan Message
	finish   chan bool
}

func main() {

	msg := MessageQueueImplementation{
		Messages: []Message{
			{ID: 1, MessgeData: "Hello 1"},
			{ID: 2, MessgeData: "Hello 2"},
			{ID: 3, MessgeData: "Hello 3"},
			{ID: 4, MessgeData: "Hello 4"},
			{ID: 5, MessgeData: "Hello 5"},
		},
		wg:     sync.WaitGroup{},
		ch:     make(chan Message),
		finish: make(chan bool),
	}

	msg.wg.Add(1)
	go msg.publisher()
	go msg.subscriber()

	msg.wg.Wait()
}

func (m *MessageQueueImplementation) publisher() {

	for _, msg := range m.Messages {
		fmt.Printf("Publisher has published message: %v\n", msg)
		m.ch <- msg
	}

	m.finish <- true
}

func (m *MessageQueueImplementation) subscriber() {
	for {
		select {
		case msg := <-m.ch:
			fmt.Printf("Subscriber has received message: %v\n", msg)
		case <-m.finish:
			fmt.Println("Finished the execution")
			m.wg.Done()
			break
		}
	}
}
