package mdlogger

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/short-d/app/fw"
)

// DataDog logging API =>
//   https://docs.datadoghq.com/api/?lang=bash#logs
const dataDogLoggingApi = "https://http-intake.logs.datadoghq.com/v1/input"

var _ EntryRepository = (*DataDogEntryRepo)(nil)

type DataDogEntryRepo struct {
	apiKey      string
	httpRequest fw.HTTPRequest
}

func (d DataDogEntryRepo) createLogEntry(level fw.LogLevel, prefix string, message string, date time.Time) {
	headers := d.authHeaders()
	body := d.requestBody(level, prefix, message)
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return
	}
	var res = make(map[string]interface{})
	_ = d.httpRequest.JSON(http.MethodPost, dataDogLoggingApi, headers, string(jsonBody), &res)
}

func getSeverity(level fw.LogLevel) string {
	switch level {
	case fw.LogFatal:
		return "critical"
	case fw.LogError:
		return "error"
	case fw.LogWarn:
		return "warning"
	case fw.LogInfo:
		return "info"
	case fw.LogDebug:
		return "debug"
	case fw.LogTrace:
		return "debug"
	default:
		return "Should not happen"
	}
}

func (d DataDogEntryRepo) requestBody(level fw.LogLevel, prefix string, message string) map[string]string {
	severity := getSeverity(level)
	return map[string]string{
		"service": prefix,
		"status":  severity,
		"message": message,
	}
}

func (d DataDogEntryRepo) authHeaders() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
		"DD-API-KEY":   d.apiKey,
	}
}

func NewDataDogEntryRepo(apiKey string, httpRequest fw.HTTPRequest) DataDogEntryRepo {
	return DataDogEntryRepo{
		apiKey:      apiKey,
		httpRequest: httpRequest,
	}
}
