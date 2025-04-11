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
	repo            repo.EventRepo
	sender          sender.EventSender
	events          <-chan payment.InvoiceEvent
	workerPool      *workerpool.WorkerPool
}

func NewProducer(cfg *Config) Producer {
	return &producer{
		numberProducers: cfg.numberProducers,
		timeout:         cfg.timeout,

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
	unlockedEventIDs := make([]uint64, 0)
	removedEventIDs := make([]uint64, 0)
	mu := sync.Mutex{}

	for i := uint64(0); i < prod.numberProducers; i++ {
		prod.wg.Add(1)

		go func() {
			defer prod.wg.Done()
			ticker := time.NewTicker(prod.timeout)

			for {
				select {
				case event := <-prod.events:
					_, ok := prod.IDCreatedInvoices[event.Entity.ID]
					if ok && event.Type == payment.Created {
						mu.Lock()
						unlockedEventIDs = append(unlockedEventIDs, event.ID)
						mu.Unlock()
						continue
					}
					if !ok && (event.Type == payment.Updated || event.Type == payment.Removed) {
						mu.Lock()
						unlockedEventIDs = append(unlockedEventIDs, event.ID)
						mu.Unlock()
						continue
					}

					if err := prod.sender.Send(&event); err != nil {
						mu.Lock()
						unlockedEventIDs = append(unlockedEventIDs, event.ID)
						mu.Lock()
						continue
					}

					mu.Lock()
					removedEventIDs = append(removedEventIDs, event.ID)
					mu.Unlock()

					if event.Type == payment.Created {
						mu.Lock()
						prod.IDCreatedInvoices[event.Entity.ID] = struct{}{}
						mu.Unlock()
					} else if event.Type == payment.Removed {
						mu.Lock()
						delete(prod.IDCreatedInvoices, event.Entity.ID)
						mu.Unlock()
					}
				case <-ticker.C:
					mu.Lock()
					if len(unlockedEventIDs) > 0 {
						prod.workerPool.Submit(func() {
							if err := prod.repo.Unlock(unlockedEventIDs); err == nil {
								unlockedEventIDs = unlockedEventIDs[:0]
							}
						})
					}
					mu.Unlock()
					mu.Lock()
					if len(removedEventIDs) > 0 {
						prod.workerPool.Submit(func() {
							if err := prod.repo.Remove(removedEventIDs); err == nil {
								removedEventIDs = removedEventIDs[:0]
							}
						})
					}
					mu.Unlock()
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
