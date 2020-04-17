package mdmetrics

import (
	"encoding/json"
	"fmt"
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
	rate  metricType = "rate"
	gauge metricType = "gauge"
)

type metricPoints struct {
	Metric   string      `json:"metric"`
	Type     metricType  `json:"type"`
	Interval int         `json:"interval"`
	Points   [][]float64 `json:"points"`
	Tags     []string    `json:"tags"`
}

type timeSeries struct {
	Series []metricPoints `json:"series"`
}

type DataDog struct {
	apiKey      string
	httpRequest fw.HTTPRequest
	timer       fw.Timer
	serverEnv   fw.ServerEnv
}

func (d DataDog) Count(metricID string, point int, interval int, ctx fw.ExecutionContext) {
	d.recordPoint(metricID, count, float64(point), interval, ctx)
}

func (d DataDog) Rate(metricID string, point float32, interval int, ctx fw.ExecutionContext) {
	d.recordPoint(metricID, rate, float64(point), interval, ctx)
}

func (d DataDog) Gauge(metricID string, point float32, ctx fw.ExecutionContext) {
	d.recordPoint(metricID, gauge, float64(point), 0, ctx)
}

func (d DataDog) recordPoint(
	metricID string,
	metricType metricType,
	point float64,
	interval int,
	ctx fw.ExecutionContext,
) {
	headers := d.authHeaders()
	now := d.timer.Now()
	body := d.requestBody(metricID, metricType, point, interval, now, ctx)
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
	ctx fw.ExecutionContext,
) timeSeries {
	tags := map[string]string{
		"env":               string(d.serverEnv),
		"feature-toggle-id": ctx.FeatureToggleID,
		"experiment-bucket": ctx.ExperimentBucket,
	}
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
				Tags: d.dataDogTags(tags),
			},
		},
	}
}

func (d DataDog) dataDogTags(tags map[string]string) []string {
	var tagsList []string

	for key, val := range tags {
		pair := fmt.Sprintf("%s:%s", key, val)
		tagsList = append(tagsList, pair)
	}
	return tagsList
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
	serverEnv fw.ServerEnv,
) DataDog {
	return DataDog{
		apiKey:      apiKey,
		httpRequest: httpRequest,
		timer:       timer,
		serverEnv:   serverEnv,
	}
}
