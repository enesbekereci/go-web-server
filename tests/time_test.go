package tests

import (
	"go-web-server/utils"
	"testing"
)

func Test1(t *testing.T) {
	epoch_time := 5
	want := "1970-01-01 03:00:05 +0300 +03"
	result := utils.Epoch_to_iso(epoch_time)
	if want != result {
		t.Fatalf(`Not Equal %q %q`, result, want)
	}
}

func Test2(t *testing.T) {
	want := 5
	iso_time := "1970-01-01 03:00:05 +0300"
	result, _ := utils.Iso_to_epoch(iso_time)
	if want != result {
		t.Fatalf(`Not Equal %d %d`, result, want)
	}
}
