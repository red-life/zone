package management

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ Repository = (*ManagementRepository)(nil)

func NewManagementRepository(db *gorm.DB) *ManagementRepository {
	return &ManagementRepository{
		db: db,
	}
}

type ManagementRepository struct {
	db *gorm.DB
}

func (m *ManagementRepository) SaveZone(zone Zone) error {
	return m.db.Create(zone).Error
}

func (m *ManagementRepository) DeleteZoneByID(zoneID uuid.UUID) error {
	return m.db.Delete(Zone{
		ID: zoneID,
	}).Error
}

func (m *ManagementRepository) SaveRecord(record Record) error {
	return m.db.Create(record).Error
}

func (m *ManagementRepository) FindRecords() ([]Record, error) {
	records := make([]Record, 0)
	result := m.db.Find(&records)
	return records, result.Error
}

func (m *ManagementRepository) FindRecord(recordID uuid.UUID) (Record, error) {
	var record Record
	result := m.db.Find(Record{
		ID: recordID,
	}).First(&record)
	return record, result.Error
}

func (m *ManagementRepository) UpdateRecordByID(recordID uuid.UUID, record Record) error {
	return m.db.Save(Record{
		ID:    recordID,
		Name:  record.Name,
		TTL:   record.TTL,
		Value: record.Value,
	}).Error
}

func (m *ManagementRepository) DeleteRecordByID(recordID uuid.UUID) error {
	return m.db.Delete(Record{
		ID: recordID,
	}).Error
}
