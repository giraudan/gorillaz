/*
Provides pubsub of messages over channels.
A provider has a broadcaster into which it Submits messages and into
which subscribers Register to pick up those messages.

If one of the subscribers is not able to consume the message, then messages will be dropped for this consumer.
It is possible to pass a function to handle dropped messages.

*/
package mux

import (
	"fmt"
)

type Broadcaster struct {
	input   chan interface{}
	reg     chan registration
	unreg   chan unregistration
	outputs map[chan<- interface{}]ConsumerConfig
	*BroadcasterConfig
}

// Register a new channel to receive broadcasts
func (b *Broadcaster) Register(newch chan<- interface{}, options ...ConsumerOptionFunc) error {
	done := make(chan bool)
	config := &ConsumerConfig{}
	for _, option := range options {
		if err := option(config); err != nil {
			return err
		}
	}
	b.reg <- registration{consumer{*config, newch}, done}
	<-done
	return nil
}

// Unregister a channel so that it no longer receives broadcasts.
func (b *Broadcaster) Unregister(newch chan<- interface{}) {
	done := make(chan bool)
	b.unreg <- unregistration{newch, done}
	<-done
}

// Shut this StateBroadcaster down.
func (b *Broadcaster) Close() error {
	close(b.reg)
	return nil
}

// Submit a new object to all subscribers, this call can block if the input channel is full
func (b *Broadcaster) SubmitBlocking(i interface{}) error {
	if b != nil && i != nil {
		b.input <- i
		return nil
	}
	return fmt.Errorf("nil value")
}

// Submit a new object to all subscribers, this call will drop the message if the input channel is full
func (b *Broadcaster) SubmitNonBlocking(i interface{}) error {
	if b != nil && i != nil {
		select {
		case b.input <- i:
			return nil
		default:
			return fmt.Errorf("value dropped")
		}
	}
	return fmt.Errorf("nil value")
}

func (b *Broadcaster) broadcast(m interface{}) {
	for ch := range b.outputs {
		select {
		case ch <- m:
			//message sent
		default:
			//consumer is not ready to receive a message, drop it and execute provided action on backpressure
			subConfig := b.outputs[ch]
			if subConfig.onBackpressure != nil {
				subConfig.onBackpressure(m)
			}
		}
	}
	if b.postBroadcast != nil {
		b.postBroadcast(m)
	}
}

// onBackPressureState can be nil
func (b *Broadcaster) run() {
	subscriberCount := 0
	for {
		if b.eagerBroadcast == false {
			// wait for first consumer
			r, ok := <-b.reg
			if ok {
				subscriberCount = b.addSubscriber(r, subscriberCount)
			} else {
				// close all output channel to notify them that the broadcaster is closed
				for output := range b.outputs {
					close(output)
				}
			}
		}
		for b.eagerBroadcast == true || subscriberCount != 0 {
			select {
			case r, ok := <-b.reg:
				if ok {
					subscriberCount = b.addSubscriber(r, subscriberCount)
				} else {
					return
				}
			case u := <-b.unreg:
				delete(b.outputs, u.channel)
				subscriberCount--
				u.done <- true
			case m := <-b.input:
				b.broadcast(m)

			}
		}
	}
}

func (b *Broadcaster) addSubscriber(r registration, subscriberCount int) int {
	b.outputs[r.consumer.channel] = r.consumer.config
	r.done <- true
	return subscriberCount + 1
}

// NewBroadcaster creates a new Broadcaster with the given input channel buffer length.
// onBackPressureState is an action to execute when messages are dropped on back pressure (typically logging), it can be nil
func NewNonBlockingBroadcaster(bufLen int, options ...BroadcasterOptionFunc) (*Broadcaster, error) {
	b := &Broadcaster{
		input:             make(chan interface{}, bufLen),
		reg:               make(chan registration),
		unreg:             make(chan unregistration),
		outputs:           make(map[chan<- interface{}]ConsumerConfig),
		BroadcasterConfig: &BroadcasterConfig{eagerBroadcast: true},
	}

	for _, option := range options {
		if err := option(b.BroadcasterConfig); err != nil {
			return nil, err
		}
	}

	go b.run()
	return b, nil
}
