package repo

import "consumer-producer/internal/model"

type EventRepo interface {
	Lock(n uint64) ([]model.EntityEvent, error)
	Unlock(eventIDs []uint64) error

	Add(event []model.EntityEvent) error
	Remove(eventIDs []uint64) error
}
