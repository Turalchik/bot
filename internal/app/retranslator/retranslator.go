package retranslator

import (
	"github.com/Turalchik/bot/internal/app/consumer"
	"github.com/Turalchik/bot/internal/app/producer"
	"github.com/Turalchik/bot/internal/app/repo"
	"github.com/Turalchik/bot/internal/app/sender"
	"github.com/Turalchik/bot/internal/model/payment"
	"github.com/gammazero/workerpool"
	"time"
)

type Retranslator interface {
	Start()
	Close()
}

type Config struct {
	ConsumerNumber    uint64
	ConsumerTimeout   time.Duration
	ConsumerBatchSize uint64

	ProducerNumber     uint64
	ProducerTimeout    time.Duration
	ProducerBufferSize uint64

	ChannelEventsSize uint64
	WorkerCount       int

	Repo   repo.EventRepo
	Sender sender.EventSender
}

type retranslator struct {
	events     chan payment.InvoiceEvent
	consumer   consumer.Consumer
	producer   producer.Producer
	workerPool *workerpool.WorkerPool
}

func NewRetranslator(cfg Config) Retranslator {
	events := make(chan payment.InvoiceEvent, cfg.ChannelEventsSize)
	workerPool := workerpool.New(cfg.WorkerCount)

	producerCfg := producer.Config{
		NumberProducers: cfg.ProducerNumber,
		Timeout:         cfg.ProducerTimeout,
		BufferSize:      cfg.ProducerBufferSize,
		Repo:            cfg.Repo,
		Sender:          cfg.Sender,
		Events:          events,
		WorkerPool:      workerPool,
	}

	consumerCfg := consumer.Config{
		NumberConsumers: cfg.ConsumerNumber,
		Events:          events,
		Repo:            cfg.Repo,
		BatchSize:       cfg.ConsumerBatchSize,
		Timeout:         cfg.ConsumerTimeout,
	}

	return &retranslator{
		events:     events,
		consumer:   consumer.NewConsumer(consumerCfg),
		producer:   producer.NewProducer(producerCfg),
		workerPool: workerPool,
	}
}

func (retrans *retranslator) Start() {
	retrans.producer.Start()
	retrans.consumer.Start()
}

func (retrans *retranslator) Close() {
	retrans.consumer.Close()
	retrans.producer.Close()
	retrans.workerPool.StopWait()
}
