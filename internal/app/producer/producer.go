package producer

import (
	"github.com/Turalchik/bot/internal/app/repo"
	"github.com/Turalchik/bot/internal/app/sender"
	"github.com/Turalchik/bot/internal/model/payment"
	"github.com/gammazero/workerpool"
	"sync"
	"time"
)

type Producer interface {
	Start()
	Close()
}

type producer struct {
	numberProducers uint64
	timeout         time.Duration
	bufferSize      uint64

	repo   repo.EventRepo
	sender sender.EventSender
	events <-chan payment.InvoiceEvent

	workerPool *workerpool.WorkerPool

	IDCreatedInvoices map[uint64]interface{}
	wg                *sync.WaitGroup
	done              chan interface{}
}

type Config struct {
	numberProducers uint64
	timeout         time.Duration
	bufferSize      uint64
	repo            repo.EventRepo
	sender          sender.EventSender
	events          <-chan payment.InvoiceEvent
	workerPool      *workerpool.WorkerPool
}

func NewProducer(cfg *Config) Producer {
	return &producer{
		numberProducers: cfg.numberProducers,
		timeout:         cfg.timeout,
		bufferSize:      cfg.bufferSize,

		repo:   cfg.repo,
		sender: cfg.sender,
		events: cfg.events,

		workerPool: cfg.workerPool,

		IDCreatedInvoices: make(map[uint64]interface{}),
		wg:                &sync.WaitGroup{},
		done:              make(chan interface{}),
	}
}

func (prod *producer) Start() {
	chanUnlockedEventIDs := make(chan payment.InvoiceEvent, prod.bufferSize)
	chanRemovedEventIDs := make(chan payment.InvoiceEvent, prod.bufferSize)

	mu := sync.Mutex{}

	for i := uint64(0); i < prod.numberProducers; i++ {
		prod.wg.Add(1)

		go func() {
			defer prod.wg.Done()
			ticker := time.NewTicker(prod.timeout)

			for {
				select {
				case event := <-prod.events:
					mu.Lock()
					_, ok := prod.IDCreatedInvoices[event.Entity.ID]
					mu.Unlock()

					if ok && event.Type == payment.Created || !ok && (event.Type == payment.Updated || event.Type == payment.Removed) {
						chanUnlockedEventIDs <- event
						continue
					}

					if err := prod.sender.Send(&event); err != nil {
						chanUnlockedEventIDs <- event
						continue
					}

					chanRemovedEventIDs <- event

				case <-ticker.C:
					_, sliceUnlockedIDs := chanInvoiceEvents2SliceInvoiceEventsANDIDs(chanUnlockedEventIDs)
					if len(sliceUnlockedIDs) > 0 {
						prod.workerPool.Submit(func() {
							if err := prod.repo.Unlock(sliceUnlockedIDs); err != nil {
								for _, id := range sliceUnlockedIDs {
									chanUnlockedEventIDs <- payment.InvoiceEvent{ID: id}
								}
							}
						})
					}
					sliceRemovedEvents, sliceRemovedIDs := chanInvoiceEvents2SliceInvoiceEventsANDIDs(chanRemovedEventIDs)
					if len(sliceRemovedIDs) > 0 {
						prod.workerPool.Submit(func() {
							if err := prod.repo.Remove(sliceRemovedIDs); err != nil {
								for _, event := range sliceRemovedEvents {
									chanRemovedEventIDs <- event
								}
							} else {
								mu.Lock()
								for _, event := range sliceRemovedEvents {
									if event.Type == payment.Created {
										prod.IDCreatedInvoices[event.Entity.ID] = struct{}{}
									} else if event.Type == payment.Removed {
										delete(prod.IDCreatedInvoices, event.Entity.ID)
									}
								}
								mu.Unlock()
							}
						})
					}
				case <-prod.done:
					return
				}
			}
		}()
	}
}

func (prod *producer) Close() {
	close(prod.done)
	prod.wg.Wait()
}
