package management

import "github.com/google/uuid"

type Repository interface {
	SaveZone(zone Zone) (Zone, error)
	FindZones() ([]Zone, error)
	FindZoneByID(zoneID uuid.UUID) (Zone, error)
	DeleteZoneByID(zoneID uuid.UUID) error
	SaveZoneRecord(record Record) (Record, error)
	FindZoneRecords(zoneID uuid.UUID) ([]Record, error)
	FindZoneRecordByID(zoneID uuid.UUID, recordID uuid.UUID) (Record, error)
	UpdateZoneRecordByID(zoneID uuid.UUID, recordID uuid.UUID, record Record) (Record, error)
	DeleteZoneRecordByID(zoneID uuid.UUID, recordID uuid.UUID) error
}
