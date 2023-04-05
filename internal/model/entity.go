package model

type EventType uint64
type EventStatus uint64

type Entity struct {
	ID uint64
}

const (
	Created EventType = iota
	Updated
	Removed

	Deferred EventStatus = iota
	Processed
)

type EntityEvent struct {
	ID     uint64
	Type   EventType
	Status EventStatus
	Entity *Entity
}
