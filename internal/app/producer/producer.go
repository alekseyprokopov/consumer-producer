package producer

import (
	"consumer-producer/internal/app/sender"
	"consumer-producer/internal/model"
	"github.com/gammazero/workerpool"
	"sync"
	"time"
)

type Producer interface {
	Start()
	Close()
}

type producer struct {
	n       uint64
	timeout time.Duration

	sender sender.EventSender
	events <-chan model.EntityEvent

	workerPool *workerpool.WorkerPool

	wg   *sync.WaitGroup
	done chan bool
}

func NewProducer(
	n uint64,
	sender sender.EventSender,
	events <-chan model.EntityEvent,
	workerPool *workerpool.WorkerPool,
) Producer {
	wg := &sync.WaitGroup{}
	done := make(chan bool)
	return &producer{
		n:          n,
		sender:     sender,
		events:     events,
		workerPool: workerPool,
		wg:         wg,
		done:       done,
	}
}

func (p *producer) Start() {
	for i := uint64(0); i < p.n; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for {
				select {
				case event := <-p.events:
					if err := p.sender.Send(&event); err != nil {
						p.workerPool.Submit(func() {

						})
					} else {
						p.workerPool.Submit(func() {

						})
					}
				case <-p.done:
					return

				}
			}
		}()

	}

}

func (p *producer) Close() {
	close(p.done)
	p.wg.Wait()
}
