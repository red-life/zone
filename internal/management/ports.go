package management

import "github.com/google/uuid"

type Repository interface {
	SaveZone(zone Zone) error
	DeleteZoneByID(zoneID uuid.UUID) error
	SaveZoneRecord(record Record) error
	FindZoneRecords(zoneID uuid.UUID) ([]Record, error)
	FindZoneRecordByID(zoneID uuid.UUID, recordID uuid.UUID) (Record, error)
	UpdateZoneRecordByID(zoneID uuid.UUID, recordID uuid.UUID, record Record) error
	DeleteZoneRecordByID(zoneID uuid.UUID, recordID uuid.UUID) error
}
