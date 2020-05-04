package logger

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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
	env         fw.ServerEnv
}

func (d DataDogEntryRepo) createLogEntry(
	level LogLevel,
	prefix string,
	line int,
	filename string,
	message string,
	date time.Time) {
	headers := d.authHeaders()

	body := d.requestBody(level, prefix, line, filename, message)
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return
	}
	var res = make(map[string]interface{})
	_ = d.httpRequest.JSON(http.MethodPost, dataDogLoggingApi, headers, string(jsonBody), &res)
}

// getSeverity converts internal log severity to DataDog's log status.
// Here is DataDog's documentation: https://docs.datadoghq.com/logs/processing/processors/?tab=ui#log-status-remapper
func getSeverity(level LogLevel) string {
	switch level {
	case LogFatal:
		return "critical"
	case LogError:
		return "error"
	case LogWarn:
		return "warning"
	case LogInfo:
		return "info"
	case LogDebug:
		return "debug"
	case LogTrace:
		return "debug"
	default:
		return "Should not happen"
	}
}

func (d DataDogEntryRepo) requestBody(
	level LogLevel,
	prefix string,
	line int,
	filename string,
	message string) map[string]string {
	severity := getSeverity(level)
	tags := map[string]string{
		"env":       string(d.env),
		"line":      fmt.Sprintf("%d", line),
		"file-name": filename,
	}
	return map[string]string{
		"service": prefix,
		"status":  severity,
		"message": message,
		"ddtags":  d.dataDogTags(tags),
	}
}

func (d DataDogEntryRepo) dataDogTags(tags map[string]string) string {
	var tagsList []string

	for key, val := range tags {
		pair := fmt.Sprintf("%s:%s", key, val)
		tagsList = append(tagsList, pair)
	}
	return strings.Join(tagsList, ",")
}

func (d DataDogEntryRepo) authHeaders() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
		"DD-API-KEY":   d.apiKey,
	}
}

func NewDataDogEntryRepo(apiKey string, httpRequest fw.HTTPRequest, env fw.ServerEnv) DataDogEntryRepo {
	return DataDogEntryRepo{
		apiKey:      apiKey,
		httpRequest: httpRequest,
		env:         env,
	}
}
