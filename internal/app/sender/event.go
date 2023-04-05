package sender

import "consumer-producer/internal/model"

type EventSender interface {
	Send(entity *model.EntityEvent) error
}
