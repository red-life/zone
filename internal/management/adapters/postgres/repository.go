package postgres

import (
	"github.com/google/uuid"
	"github.com/red-life/zone/internal/management"
	"gorm.io/gorm"
)

var _ management.Repository = (*ManagementRepository)(nil)

func NewManagementRepository(db *gorm.DB) *ManagementRepository {
	return &ManagementRepository{
		db: db,
	}
}

type ManagementRepository struct {
	db *gorm.DB
}

func (m *ManagementRepository) SaveZone(zone management.Zone) (management.Zone, error) {
	result := m.db.Create(&zone)
	return zone, management.GormToCustomError(result.Error)
}

func (m *ManagementRepository) FindZones() ([]management.Zone, error) {
	zones := make([]management.Zone, 0)
	result := m.db.Find(&zones)
	return zones, management.GormToCustomError(result.Error)
}

func (m *ManagementRepository) FindZoneByID(zoneID uuid.UUID) (management.Zone, error) {
	var zone management.Zone
	result := m.db.Where("id = ?", zoneID).First(&zone)
	return zone, management.GormToCustomError(result.Error)
}

func (m *ManagementRepository) DeleteZoneByID(zoneID uuid.UUID) error {
	result := m.db.Delete(management.Zone{
		ID: zoneID,
	})
	return management.GormToCustomError(result.Error)
}

func (m *ManagementRepository) SaveZoneRecord(record management.Record) (management.Record, error) {
	result := m.db.Create(&record)
	return record, management.GormToCustomError(result.Error)
}

func (m *ManagementRepository) FindZoneRecords(zoneID uuid.UUID) ([]management.Record, error) {
	records := make([]management.Record, 0)
	result := m.db.Where("zone_id = ?", zoneID).Find(&records)
	if result.RowsAffected <= 0 {
		return records, management.ErrNotFound
	}
	return records, management.GormToCustomError(result.Error)
}

func (m *ManagementRepository) FindZoneRecordByID(zoneID uuid.UUID, recordID uuid.UUID) (management.Record, error) {
	var record management.Record
	result := m.db.Where("id = ? AND zone_id = ?", recordID, zoneID).First(&record)
	return record, management.GormToCustomError(result.Error)
}

func (m *ManagementRepository) UpdateZoneRecordByID(zoneID uuid.UUID, recordID uuid.UUID, record management.Record) (management.Record, error) {
	updateRecord := management.Record{
		Name:  record.Name,
		TTL:   record.TTL,
		Value: record.Value,
	}
	result := m.db.Where("id = ? AND zone_id = ?", recordID, zoneID).Updates(&updateRecord)
	updateRecord.ID = recordID
	updateRecord.ZoneID = zoneID
	return updateRecord, management.GormToCustomError(result.Error)
}

func (m *ManagementRepository) DeleteZoneRecordByID(zoneID uuid.UUID, recordID uuid.UUID) error {
	result := m.db.Delete(management.Record{
		ID:     recordID,
		ZoneID: zoneID,
	})
	return management.GormToCustomError(result.Error)
}
