package resolver

import (
	"encoding/json"
	"github.com/google/uuid"
	"net"
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

type Answer[T any] struct {
	TTL   uint32 `json:"ttl"`
	Value T      `bson:"value"`
}

type ARecord []struct {
	IP net.IP `json:"ip"`
}
type AAAARecord []struct {
	IP net.IP `json:"ip"`
}
type MXRecord []struct {
	Preference int    `json:"preference"`
	Domain     string `json:"domain"`
}
type NSRecord []struct {
	Domain string `json:"domain"`
}
type TXTRecord []struct {
	Text string `json:"text"`
}
type CNAMERecord struct {
	Target string `json:"target"`
}
type PTRRecord struct {
	Domain string `json:"domain"`
}
