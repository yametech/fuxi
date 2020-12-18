package handler

import "strings"

/*
4
/{service}/{api_type}/{api_version}/{resource_type}
5
/{service}/{api_type}/{api_version}/{resource_type}/{resource_name}
6
/{service}/{api_type}/{api_version}/namespaces/{namespace_name}/{resource_name}
7
/{service}/{api_type}/{api_version}/namespaces/{namespace_name}/{resource_type}/{resource_name}

// parameter filter
/api/metrics?start=1608184020&end=1608187620&step=60&kubernetes_namespace=im-ops

// watch filter
/api/watch?api=/apis/tekton.dev/v1alpha1/namespaces/im-ops/tasks?watch=1&resourceVersion=139071989
*/

type APIType uint8

const (
	/*
		/{service}/{api_type}/{api_version}/{resource_type}
	*/
	FOUR APIType = iota

	/*
		/{service}/{api_type}/{api_version}/{resource_type}/{resource_name}
	*/
	FIVE

	/*
		/{service}/{api_type}/{api_version}/namespaces/{namespace_name}/{resource_name}
	*/
	SIX

	/*
		/{service}/{api_type}/{api_version}/namespaces/{namespace_name}/{resource_type}/{resource_name}
	*/
	SEVEN

	/*
		/api/watch?api=/apis/tekton.dev/v1alpha1/namespaces/im-ops/tasks?watch=1&resourceVersion=139071989
	*/
	WATCH

	/*
		/api/metrics?start=1608184020&end=1608187620&step=60&kubernetes_namespace=im-ops
	*/
	METRICS
)

type uriFilter struct{}

func (f *uriFilter) Parse(uri string) (service, resourceType, namespaceName, resourceName string, err error) {
	uriItems := uriLength(uri)
	switch len(uriItems) {
	case 4:
		service, resourceName = uriItems[0], uriItems[3]
	case 5:
		service, resourceType, resourceName = uriItems[0], uriItems[3], uriItems[4]
	case 6:
		service, namespaceName, resourceName = uriItems[0], uriItems[4], uriItems[5]
	case 7:
		service, namespaceName, resourceType, resourceName = uriItems[0], uriItems[4], uriItems[5], uriItems[6]
	}
	return
}

func (f *uriFilter) watchUri(uri string) (services []string, namespaces []string, resources []string) {
	return
}

func (f *uriFilter) metricsUri(uri string) (service string, namespace string, resource string) {
	return
}

func trimSpace(list []string) []string {
	for i, s := range list {
		if strings.TrimSpace(s) == "" {
			list = append(list[:i], list[i+1:]...)
		}
	}
	return list
}

func index(list []string, item string) int {
	for idx, _item := range list {
		if _item == item {
			return idx
		}
	}
	return -1
}

func uriLength(uri string) []string {
	return trimSpace(strings.Split(uri, "/"))
}
