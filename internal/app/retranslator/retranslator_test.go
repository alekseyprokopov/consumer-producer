package retranslator

import (
	"consumer-producer/internal/mocks"
	"github.com/golang/mock/gomock"
	"log"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	log.Printf("HELLO")
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)

	repo.EXPECT().Lock(gomock.Any()).AnyTimes()

	cfg := Config{
		ChannelSize:     512,
		ConsumerCount:   2,
		ConsumerSize:    10,
		ConsumerTimeout: 10 * time.Second,
		ProducerCount:   2,
		WorkerCount:     2,
		Repo:            repo,
		Sender:          sender,
	}

	retranslator := NewRetranslator(cfg)
	retranslator.Start()
	retranslator.Close()

}
