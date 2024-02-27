package management

import "github.com/google/uuid"

type Repository interface {
	SaveZone(zone Zone) error
	DeleteZoneByID(zoneID uuid.UUID) error
	SaveRecord(record Record) error
	FindRecords() ([]Record, error)
	FindRecord(recordID uuid.UUID) (Record, error)
	UpdateRecordByID(recordID uuid.UUID, record Record) error
	DeleteRecordByID(recordID uuid.UUID) error
}
