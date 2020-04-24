package workload

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"
)

type Metrics struct{}

func NewMetrics() *Metrics {
	return &Metrics{}
}

func (m *Metrics) PostMetrics(arguments string) (map[string]string, error) {
	//url := fmt.Sprintf("/metrics?%s", arguments)
	data, err := sharedK8sClient.
		clientSetV1.
		RESTClient().
		Post().
		AbsPath("metrics").
		DoRaw()
	if err != nil {
		return nil, err
	}
	result := make(map[string]string)
	for _, line := range strings.Split(string(data), "\n") {
		lines := strings.Split(line, " ")
		if len(lines) < 2 || len(lines) >= 3 {
			continue
		}
		result[lines[0]] = lines[1]
	}
	return result, nil
}

func (m *Metrics) PostProxyToPrometheus(params map[string]string, body string, result io.Writer) error {
	req := sharedK8sClient.clientSetV1.CoreV1().RESTClient().Get().
		Namespace("lens-metrics").
		Resource("services").
		Name("prometheus:80").
		SubResource("proxy").
		// The server URL path, without leading "/" goes here...
		Suffix("api/v1/query_range")

	for k, v := range params {
		req.Param(k, v)
	}
	req.Body(body)

	raw, err := req.DoRaw()

	fmt.Printf("recv from proxy end %s\r\n", raw)

	if err != nil {
		panic(err.Error())
	}
	_, err = io.Copy(result, bytes.NewBuffer(raw))
	if err != nil {
		return err
	}

	return nil
}

func (m *Metrics) GetMetrics(pods *PodMetricsList) error {
	data, err := sharedK8sClient.
		clientSetV1.
		RESTClient().
		Get().
		AbsPath("apis/metrics.k8s.io/v1beta1/pods").DoRaw()
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &pods)
	return err
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
