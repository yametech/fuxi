package workload

import (
	"encoding/json"
)

type Metrics struct{}

func NewMetrics() *Metrics {
	return &Metrics{}
}

type MetricsContext struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric struct {
			} `json:"metric"`
			Values [][]interface{} `json:"values"`
		} `json:"result"`
	} `json:"data"`
}

type MetricsReport struct {
	MemoryUsage    string `json:"memoryUsage"`
	MemoryRequests string `json:"memoryRequests"`
	MemoryLimits   string `json:"memoryLimits"`
	MemoryCapacity string `json:"memoryCapacity"`
	CPUUsage       string `json:"cpuUsage"`
	CPURequests    string `json:"cpuRequests"`
	CPULimits      string `json:"cpuLimits"`
	CPUCapacity    string `json:"cpuCapacity"`
	PodUsage       string `json:"podUsage"`
	PodCapacity    string `json:"podCapacity"`
	FsSize         string `json:"fsSize"`
	FsUsage        string `json:"fsUsage"`
}

func (m *Metrics) ProxyToPrometheus(params map[string]string, body []byte) (map[string]MetricsContext, error) {
	var bodyMap map[string]string
	var resultMap = make(map[string]MetricsContext)
	err := json.Unmarshal(body, &bodyMap)
	_ = err

	for bodyKey, bodyValue := range bodyMap {
		req := sharedK8sClient.clientSetV1.
			CoreV1().
			RESTClient().
			Get().
			Namespace("lens-metrics").
			Resource("services").
			Name("prometheus:80").
			SubResource("proxy").
			Suffix("api/v1/query_range")
		for k, v := range params {
			req.Param(k, v)
		}
		req.Param("query", bodyValue)

		raw, err := req.DoRaw()
		if err != nil {
			return nil, err
		}
		metricsContext := MetricsContext{}
		err = json.Unmarshal(raw, &metricsContext)
		if err != nil {
			panic(err)
		}
		resultMap[bodyKey] = metricsContext
	}

	return resultMap, nil
}
