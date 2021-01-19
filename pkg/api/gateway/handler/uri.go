package handler

import (
	"regexp"
	"strings"
)

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
		/{service}/api/v1/namespaces/{namespace_name}
		eg: /workload/api/v1/namespaces
	*/
	LIST_ALL_RESOURCE APIType = iota

	/*
		/{service}/api/v1/namespaces/{namespace_name}
		eg: /workload/api/v1/namespaces/im-ops
	*/
	LIST_NAMESPACE_RESOURCE

	WATCH_NAMESPACE_RESOURCE

	WATCH_ALL_NAMESPACES_RESOURCE

	/*
		/api/metrics?start=1608184020&end=1608187620&step=60&kubernetes_namespace=im-ops
	*/
	METRICS
)

type uriFilter struct{}

func (f *uriFilter) ParseQuery(uri string) (service, resourceType, namespaceName, resourceName, op string, err error) {
	uriItems := uriLength(uri)
	switch len(uriItems) {
	case 4:
		// /workload/api/v1/pods
		service, resourceName = uriItems[0], uriItems[3]
	case 5:
		// /workload/api/v1/namespaces/im-ops
		service, resourceType, namespaceName = uriItems[0], uriItems[3], uriItems[4]
	case 6:
		// /workload/api/v1/namespaces/tekton-store/pods
		service, namespaceName, resourceType = uriItems[0], uriItems[4], uriItems[5]
	case 7:
		// need to distinguish the url of api/apis

		// /workload/api/v1/namespaces/im-test/secrets/default-token-l7mgh
		matchAPI, matchErr := regexp.MatchString("/*/api", uri)
		if err != nil {
			err = matchErr
			return
		}
		if matchAPI {
			service, namespaceName, resourceType, resourceName = uriItems[0], uriItems[4], uriItems[5], uriItems[6]
		}

		matchAPIS, matchErr := regexp.MatchString("/*/apis", uri)
		if err != nil {
			err = matchErr
			return
		}
		// /workload/apis/nuwa.nip.io/v1/namespaces/im-ops/stones
		if matchAPIS {
			service, namespaceName, resourceType = uriItems[0], uriItems[5], uriItems[6]
		}

	case 8:
		// /workload/api/v1/namespaces/im/pods/sky-idmp-ui-0-b-0/log?container=main&timestamps=true&tailLines=1000&sinceTime=1970-01-01T00%3A00%3A00.000Z
		service, namespaceName, resourceType, resourceName, op = uriItems[0], uriItems[4], uriItems[5], uriItems[6], uriItems[7]
	}

	return
}

func (f *uriFilter) ParseApply(uri string) (service, resourceType, namespaceName, resourceName string, err error) {
	uriItems := uriLength(uri)
	switch len(uriItems) {
	case 4:

	case 5:
		// /workload/apis/fuxi.nip.io/v1/workload
		service, resourceType = uriItems[0], uriItems[4]
	case 6:
		// /workload/api/v1/namespaces/tekton-store/pods
		service, namespaceName, resourceType = uriItems[0], uriItems[4], uriItems[5]
	case 7:
		// /{service}/{api_type}/{api_version}/namespaces/{namespace_name}/{resource_type}/{resource_name}
		// /workload/apis/batch/v1beta1/namespaces/im-ops/cronjobs
		service, namespaceName, resourceType = uriItems[0], uriItems[5], uriItems[6]
	}
	// fmt.Printf("########--------------uri (%s) items (%v) namespaceName=%s resourceType=%s\n", uri, uriItems, namespaceName, resourceType)
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
