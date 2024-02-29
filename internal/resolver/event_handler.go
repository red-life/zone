package resolver

import (
	"context"
	"fmt"
	"github.com/red-life/zone/internal/shared/events"
	"github.com/red-life/zone/internal/shared/mq"
)

func NewEventHandler(mq mq.MessageQueue, repo RecordRepository, cache Cache) *EventHandler {
	handler := &EventHandler{
		mq:    mq,
		repo:  repo,
		cache: cache,
	}
	go handler.ZoneCreated()
	go handler.ZoneDeleted()
	go handler.RecordAdded()
	go handler.RecordUpdated()
	go handler.RecordDeleted()
	return handler
}

type EventHandler struct {
	mq    mq.MessageQueue
	repo  RecordRepository
	cache Cache
}

func (E *EventHandler) ZoneCreated() {
	// TODO: less code duplication
	sub := E.mq.Subscribe(context.Background(), events.ZoneCreatedEvent)
	for {
		msg, err := sub(context.Background())
		if err != nil {
			continue //TODO: add logging
		}
		zoneCreated := MustFromString[events.ZoneCreated](msg.Payload)
		zone := MapZoneCreatedToZone(zoneCreated)
		err = E.repo.SaveZone(zone)
	}
}

func (E *EventHandler) ZoneDeleted() {
	sub := E.mq.Subscribe(context.Background(), events.ZoneDeletedEvent)
	for {
		msg, err := sub(context.Background())
		if err != nil {
			continue //TODO: add logging
		}
		zoneDeleted := MustFromString[events.ZoneDeleted](msg.Payload)
		zoneID := MapZoneDeleted(zoneDeleted)
		zone, err := E.repo.FindZoneByZoneID(zoneID)
		if err != nil {
			continue
		}
		err = E.repo.DeleteZoneByID(zoneID)
		if err != nil {
			continue
		}
		cacheKey := fmt.Sprintf("record:*%s.:*", zone.Zone)
		err = E.cache.DeleteWildcard(cacheKey)
	}
}

func (E *EventHandler) RecordAdded() {
	sub := E.mq.Subscribe(context.Background(), events.RecordAddedEvent)
	for {
		msg, err := sub(context.Background())
		if err != nil {
			continue //TODO: add logging
		}
		recordAdded := MustFromString[events.RecordAdded](msg.Payload)
		record := MapRecordAddedToRecord(recordAdded)
		zone, err := E.repo.FindZoneByZoneID(record.ZoneID)
		if err != nil {
			continue
		}
		record.Zone = zone.Zone
		record.FullName = fmt.Sprintf("%s.%s.", record.Name, zone.Zone)
		err = E.repo.SaveRecord(record)
		if err != nil {
			continue
		}
		cacheKey := fmt.Sprintf("record:%s:%s", record.FullName, record.RecordType)
		cacheValue := Answer[any]{
			TTL:   record.TTL,
			Value: record.Value,
		}
		err = E.cache.Set(cacheKey, MustString(cacheValue))
	}
}

func (E *EventHandler) RecordUpdated() {
	sub := E.mq.Subscribe(context.Background(), events.RecordUpdatedEvent)
	for {
		msg, err := sub(context.Background())
		if err != nil {
			continue //TODO: add logging
		}
		recordUpdated := MustFromString[events.RecordUpdated](msg.Payload)
		record := MapRecordUpdatedToRecord(recordUpdated)
		zone, err := E.repo.FindZoneByZoneID(record.ZoneID)
		if err != nil {
			continue
		}
		record.Zone = zone.Zone
		record.FullName = fmt.Sprintf("%s.%s.", record.Name, zone.Zone)
		err = E.repo.SaveRecord(record)
		if err != nil {
			continue
		}
		cacheKey := fmt.Sprintf("record:%s:%s", record.FullName, record.RecordType)
		cacheValue := Answer[any]{
			TTL:   record.TTL,
			Value: record.Value,
		}
		err = E.cache.Set(cacheKey, MustString(cacheValue))
	}
}

func (E *EventHandler) RecordDeleted() {
	sub := E.mq.Subscribe(context.Background(), events.RecordDeletedEvent)
	for {
		msg, err := sub(context.Background())
		if err != nil {
			continue //TODO: add logging
		}
		recordDeleted := MustFromString[events.RecordDeleted](msg.Payload)
		_, recordID := MapRecordDeleted(recordDeleted)
		record, err := E.repo.FindRecordByRecordID(recordID)
		if err != nil {
			continue
		}
		err = E.repo.DeleteRecordByID(recordID)
		if err != nil {
			continue
		}
		cacheKey := fmt.Sprintf("record:%s:%s", record.FullName, record.RecordType)
		err = E.cache.Delete(cacheKey)
	}
}
