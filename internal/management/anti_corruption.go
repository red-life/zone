package management

import (
	"github.com/google/uuid"
	"github.com/red-life/zone/internal/shared/events"
)

func MapZoneToZoneCreated(zone Zone) events.ZoneCreated {
	return events.ZoneCreated{
		ZoneID: zone.ID,
		Zone:   zone.Zone,
	}
}

func MapToZoneDeleted(zoneID uuid.UUID) events.ZoneDeleted {
	return events.ZoneDeleted{
		ZoneID: zoneID,
	}
}

func MapRecordToRecordAdded(record Record) events.RecordAdded {
	b, _ := record.Value.MarshalJSON()
	return events.RecordAdded{
		ZoneID:   record.ZoneID,
		RecordID: record.ID,
		Name:     record.Name,
		TTL:      record.TTL,
		Value:    b,
	}
}

func MapRecordToRecordUpdated(record Record) events.RecordUpdated {
	b, _ := record.Value.MarshalJSON()
	return events.RecordUpdated{
		ZoneID:   record.ZoneID,
		RecordID: record.ID,
		Name:     record.Name,
		TTL:      record.TTL,
		Value:    b,
	}
}

func MapToRecordDeleted(zoneID uuid.UUID, recordID uuid.UUID) events.RecordDeleted {
	return events.RecordDeleted{
		ZoneID:   zoneID,
		RecordID: recordID,
	}
}
