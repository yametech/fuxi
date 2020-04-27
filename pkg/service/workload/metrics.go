package workload

import (
	"encoding/json"
	"time"
)

type Metrics struct{}

func NewMetrics() *Metrics {
	return &Metrics{}
}

type MetricsContent struct {
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

func (m *Metrics) ProxyToPrometheus(params map[string]string, body []byte) (map[string]MetricsContent, error) {
	var bodyMap map[string]string
	var resultMap = make(map[string]MetricsContent)
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
		metricsContext := MetricsContent{}
		err = json.Unmarshal(raw, &metricsContext)
		if err != nil {
			panic(err)
		}
		resultMap[bodyKey] = metricsContext
	}

	return resultMap, nil
}

func (m *Metrics) GetMetrics(pods *PodMetricsList) error {
	data, err := sharedK8sClient.
		clientSetV1.
		RESTClient().
		Get().
		AbsPath("apis/metrics.k8s.io/v1beta1/pods").
		DoRaw()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &pods)
}

type PodMetricsList struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata   struct {
		SelfLink string `json:"selfLink"`
	} `json:"metadata"`
	Items []struct {
		Metadata struct {
			Name              string    `json:"name"`
			Namespace         string    `json:"namespace"`
			SelfLink          string    `json:"selfLink"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
		} `json:"metadata"`
		Timestamp  time.Time `json:"timestamp"`
		Window     string    `json:"window"`
		Containers []struct {
			Name  string `json:"name"`
			Usage struct {
				CPU    string `json:"cpu"`
				Memory string `json:"memory"`
			} `json:"usage"`
		} `json:"containers"`
	} `json:"items"`
}
