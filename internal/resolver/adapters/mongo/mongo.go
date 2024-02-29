package mongo

import (
	"context"
	"github.com/google/uuid"
	"github.com/red-life/zone/internal/resolver"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ resolver.RecordRepository = (*MongoRecordRepository)(nil)

func NewMongoRecordRepository(zoneCollection *mongo.Collection, recordCollection *mongo.Collection) *MongoRecordRepository {
	return &MongoRecordRepository{
		zoneCollection:   zoneCollection,
		recordCollection: recordCollection,
	}
}

type MongoRecordRepository struct {
	zoneCollection   *mongo.Collection
	recordCollection *mongo.Collection
}

func (m *MongoRecordRepository) FindZoneByZoneID(zoneID uuid.UUID) (resolver.Zone, error) {
	result := m.zoneCollection.FindOne(context.Background(), bson.M{"zone_id": zoneID})
	if err := result.Err(); err != nil {
		return resolver.Zone{}, err
	}
	var zone resolver.Zone
	err := result.Decode(&zone)
	return zone, err
}

func (m *MongoRecordRepository) SaveZone(zone resolver.Zone) error {
	_, err := m.zoneCollection.InsertOne(context.Background(), zone)
	return err
}

func (m *MongoRecordRepository) DeleteZoneByID(zoneID uuid.UUID) error {
	_, err := m.zoneCollection.DeleteOne(context.Background(), bson.M{"zone_id": zoneID})
	if err != nil {
		return err
	}
	_, err = m.recordCollection.DeleteMany(context.Background(), bson.M{"zone_id": zoneID})
	return err
}

func (m *MongoRecordRepository) FindRecord(recordType string, zone string) (resolver.Record, error) {
	result := m.recordCollection.FindOne(context.Background(), bson.M{"record_type": recordType, "zone": zone})
	if err := result.Err(); err != nil {
		return resolver.Record{}, err
	}
	var record resolver.Record
	err := result.Decode(&record)
	return record, err

}

func (m *MongoRecordRepository) SaveRecord(record resolver.Record) error {
	_, err := m.recordCollection.InsertOne(context.Background(), record)
	return err
}

func (m *MongoRecordRepository) UpdateRecordByID(recordID uuid.UUID, record resolver.Record) error {
	_, err := m.recordCollection.UpdateOne(context.Background(), bson.M{"record_id": recordID}, record)
	return err
}

func (m *MongoRecordRepository) DeleteRecordByID(recordID uuid.UUID) error {
	_, err := m.recordCollection.DeleteOne(context.Background(), bson.M{"record_id": recordID})
	return err
}
