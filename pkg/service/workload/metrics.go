package workload

import (
	"encoding/json"
	"fmt"
	"github.com/yametech/fuxi/pkg/service/common"
	metrics "k8s.io/metrics/pkg/apis/metrics"
)

type Metrics struct{}

func NewMetrics() *Metrics {
	return &Metrics{}
}

type MetricsContentMap map[string]interface{}

func (m *Metrics) ProxyToPrometheus(params map[string]string, body []byte) (map[string]MetricsContentMap, error) {
	var bodyMap map[string]string
	var resultMap = make(map[string]MetricsContentMap)
	err := json.Unmarshal(body, &bodyMap)
	if err != nil {
		return nil, err
	}

	for bodyKey, bodyValue := range bodyMap {
		req := common.SharedK8sClient.
			ClientV1.
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
		var metricsContextMap MetricsContentMap
		err = json.Unmarshal(raw, &metricsContextMap)
		if err != nil {
			return nil, err
		}
		resultMap[bodyKey] = metricsContextMap
	}

	return resultMap, nil
}

func (m *Metrics) GetPodMetrics(namespace, name string, pods *metrics.PodMetrics) error {
	uri := fmt.Sprintf("apis/metrics.k8s.io/v1beta1/%s/%s/pods", namespace, name)
	data, err := common.SharedK8sClient.
		ClientV1.
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
	data, err := common.SharedK8sClient.
		ClientV1.
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
	data, err := common.SharedK8sClient.
		ClientV1.
		RESTClient().
		Get().
		AbsPath("apis/metrics.k8s.io/v1beta1/nodes").
		DoRaw()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &nodes)
}
