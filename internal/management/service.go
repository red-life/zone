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

func (m *ManagementService) CreateZone(zone Zone) error {
	return m.repo.SaveZone(zone)
}

func (m *ManagementService) DeleteZone(zoneID uuid.UUID) error {
	return m.repo.DeleteZoneByID(zoneID)
}

func (m *ManagementService) AddRecord(record Record) error {
	return m.repo.SaveRecord(record)
}

func (m *ManagementService) GetRecords() ([]Record, error) {
	return m.repo.FindRecords()
}

func (m *ManagementService) GetRecord(recordID uuid.UUID) (Record, error) {
	return m.repo.FindRecordByID(recordID)
}

func (m *ManagementService) UpdateRecord(recordID uuid.UUID, record Record) error {
	return m.repo.UpdateRecordByID(recordID, record)
}

func (m *ManagementService) DeleteRecord(recordID uuid.UUID) error {
	return m.repo.DeleteRecordByID(recordID)
}
