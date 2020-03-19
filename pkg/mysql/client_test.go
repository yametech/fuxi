package mysql

import "testing"

func TestOptions(t *testing.T) {
	opts := make([]Options, 0)
	opts = append(opts, SetOptionPassword("123"))
	opts = append(opts, SetOptionUser("test"))
	opts = append(opts, SetOptionDB("skp"))

	option := NewOption(opts...)
	if option.User != "test" && option.DBName != "skp" && option.Password != "123" {
		t.Fatal("value not equal of expected")
	}
}
