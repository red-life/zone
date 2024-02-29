package resolver

import (
	"github.com/google/uuid"
	"github.com/red-life/zone/internal/shared/events"
)

func MapZoneCreatedToZone(z events.ZoneCreated) Zone {
	return Zone{
		ZoneID: z.ZoneID,
		Zone:   z.Zone,
	}
}

func MapZoneDeleted(z events.ZoneDeleted) uuid.UUID {
	return z.ZoneID
}

func MapRecordAddedToRecord(r events.RecordAdded) Record {
	// the events.RecordAdded doesn't have Zone, so the Record will be missing FullName and Zone
	return Record{
		ZoneID:     r.ZoneID,
		RecordID:   r.RecordID,
		RecordType: r.RecordType,
		Name:       r.Name,
		TTL:        r.TTL,
		Value:      r.Value,
	}
}

func MapRecordUpdatedToRecord(r events.RecordUpdated) Record {
	// the events.RecordUpdated doesn't have Zone, so the Record will be missing FullName and Zone
	return Record{
		ZoneID:     r.ZoneID,
		RecordID:   r.RecordID,
		RecordType: r.RecordType,
		Name:       r.Name,
		TTL:        r.TTL,
		Value:      r.Value,
	}
}

func MapRecordDeleted(r events.RecordDeleted) (uuid.UUID, uuid.UUID) {
	return r.ZoneID, r.RecordID
}
