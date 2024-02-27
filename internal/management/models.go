package management

import (
	"encoding/json"
	"github.com/google/uuid"
)

type RecordType string

const (
	A     RecordType = "A"
	AAAA  RecordType = "AAAA"
	MX    RecordType = "MX"
	NS    RecordType = "NS"
	TXT   RecordType = "TXT"
	CNAME RecordType = "CNAME"
	PTR   RecordType = "PTR"
)

type Zone struct {
	ID   uuid.UUID
	Zone string
}

type Record struct {
	ID     uuid.UUID
	ZoneID uuid.UUID
	Name   string
	Type   RecordType
	TTL    uint32
	Value  json.RawMessage
}
