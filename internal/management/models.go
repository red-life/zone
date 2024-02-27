package management

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
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
	ID   uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Zone string    `gorm:"unique;not null"`
}

type Record struct {
	ID     uuid.UUID    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	ZoneID uuid.UUID    `gorm:"type:uuid"`
	Name   string       `gorm:"type:varchar(255);unique;not null"`
	Type   RecordType   `gorm:"type:varchar(7);unique;not null"` // not using enums due to further update that might support other record types
	TTL    uint32       `gorm:"type:integer;check:ttl >= 0;not null"`
	Value  pgtype.JSONB `gorm:"type:jsonb;not null"`
	Zone   Zone         `gorm:"foreignKey:ZoneID;constraint:OnDelete:CASCADE"`
}
