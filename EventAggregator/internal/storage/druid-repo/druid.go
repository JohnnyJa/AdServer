package druid

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/grafadruid/go-druid"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Repo struct {
	cfg    *Config
	client *druid.Client
}

func NewRepo(cfg *Config) *Repo {
	return &Repo{
		cfg: cfg,
	}
}

func (r *Repo) Start() error {
	cl, err := druid.NewClient(r.cfg.IngestURL)
	if err != nil {
		return err
	}

	r.client = cl
	return nil
}

func (r *Repo) WriteBatch(ctx context.Context, events []Event) error {
	if len(events) == 0 {
		return nil
	}

	eventsJson, err := json.Marshal(events)
	if err != nil {
		return err
	}

	// Формуємо ingestion payload
	payload := map[string]interface{}{
		"type": "index",
		"spec": map[string]interface{}{
			"dataSchema": map[string]interface{}{
				"dataSource": r.cfg.Datasource,
				"timestampSpec": map[string]string{
					"column": "time",
					"format": "iso",
				},
				"dimensionsSpec": map[string]interface{}{
					"dimensions": []string{
						"request_id", "event_type", "profile_id", "publisher_id", "user_id", "ip", "user_agent", "placement_id", "currency",
					},
				},
				"metricSpec": []map[string]interface{}{
					{
						"type":      "doubleSum",
						"name":      "bid_price",
						"fieldName": "bid_price",
					},
				},
				"granularitySpec": map[string]interface{}{
					"type":               "uniform",
					"segmentGranularity": "hour",
					"queryGranularity":   "minute",
					"rollup":             false,
				},
			},
			"ioConfig": map[string]interface{}{
				"type": "index",
				"inputSource": map[string]interface{}{
					"type": "inline",
					"data": strconv.Quote(string(eventsJson)),
				},
				"inputFormat": map[string]interface{}{
					"type": "json",
				},
			},
			"tuningConfig": map[string]interface{}{
				"type": "index",
			},
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal ingestion spec: %w", err)
	}

	var str string
	err = json.Unmarshal(body, &str)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, r.cfg.IngestURL, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := r.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		responseBody, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("druid ingestion failed: %s", string(responseBody))
	}

	return nil
}
