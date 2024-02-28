package http

import "github.com/gin-gonic/gin"

func NewHTTPServer(api *API, engine *gin.Engine, addr string) *HTTP {
	return &HTTP{
		api:    api,
		engine: engine,
		addr:   addr,
	}
}

type HTTP struct {
	api    *API
	engine *gin.Engine
	addr   string
}

func (h *HTTP) RegisterRoutes() {
	api := h.engine.Group("/api/v1")

	api.GET("/zone", h.api.Zones)
	api.GET("/zone/:zone_id", h.api.Zone)
	api.POST("/zone", h.api.CreateZone)
	api.DELETE("/zone/:zone_id", h.api.DeleteZone)

	api.GET("/zone/:zone_id/record", h.api.Records)
	api.GET("/zone/:zone_id/record/:record_id", h.api.Record)
	api.POST("/zone/:zone_id/record", h.api.CreateRecord)
	api.PUT("/zone/:zone_id/record/:record_id", h.api.UpdateRecord)
	api.DELETE("/zone/:zone_id/record/:record_id", h.api.DeleteRecord)
}

func (h *HTTP) Run() error {
	return h.engine.Run(h.addr)
}
