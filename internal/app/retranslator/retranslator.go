package retranslator

import (
	"consumer-producer/internal/app/consumer"
	"consumer-producer/internal/app/producer"
	"consumer-producer/internal/app/repo"
	"consumer-producer/internal/app/sender"
	"consumer-producer/internal/model"
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
	WorkerCount   uint64

	Repo   repo.EventRepo
	Sender sender.EventSender
}

type retranslator struct {
	events   chan model.EntityEvent
	producer producer.Producer
	consumer consumer.Consumer
	//workerPool workerpool.WorkerPool
}
