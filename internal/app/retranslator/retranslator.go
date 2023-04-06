package retranslator

import (
	"consumer-producer/internal/app/consumer"
	"consumer-producer/internal/app/producer"
	"consumer-producer/internal/app/repo"
	"consumer-producer/internal/app/sender"
	"consumer-producer/internal/model"
	"github.com/gammazero/workerpool"
	"time"
)

type Retranslator interface {
	Start()
	Close()
}

type Config struct {
	ChannelSize uint64

	ConsumerSize    uint64
	ConsumerCount   uint64
	ConsumerTimeout time.Duration

	ProducerCount uint64
	WorkerCount   int

	Repo   repo.EventRepo
	Sender sender.EventSender
}

type retranslator struct {
	events     chan model.EntityEvent
	producer   producer.Producer
	consumer   consumer.Consumer
	workerPool *workerpool.WorkerPool
}

func NewRetranslator(cfg Config) Retranslator {
	events := make(chan model.EntityEvent)
	workerPool := workerpool.New(cfg.WorkerCount)

	consumer := consumer.NewdbConsumer(
		cfg.ConsumerCount,
		events,
		cfg.Repo,
		cfg.ConsumerSize,
		cfg.ConsumerTimeout,
	)

	producer := producer.NewProducer(
		cfg.ProducerCount,
		cfg.Sender,
		events,
		workerPool,
	)

	return &retranslator{
		events:     events,
		producer:   producer,
		consumer:   consumer,
		workerPool: workerPool,
	}
}

func (r *retranslator) Start() {
	r.producer.Start()
	r.consumer.Start()
}

func (r *retranslator) Close() {
	r.producer.Close()
	r.consumer.Close()
	r.workerPool.StopWait()
}
