package consumer

import (
	"consumer-producer/internal/app/repo"
	"consumer-producer/internal/model"
	"sync"
	"time"
)

type Consumer interface {
	Start()
	Close()
}

type consumer struct {
	n      uint64
	events chan<- model.EntityEvent

	repo repo.EventRepo

	batchSize uint64
	timeout   time.Duration

	done chan bool
	wg   *sync.WaitGroup
}

type Config struct {
	n         uint64
	events    chan<- model.EntityEvent
	repo      repo.EventRepo
	batchSize uint64
	timeout   time.Duration
}

func NewdbConsumer(
	n uint64,
	events chan<- model.EntityEvent,
	repo repo.EventRepo,
	batchSize uint64,
	consumerTimeout time.Duration) Consumer {
	wg := &sync.WaitGroup{}
	done := make(chan bool)
	return &consumer{
		n:         n,
		events:    events,
		repo:      repo,
		batchSize: batchSize,
		timeout:   consumerTimeout,
		wg:        wg,
		done:      done,
	}
}

func (c *consumer) Start() {

	for i := uint64(0); i < c.n; i++ {
		c.wg.Add(1)
		go func() {
			defer c.wg.Done()
			ticker := time.NewTicker(c.timeout)
			for {

				select {
				case <-ticker.C:
					events, err := c.repo.Lock(c.batchSize)
					if err != nil {
						continue
					}
					for _, event := range events {
						c.events <- event
					}
				case <-c.done:
					return
				}

			}

		}()

	}
}

func (c *consumer) Close() {
	close(c.events)
	c.wg.Wait()
}
