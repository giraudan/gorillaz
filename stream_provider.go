package gorillaz

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/common/log"
	"github.com/skysoft-atm/gorillaz/mux"
	"github.com/skysoft-atm/gorillaz/stream"
	"go.uber.org/zap"
	"sync"
)

// NewStreamProvider returns a new provider ready to be used.
// only one instance of provider should be created for a given streamName
func (g *Gaz) NewStreamProvider(streamName string, opts ...ProviderConfigOpt) (*StreamProvider, error) {
	Log.Info("creating stream", zap.String("stream", streamName))

	config := defaultProviderConfig()
	for _, opt := range opts {
		opt(config)
	}

	var broadcaster *mux.Broadcaster
	var err error

	if config.LazyBroadcast {
		broadcaster, err = mux.NewNonBlockingBroadcaster(config.InputBufferLen, mux.LazyBroadcast)
	} else {
		broadcaster, err = mux.NewNonBlockingBroadcaster(config.InputBufferLen)
	}
	if err != nil {
		Log.Error("could not create stream broadcaster", zap.Error(err))
		return nil, err
	}
	p := &StreamProvider{
		streamName:  streamName,
		config:      config,
		broadcaster: broadcaster,
		metrics:     pMetricHolder(streamName),
	}
	g.streamRegistry.register(streamName, p)

	return p, nil
}

func (g *Gaz) CloseStream(streamName string) error {
	log.Info("closing stream", zap.String("stream", streamName))
	provider, ok := g.streamRegistry.find(streamName)
	if !ok {
		return fmt.Errorf("cannot find stream " + streamName)
	}
	g.streamRegistry.unregister(streamName)
	provider.close()
	return nil
}

type StreamProvider struct {
	streamName  string
	config      *ProviderConfig
	broadcaster *mux.Broadcaster
	metrics     providerMetricsHolder
}

var pMetricHolderMu sync.Mutex
var pMetrics = make(map[string]providerMetricsHolder)

func pMetricHolder(streamName string) providerMetricsHolder {
	pMetricHolderMu.Lock()
	defer pMetricHolderMu.Unlock()
	if h, ok := pMetrics[streamName]; ok {
		return h
	}

	h := providerMetricsHolder{
		sentCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "stream_event_sent",
			Help: "The total number of messages sent",
			ConstLabels: prometheus.Labels{
				"stream": streamName,
			},
		}),

		backPressureCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "stream_backpressure_dropped",
			Help: "The total number of messages dropped due to backpressure",
			ConstLabels: prometheus.Labels{
				"stream": streamName,
			},
		}),

		clientCounter: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "stream_connected_clients",
			Help: "The total number of clients connected",
			ConstLabels: prometheus.Labels{
				"stream": streamName,
			},
		}),

		lastEventTimestamp: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "stream_last_evt_timestamp",
			Help: "Timestamp of the last event produced",
			ConstLabels: prometheus.Labels{
				"stream": streamName,
			},
		}),
	}
	pMetrics[streamName] = h
	return h
}

type providerMetricsHolder struct {
	sentCounter         prometheus.Counter
	backPressureCounter prometheus.Counter
	clientCounter       prometheus.Gauge
	lastEventTimestamp  prometheus.Gauge
}

// ProviderConfig is the configuration that will be applied for the stream StreamProvider
type ProviderConfig struct {
	InputBufferLen           int                     // InputBufferLen is the size of the input channel (default: 256)
	SubscriberInputBufferLen int                     // SubscriberInputBufferLen is the size of the channel used to forward events to each client. (default: 256)
	OnBackPressure           func(streamName string) // OnBackPressure is the function called when a customer cannot consume fast enough and event are dropped. (default: log)
	LazyBroadcast            bool                    // if lazy broadcaster, then the provider doesn't consume messages as long as there is no consumer
}

func defaultProviderConfig() *ProviderConfig {
	return &ProviderConfig{
		InputBufferLen:           256,
		SubscriberInputBufferLen: 256,
		OnBackPressure: func(streamName string) {
			Log.Warn("backpressure applied, an event won't be delivered because it can't consume fast enough", zap.String("stream", streamName))
		},
		LazyBroadcast: false,
	}
}

// ProviderConfigOpt is a ProviderConfig option function to modify the ProviderConfig used by the stream StreamProvider
type ProviderConfigOpt func(p *ProviderConfig)

var LazyBroadcast = func(p *ProviderConfig) {
	p.LazyBroadcast = true
}

// Submit pushes the event to all subscribers
func (p *StreamProvider) Submit(evt *stream.Event) {
	metadata, err := stream.ContextToMetadata(evt.Ctx)
	if err != nil {
		Log.Error("error while creating Metadata from event.Context", zap.Error(err))
	}
	streamEvent := &stream.StreamEvent{
		Key:      evt.Key,
		Value:    evt.Value,
		Metadata: metadata,
	}
	p.metrics.sentCounter.Inc()
	p.metrics.lastEventTimestamp.SetToCurrentTime()
	p.broadcaster.SubmitBlocking(streamEvent)
}

func (p *StreamProvider) sendLoop(streamName string, strm stream.Stream_StreamServer) {
	p.metrics.clientCounter.Inc()
	broadcaster := p.broadcaster
	streamCh := make(chan interface{}, p.config.SubscriberInputBufferLen)
	broadcaster.Register(streamCh, func(config *mux.ConsumerConfig) error {
		config.OnBackpressure(func(interface{}) {
			p.config.OnBackPressure(streamName)
			p.metrics.backPressureCounter.Inc()
		})
		return nil
	})

	for val := range streamCh {
		evt := val.(*stream.StreamEvent)
		err := strm.Send(evt)
		if err != nil {
			Log.Info("consumer disconnected", zap.Error(err))
			broadcaster.Unregister(streamCh)
			break
		}
	}
	p.metrics.clientCounter.Dec()
}

func (p *StreamProvider) close() {
	p.broadcaster.Close()
}

type streamRegistry struct {
	sync.RWMutex
	providers map[string]*StreamProvider
}

func (r *streamRegistry) find(streamName string) (*StreamProvider, bool) {
	r.RLock()
	p, ok := r.providers[streamName]
	r.RUnlock()
	return p, ok
}

func (r *streamRegistry) register(streamName string, p *StreamProvider) {
	r.Lock()
	if _, found := r.providers[streamName]; found {
		panic("cannot register 2 providers with the same streamName")
	}
	r.providers[streamName] = p
	r.Unlock()
}

func (r *streamRegistry) unregister(streamName string) {
	r.Lock()
	delete(r.providers, streamName)
	r.Unlock()
}

// Stream implements streaming.proto Stream.
// should not be called by the client
func (r *streamRegistry) Stream(req *stream.StreamRequest, strm stream.Stream_StreamServer) error {
	Log.Info("new stream consumer", zap.String("stream", req.Name))
	streamName := req.Name
	r.RLock()
	provider, ok := r.providers[req.Name]
	r.RUnlock()
	if !ok {
		Log.Error("unknown stream %s", zap.String("stream", streamName))
		return fmt.Errorf("unknown stream %s", streamName)
	}
	provider.sendLoop(streamName, strm)
	return nil
}
