package resolver

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Zone struct {
	ZoneID uuid.UUID `bson:"zone_id"`
	Zone   string    `bson:"zone"`
}

type Record struct {
	RecordID   uuid.UUID       `bson:"record_id"`
	ZoneID     uuid.UUID       `bson:"zone_id"`
	Zone       string          `bson:"zone"`
	Name       string          `bson:"name"`
	FullName   string          `bson:"full_name"` // concatenated of "{Name}.{Zone}."
	RecordType string          `bson:"record_type"`
	TTL        uint32          `bson:"ttl"`
	Value      json.RawMessage `bson:"value"`
}
