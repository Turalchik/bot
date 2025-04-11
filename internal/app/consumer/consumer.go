package consumer

import (
	"github.com/Turalchik/bot/internal/app/repo"
	"github.com/Turalchik/bot/internal/model/payment"
	"sync"
	"time"
)

type Consumer interface {
	Start()
	Close()
}

type consumer struct {
	numberConsumers uint64
	events          chan<- payment.InvoiceEvent

	repo repo.EventRepo

	batchSize uint64
	timeout   time.Duration

	done chan interface{}
	wg   *sync.WaitGroup
}

type Config struct {
	numberConsumers uint64
	events          chan<- payment.InvoiceEvent
	repo            repo.EventRepo
	batchSize       uint64
	timeout         time.Duration
}

func NewConsumer(cfg *Config) Consumer {
	return &consumer{
		numberConsumers: cfg.numberConsumers,
		events:          cfg.events,

		repo: cfg.repo,

		batchSize: cfg.batchSize,
		timeout:   cfg.timeout,

		done: make(chan interface{}),
		wg:   &sync.WaitGroup{},
	}
}

func (cons *consumer) Start() {
	for i := uint64(0); i < cons.numberConsumers; i++ {
		cons.wg.Add(1)

		go func() {
			defer cons.wg.Done()
			ticker := time.NewTicker(cons.timeout)
			for {
				select {
				case <-ticker.C:
					events, err := cons.repo.Lock(cons.batchSize)
					if err != nil {
						continue
					}
					for _, event := range events {
						cons.events <- event
					}
				case <-cons.done:
					return
				}
			}
		}()
	}
}

func (cons *consumer) Close() {
	close(cons.done)
	cons.wg.Wait()
}
