package http

import (
	"encoding/json"
	"github.com/red-life/zone/internal/management"
	"net"
)

type ARecord []struct {
	IP net.IP `json:"ip"`
}
type AAAARecord []struct {
	IP net.IP `json:"ip"`
}
type MXRecord []struct {
	Preference int    `json:"preference"`
	Domain     string `json:"domain"`
}
type NSRecord []struct {
	Domain string `json:"domain"`
}
type TXTRecord []struct {
	Text string `json:"text"`
}
type CNAMERecord struct {
	Target string `json:"target"`
}
type PTRRecord struct {
	Domain string `json:"domain"`
}

func ValidateRecordValue(recordType management.RecordType, value json.RawMessage) error {
	var err error
	switch recordType {
	case management.A:
		var record ARecord
		err = json.Unmarshal(value, &record)
	case management.AAAA:
		var record AAAARecord
		err = json.Unmarshal(value, &record)
	case management.MX:
		var record MXRecord
		err = json.Unmarshal(value, &record)
	case management.NS:
		var record NSRecord
		err = json.Unmarshal(value, &record)
	case management.TXT:
		var record TXTRecord
		err = json.Unmarshal(value, &record)
	case management.CNAME:
		var record CNAMERecord
		err = json.Unmarshal(value, &record)
	case management.PTR:
		var record PTRRecord
		err = json.Unmarshal(value, &record)
	default:
		return management.ErrValidation
	}
	if err != nil {
		return management.ErrValidation
	}
	return nil
}
