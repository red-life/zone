package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/red-life/zone/internal/management"
	"net/http"
)

func NewAPI(service *management.ManagementService) *API {
	return &API{
		managementService: service,
	}
}

type API struct {
	managementService *management.ManagementService
}

func (a *API) Zones(c *gin.Context) {
	var resp ZonesResponse
	zones, err := a.managementService.GetZones()
	resp.Zones = zones
	if err != nil {
		resp.Zones = nil
		resp.Error = err.Error()

	}
	c.JSON(management.CustomErrorToHTTPStatusCode(err), resp)
}

func (a *API) Zone(c *gin.Context) {
	var req ZoneRequest
	var resp ZoneResponse
	if err := c.ShouldBindUri(&req); err != nil {
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	zone, err := a.managementService.GetZone(uuid.MustParse(req.ZoneID))
	resp.Zone = &zone
	if err != nil {
		resp.Zone = nil
		resp.Error = err.Error()
	}
	c.JSON(management.CustomErrorToHTTPStatusCode(err), resp)
}

func (a *API) CreateZone(c *gin.Context) {
	var req CreateZoneRequest
	var resp CreateZoneResponse
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	zone := management.Zone{Zone: req.Zone}
	createdZone, err := a.managementService.CreateZone(zone)
	resp.Zone = &createdZone
	statusCode := http.StatusCreated
	if err != nil {
		statusCode = management.CustomErrorToHTTPStatusCode(err)
		resp.Zone = nil
		resp.Error = err.Error()
	}
	c.JSON(statusCode, resp)
}

func (a *API) DeleteZone(c *gin.Context) {
	var req DeleteZoneRequest
	var resp DeleteZoneResponse
	if err := c.ShouldBindUri(&req); err != nil {
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	err := a.managementService.DeleteZone(uuid.MustParse(req.ZoneID))
	if err != nil {
		resp.Error = err.Error()
	}
	c.JSON(management.CustomErrorToHTTPStatusCode(err), resp)
}

func (a *API) Records(c *gin.Context) {
	var req RecordsRequest
	var resp RecordsResponse
	if err := c.ShouldBindUri(&req); err != nil {
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	records, err := a.managementService.GetRecords(uuid.MustParse(req.ZoneID))
	resp.Records = records
	if err != nil {
		resp.Records = nil
		resp.Error = err.Error()
	}
	c.JSON(management.CustomErrorToHTTPStatusCode(err), resp)
}

func (a *API) Record(c *gin.Context) {
	var req RecordRequest
	var resp RecordResponse
	if err := c.ShouldBindUri(&req); err != nil {
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, err)
		return
	}
	record, err := a.managementService.GetRecord(uuid.MustParse(req.ZoneID), uuid.MustParse(req.RecordID))
	resp.Record = &record
	if err != nil {
		resp.Record = nil
		resp.Error = err.Error()
	}
	c.JSON(management.CustomErrorToHTTPStatusCode(err), resp)
}

func (a *API) CreateRecord(c *gin.Context) {
	var reqURI CreateRecordRequestURI
	var reqBody CreateRecordRequestBody
	var resp CreateRecordResponse
	if err := c.ShouldBindUri(&reqURI); err != nil {
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	if err := ValidateRecordValue(reqBody.Type, reqBody.Value); err != nil {
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	value := pgtype.JSONB{}
	_ = value.Set(reqBody.Value)
	record := management.Record{
		ZoneID: uuid.MustParse(reqURI.ZoneID),
		Name:   reqBody.Name,
		Type:   reqBody.Type,
		TTL:    reqBody.TTL,
		Value:  value,
	}
	addedRecord, err := a.managementService.AddRecord(record)
	resp.Record = &addedRecord
	statusCode := http.StatusCreated
	if err != nil {
		statusCode = management.CustomErrorToHTTPStatusCode(err)
		resp.Record = nil
		resp.Error = err.Error()
	}
	c.JSON(statusCode, resp)
}

func (a *API) UpdateRecord(c *gin.Context) {
	var reqURI UpdateRecordRequestURI
	var reqBody UpdateRecordRequestBody
	var resp UpdateRecordResponse
	if err := c.ShouldBindUri(&reqURI); err != nil {
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	if err := ValidateRecordValue(reqBody.Type, reqBody.Value); err != nil {
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	value := pgtype.JSONB{}
	_ = value.Set(reqBody.Value)
	record := management.Record{
		Name:  reqBody.Name,
		Type:  reqBody.Type,
		TTL:   reqBody.TTL,
		Value: value,
	}
	updatedRecord, err := a.managementService.UpdateRecord(uuid.MustParse(reqURI.ZoneID), uuid.MustParse(reqURI.RecordID), record)
	resp.Record = &updatedRecord
	if err != nil {
		resp.Record = nil
		resp.Error = err.Error()
	}
	c.JSON(management.CustomErrorToHTTPStatusCode(err), resp)
}

func (a *API) DeleteRecord(c *gin.Context) {
	var req DeleteRecordRequest
	var resp DeleteRecordResponse
	if err := c.ShouldBindUri(&req); err != nil {
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	err := a.managementService.DeleteRecord(uuid.MustParse(req.ZoneID), uuid.MustParse(req.RecordID))
	if err != nil {
		resp.Error = err.Error()
	}
	c.JSON(management.CustomErrorToHTTPStatusCode(err), resp)
}
