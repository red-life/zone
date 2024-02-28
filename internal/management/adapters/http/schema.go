package http

import (
	"encoding/json"
	"github.com/red-life/zone/internal/management"
)

type ZonesResponse struct {
	Error string            `json:"error,omitempty"`
	Zones []management.Zone `json:"data,omitempty"`
}

type ZoneRequest struct {
	ZoneID string `uri:"zone_id" binding:"required" validate:"uuid4"`
}

type ZoneResponse struct {
	Error string           `json:"error,omitempty"`
	Zone  *management.Zone `json:"zone,omitempty"`
}

type CreateZoneRequest struct {
	Zone string `json:"zone" binding:"required" validate:"hostname"`
}

type CreateZoneResponse struct {
	Error string           `json:"error,omitempty"`
	Zone  *management.Zone `json:"zone,omitempty"`
}

type DeleteZoneRequest struct {
	ZoneID string `uri:"zone_id" binding:"required" validate:"uuid4"`
}

type DeleteZoneResponse struct {
	Error string `json:"error,omitempty"`
}

type RecordsRequest struct {
	ZoneID string `uri:"zone_id" binding:"required" validate:"uuid4"`
}

type RecordsResponse struct {
	Error   string              `json:"error,omitempty"`
	Records []management.Record `json:"records,omitempty"`
}

type RecordRequest struct {
	ZoneID   string `uri:"zone_id" binding:"required" validate:"uuid4"`
	RecordID string `uri:"record_id" binding:"required" validate:"uuid4"`
}

type RecordResponse struct {
	Error  string             `json:"error,omitempty"`
	Record *management.Record `json:"record,omitempty"`
}

type CreateRecordRequestURI struct {
	ZoneID string `uri:"zone_id" binding:"required" validate:"uuid4"`
}

type CreateRecordRequestBody struct {
	Name  string                `json:"name" binding:"required" validate:"ascii,min=1,max=255"`
	Type  management.RecordType `json:"type" binding:"required"`
	TTL   uint32                `json:"ttl" binding:"required"`
	Value json.RawMessage       `json:"value" binding:"required"`
}

type CreateRecordResponse struct {
	Error  string             `json:"error,omitempty"`
	Record *management.Record `json:"record,omitempty"`
}

type UpdateRecordRequestURI struct {
	ZoneID   string `uri:"zone_id" binding:"required" validate:"uuid4"`
	RecordID string `uri:"record_id" binding:"required" validate:"uuid4"`
}

type UpdateRecordRequestBody struct {
	Name  string                `json:"name" binding:"required" validate:"ascii,min=1,max=255"`
	Type  management.RecordType `json:"type" binding:"required"`
	TTL   uint32                `json:"ttl" binding:"required"`
	Value json.RawMessage       `json:"value" binding:"required"`
}

type UpdateRecordResponse struct {
	Error  string             `json:"error,omitempty"`
	Record *management.Record `json:"record,omitempty"`
}

type DeleteRecordRequest struct {
	ZoneID   string `uri:"zone_id" binding:"required" validate:"uuid4"`
	RecordID string `uri:"record_id" binding:"required" validate:"uuid4"`
}

type DeleteRecordResponse struct {
	Error string `json:"error"`
}
