package stream

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	gaz "github.com/skysoft-atm/gorillaz"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
	"math"
	"strings"
	"sync"
	"time"
)

var mu sync.RWMutex

func NewConsumer(streamName string, endpoints ...string) (chan *Event, error) {
	// TODO: hacky hack to create a resolver to use with round robin
	mu.Lock()
	r, _ := manual.GenerateAndRegisterManualResolver()
	mu.Unlock()

	addresses := make([]resolver.Address, len(endpoints))
	for i := 0; i < len(endpoints); i++ {
		addresses[i] = resolver.Address{Addr: endpoints[i]}
	}
	r.InitialAddrs(addresses)
	target := r.Scheme() + ":///fake"

	ch := make(chan *Event, 256)
	go func() {
		run(streamName, target, endpoints, ch)
	}()
	return ch, nil
}

func run(streamName string, target string, endpoints []string, ch chan *Event) {
	receivedCounter := promauto.NewCounter(prometheus.CounterOpts{
		Name: "received_events",
		Help: "The total number of events received",
		ConstLabels: prometheus.Labels{
			"stream":    streamName,
			"endpoints": strings.Join(endpoints, ","),
		},
	})

	conCounter := promauto.NewCounter(prometheus.CounterOpts{
		Name: "connection_attempts",
		Help: "The total number of connections to the stream",
		ConstLabels: prometheus.Labels{
			"stream":    streamName,
			"endpoints": strings.Join(endpoints, ","),
		},
	})

	delaySummary := promauto.NewSummary(prometheus.SummaryOpts{
		Name:       "streaming_delay_ms",
		Help:       "The distribution of delay between when messages are sent to from the consumer and when they are received, in milliseconds",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		ConstLabels: prometheus.Labels{
			"stream":    streamName,
			"endpoints": strings.Join(endpoints, ","),
		},
	})

	var streamClient Stream_StreamClient
	var err error
connect:
	for {
		conCounter.Inc()
		gaz.Log.Info("connection attempt to stream", zap.String("stream", streamName))
		streamClient, err = initConn(target, streamName)
		if err == nil {
			gaz.Log.Info("successful connection attempt to stream", zap.String("stream", streamName))
			break
		} else {
			gaz.Log.Error("connection attempt to stream failed, retry in 1 s", zap.String("stream", streamName), zap.Error(err))
			time.Sleep(time.Duration(time.Second))
		}
	}
	for {
		streamEvt, err := streamClient.Recv()
		if err != nil {
			gaz.Log.Error("stream is unavailable, retry connection in 1s", zap.String("stream", streamName), zap.Error(err))
			time.Sleep(time.Second)
			goto connect
		}
		gaz.Log.Debug("event received", zap.String("stream", streamName))
		receivedCounter.Inc()
		receptTime := time.Now()
		delaySummary.Observe(math.Max(0, float64(receptTime.UnixNano())/1000000.0-float64(streamEvt.Stream_Timestamp_Ns)/1000000.0))
		ch <- &Event{
			Key:             streamEvt.Key,
			Value:           streamEvt.Value,
			StreamTimestamp: streamEvt.Stream_Timestamp_Ns,
		}
	}
}

func initConn(target string, streamName string) (Stream_StreamClient, error) {
	mu.RLock()
	conn, err := grpc.Dial(target, grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
	mu.RUnlock()
	if err != nil {
		return nil, err
	}
	c := NewStreamClient(conn)
	req := &StreamRequest{Name: streamName}
	return c.Stream(context.TODO(), req)
}
