package handler

import "testing"

func Test_uriLenth(t *testing.T) {
	v1 := `/workload/api/v1/namespaces`
	if len(uriLength(v1)) != 4 {
		t.Fatal("expectd not equal")
	}
}
