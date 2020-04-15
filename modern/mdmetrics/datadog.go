package mdmetrics

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/short-d/app/fw"
)

// https://docs.datadoghq.com/api/??lang=bash#post-timeseries-points
const dataDogMetricsApi = "https://api.datadoghq.com/api/v1/series"

var _ fw.Metrics = (*DataDog)(nil)

type metricType string

const (
	count metricType = "count"
	rate             = "rate"
	gauge            = "gauge"
)

type metricPoints struct {
	Metric   string      `json:"metric"`
	Type     metricType  `json:"type"`
	Interval int         `json:"interval"`
	Points   [][]float64 `json:"points"`
}

type timeSeries struct {
	Series []metricPoints `json:"series"`
}

type DataDog struct {
	apiKey      string
	httpRequest fw.HTTPRequest
	timer       fw.Timer
}

func (d DataDog) Count(metricID string, point int, interval int) {
	d.recordPoint(metricID, count, float64(point), interval)
}

func (d DataDog) Rate(metricID string, point float32, interval int) {
	d.recordPoint(metricID, rate, float64(point), interval)
}

func (d DataDog) Gauge(metricID string, point float32) {
	d.recordPoint(metricID, gauge, float64(point), 0)
}

func (d DataDog) recordPoint(
	metricID string,
	metricType metricType,
	point float64,
	interval int,
) {
	headers := d.authHeaders()
	now := d.timer.Now()
	body := d.requestBody(metricID, metricType, point, interval, now)
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return
	}
	var res = make(map[string]interface{})
	_ = d.httpRequest.JSON(http.MethodPost, dataDogMetricsApi, headers, string(jsonBody), &res)
}

func (d DataDog) requestBody(
	metricID string,
	metricType metricType,
	point float64,
	interval int,
	date time.Time,
) timeSeries {
	return timeSeries{
		Series: []metricPoints{
			{
				Metric:   metricID,
				Type:     metricType,
				Interval: interval,
				Points: [][]float64{
					{
						float64(date.Unix()),
						point,
					},
				},
			},
		},
	}
}

func (d DataDog) authHeaders() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
		"DD-API-KEY":   d.apiKey,
	}
}

func NewDataDog(
	apiKey string,
	httpRequest fw.HTTPRequest,
	timer fw.Timer,
) DataDog {
	return DataDog{
		apiKey:      apiKey,
		httpRequest: httpRequest,
		timer:       timer,
	}
}
