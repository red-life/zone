package management

import (
	"context"
	"github.com/google/uuid"
	"github.com/red-life/zone/internal/shared/events"
	"github.com/red-life/zone/internal/shared/mq"
)

func NewManagementService(repo Repository, mq mq.MessageQueue) *ManagementService {
	return &ManagementService{
		repo: repo,
		mq:   mq,
	}
}

type ManagementService struct {
	repo Repository
	mq   mq.MessageQueue
}

func (m *ManagementService) CreateZone(zone Zone) (Zone, error) {
	zone, err := m.repo.SaveZone(zone)
	if err != nil {
		return Zone{}, err
	}
	zoneCreated := MustString(MapZoneToZoneCreated(zone))
	err = m.mq.Publish(context.Background(), events.ZoneCreatedEvent, zoneCreated)
	return zone, err
}

func (m *ManagementService) GetZones() ([]Zone, error) {
	return m.repo.FindZones()
}

func (m *ManagementService) GetZone(zoneID uuid.UUID) (Zone, error) {
	return m.repo.FindZoneByID(zoneID)
}

func (m *ManagementService) DeleteZone(zoneID uuid.UUID) error {
	err := m.repo.DeleteZoneByID(zoneID)
	if err != nil {
		return err
	}
	zoneDeleted := MustString(MapToZoneDeleted(zoneID))
	err = m.mq.Publish(context.Background(), events.ZoneDeletedEvent, zoneDeleted)
	return err
}

func (m *ManagementService) AddRecord(record Record) (Record, error) {
	record, err := m.repo.SaveZoneRecord(record)
	if err != nil {
		return Record{}, err
	}
	recordAdded := MustString(MapRecordToRecordAdded(record))
	err = m.mq.Publish(context.Background(), events.RecordAddedEvent, recordAdded)
	return record, err
}

func (m *ManagementService) GetRecords(zoneID uuid.UUID) ([]Record, error) {
	return m.repo.FindZoneRecords(zoneID)
}

func (m *ManagementService) GetRecord(zoneID uuid.UUID, recordID uuid.UUID) (Record, error) {
	return m.repo.FindZoneRecordByID(zoneID, recordID)
}

func (m *ManagementService) UpdateRecord(zoneID uuid.UUID, recordID uuid.UUID, record Record) (Record, error) {
	record, err := m.repo.UpdateZoneRecordByID(zoneID, recordID, record)
	if err != nil {
		return Record{}, err
	}
	recordUpdated := MustString(MapRecordToRecordUpdated(record))
	err = m.mq.Publish(context.Background(), events.RecordUpdatedEvent, recordUpdated)
	return record, err
}

func (m *ManagementService) DeleteRecord(zoneID uuid.UUID, recordID uuid.UUID) error {
	err := m.repo.DeleteZoneRecordByID(zoneID, recordID)
	if err != nil {
		return err
	}
	recordDeleted := MustString(MapToRecordDeleted(zoneID, recordID))
	err = m.mq.Publish(context.Background(), events.RecordDeletedEvent, recordDeleted)
	return err
}
