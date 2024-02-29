package resolver

import "github.com/google/uuid"

type Cache interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	Delete(key string) error
	DeleteWildcard(key string) error
}

type RecordRepository interface {
	FindZoneByZoneID(zoneID uuid.UUID) (Zone, error)
	SaveZone(zone Zone) error
	DeleteZoneByID(zoneID uuid.UUID) error
	FindRecord(recordType string, zone string) (Record, error)
	SaveRecord(record Record) error
	UpdateRecordByID(recordID uuid.UUID, record Record) error
	DeleteRecordByID(recordID uuid.UUID) error
}
