package handler

import "testing"

func Test_uriLenth(t *testing.T) {
	v1 := `/workload/apis/batch/v1beta1/namespaces/im-ops/cronjobs`
	if len(uriLength(v1)) != 7 {
		t.Fatal("expectd not equal")
	}
}
