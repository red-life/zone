package management

import "github.com/google/uuid"

func NewManagementService(repo Repository) *ManagementService {
	return &ManagementService{
		repo: repo,
	}
}

type ManagementService struct {
	repo Repository
}

func (m *ManagementService) CreateZone(zone Zone) (Zone, error) {
	return m.repo.SaveZone(zone)
}

func (m *ManagementService) DeleteZone(zoneID uuid.UUID) error {
	return m.repo.DeleteZoneByID(zoneID)
}

func (m *ManagementService) AddRecord(record Record) (Record, error) {
	return m.repo.SaveZoneRecord(record)
}

func (m *ManagementService) GetRecords(zoneID uuid.UUID) ([]Record, error) {
	return m.repo.FindZoneRecords(zoneID)
}

func (m *ManagementService) GetRecord(zoneID uuid.UUID, recordID uuid.UUID) (Record, error) {
	return m.repo.FindZoneRecordByID(zoneID, recordID)
}

func (m *ManagementService) UpdateRecord(zoneID uuid.UUID, recordID uuid.UUID, record Record) (Record, error) {
	return m.repo.UpdateZoneRecordByID(zoneID, recordID, record)
}

func (m *ManagementService) DeleteRecord(zoneID uuid.UUID, recordID uuid.UUID) error {
	return m.repo.DeleteZoneRecordByID(zoneID, recordID)
}
