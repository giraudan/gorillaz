package main

import (
	"context"
	"flag"
	"github.com/opentracing/opentracing-go"
	"net/http"
	_ "net/http/pprof"
	"time"
)
import "github.com/skysoft-atm/gorillaz/stream"

func main() {
	var streamName string

	var port int
	flag.StringVar(&streamName, "stream", "", "stream to receive")

	flag.IntVar(&port, "port", 0, "tcp port to listen to")
	flag.Parse()

	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	p, err := stream.NewProvider(streamName)
	if err != nil {
		panic(err)
	}

	err = stream.Run(port)
	if err != nil {
		panic(err)
	}

	tick := time.Tick(time.Nanosecond * 5)
	for {
		<-tick
		sp := opentracing.StartSpan("sending message")
		sp.SetTag("super", "tag")
		ctx := opentracing.ContextWithSpan(context.Background(), sp)

		v := []byte("something wonderful")
		event := &stream.Event{
			Value: v,
			Ctx:   ctx,
		}
		sp.Finish()
		p.Submit(event)
	}

	p.Close()

	time.Sleep(time.Second * 5)
}
