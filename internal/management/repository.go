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

func (m *ManagementRepository) SaveZone(zone Zone) (Zone, error) {
	return zone, m.db.Create(&zone).Error
}

func (m *ManagementRepository) DeleteZoneByID(zoneID uuid.UUID) error {
	result := m.db.Delete(Zone{
		ID: zoneID,
	})
	return GormToCustomError(result.Error)
}

func (m *ManagementRepository) SaveZoneRecord(record Record) (Record, error) {
	result := m.db.Create(&record)
	return record, GormToCustomError(result.Error)
}

func (m *ManagementRepository) FindZoneRecords(zoneID uuid.UUID) ([]Record, error) {
	records := make([]Record, 0)
	result := m.db.Find(&records).Where("zone_id = ?", zoneID)
	return records, GormToCustomError(result.Error)
}

func (m *ManagementRepository) FindZoneRecordByID(zoneID uuid.UUID, recordID uuid.UUID) (Record, error) {
	var record Record
	result := m.db.Find(Record{
		ID: recordID,
	}).Where("zone_id = ?", zoneID).First(&record)
	return record, GormToCustomError(result.Error)
}

func (m *ManagementRepository) UpdateZoneRecordByID(zoneID uuid.UUID, recordID uuid.UUID, record Record) (Record, error) {
	saveRecord := Record{
		ID:     recordID,
		ZoneID: zoneID,
		Name:   record.Name,
		TTL:    record.TTL,
		Value:  record.Value,
	}
	result := m.db.Save(&saveRecord)
	return saveRecord, GormToCustomError(result.Error)
}

func (m *ManagementRepository) DeleteZoneRecordByID(zoneID uuid.UUID, recordID uuid.UUID) error {
	result := m.db.Delete(Record{
		ID:     recordID,
		ZoneID: zoneID,
	})
	return GormToCustomError(result.Error)
}
