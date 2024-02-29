package events

import (
	"encoding/json"
	"github.com/google/uuid"
)

const (
	ZoneCreatedEvent   = "ZoneCreated"
	ZoneDeletedEvent   = "ZoneDeleted"
	RecordAddedEvent   = "RecordAdded"
	RecordUpdatedEvent = "RecordUpdated"
	RecordDeletedEvent = "RecordDeleted"
)

type ZoneCreated struct {
	ZoneID uuid.UUID `json:"zone_id"`
	Zone   string    `json:"zone"`
}

type ZoneDeleted struct {
	ZoneID uuid.UUID `json:"zone_id"`
}

type RecordAdded struct {
	ZoneID     uuid.UUID       `json:"zone_id"`
	RecordID   uuid.UUID       `json:"record_id"`
	RecordType string          `json:"record_type"`
	Name       string          `json:"name"`
	TTL        uint32          `json:"ttl"`
	Value      json.RawMessage `json:"value"`
}

type RecordUpdated struct {
	ZoneID     uuid.UUID       `json:"zone_id"`
	RecordID   uuid.UUID       `json:"record_id"`
	RecordType string          `json:"record_type"`
	Name       string          `json:"name"`
	TTL        uint32          `json:"ttl"`
	Value      json.RawMessage `json:"value"`
}

type RecordDeleted struct {
	ZoneID   uuid.UUID `json:"zone_id"`
	RecordID uuid.UUID `json:"record_id"`
}
