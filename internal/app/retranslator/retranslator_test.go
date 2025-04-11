package retranslator

import (
	"testing"
	"time"

	"github.com/Turalchik/bot/internal/mocks"
	"github.com/golang/mock/gomock"
)

func TestStart(t *testing.T) {

	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)

	repo.EXPECT().Lock(gomock.Any()).AnyTimes()

	cfg := Config{
		ConsumerNumber:     2,
		ConsumerTimeout:    10 * time.Second,
		ConsumerBatchSize:  10,
		ProducerNumber:     2,
		ProducerTimeout:    10 * time.Second,
		ProducerBufferSize: 128,
		ChannelEventsSize:  512,
		WorkerCount:        2,
		Repo:               repo,
		Sender:             sender,
	}

	retranslator := NewRetranslator(cfg)
	retranslator.Start()
	retranslator.Close()
}
