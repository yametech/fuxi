package workload

import (
	"encoding/json"
	"fmt"
	metrics "k8s.io/metrics/pkg/apis/metrics"
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

func (m *Metrics) GetPodMetrics(namespace, name string, pods *metrics.PodMetrics) error {
	uri := fmt.Sprintf("apis/metrics.k8s.io/v1beta1/%s/%s/pods", namespace, name)
	data, err := sharedK8sClient.
		clientSetV1.
		RESTClient().
		Get().
		AbsPath(uri).
		DoRaw()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &pods)
}

func (m *Metrics) GetPodMetricsList(namespace string, pods *metrics.PodMetricsList) error {
	uri := "apis/metrics.k8s.io/v1beta1/pods"
	if namespace != "" {
		uri = fmt.Sprintf("apis/metrics.k8s.io/v1beta1/namespaces/%s/pods", namespace)
	}
	data, err := sharedK8sClient.
		clientSetV1.
		RESTClient().
		Get().
		AbsPath(uri).
		DoRaw()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &pods)
}

func (m *Metrics) GetNodeMetricsList(nodes *metrics.NodeMetricsList) error {
	data, err := sharedK8sClient.
		clientSetV1.
		RESTClient().
		Get().
		AbsPath("apis/metrics.k8s.io/v1beta1/nodes").
		DoRaw()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &nodes)
}
