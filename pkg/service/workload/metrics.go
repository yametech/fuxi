package workload

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	constraint_common "github.com/yametech/fuxi/common"
	"github.com/yametech/fuxi/pkg/service/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metrics "k8s.io/metrics/pkg/apis/metrics"
)

type Metrics struct {
	client *resty.Client
}

func NewMetrics() *Metrics {
	return &Metrics{resty.New()}
}

type MetricsContentMap map[string]interface{}

func (m *Metrics) ProxyToPrometheus(params map[string]string, body []byte) (map[string]MetricsContentMap, error) {
	var bodyMap map[string]string
	var resultMap = make(map[string]MetricsContentMap)
	err := json.Unmarshal(body, &bodyMap)
	if err != nil {
		return nil, err
	}

	if constraint_common.DeployInCluster {
		for bodyKey, bodyValue := range bodyMap {
			resp, err := m.
				client.
				R().
				SetQueryParams(params).
				SetQueryParam("query", bodyValue).
				Get("http://prometheus.kube-system.svc.cluster.local/api/v1/query_range")
			if err != nil {
				return nil, err
			}
			var metricsContextMap MetricsContentMap
			err = json.Unmarshal([]byte(resp.String()), &metricsContextMap)
			if err != nil {
				return nil, err
			}
			resultMap[bodyKey] = metricsContextMap
		}
		return resultMap, nil
	}

	for bodyKey, bodyValue := range bodyMap {
		req := common.SharedK8sClient.
			ClientV1.
			CoreV1().
			RESTClient().
			Get().
			Namespace("kube-system").
			Resource("services").
			Name("prometheus:80").
			SubResource("proxy").
			Suffix("api/v1/query_range")

		for k, v := range params {
			req.Param(k, v)
		}
		req.Param("query", bodyValue)

		raw, err := req.DoRaw(context.Background())
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
		DoRaw(context.Background())
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &pods)
}

type PodMetricsList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	metav1.ListMeta `json:"metadata,omitempty"`

	// List of pod metrics.
	Items []metrics.PodMetrics `json:"items"`
}

func (m *Metrics) GetPodMetricsList(namespace string, pods *PodMetricsList) error {
	uri := "apis/metrics.k8s.io/v1beta1/pods"
	if namespace != "" {
		uri = fmt.Sprintf("apis/metrics.k8s.io/v1beta1/namespaces/%s/pods", namespace)
	}
	data, err := common.SharedK8sClient.
		ClientV1.
		RESTClient().
		Get().
		AbsPath(uri).
		DoRaw(context.Background())
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
		DoRaw(context.Background())
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &nodes)
}
