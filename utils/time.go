package utils

import (
	"time"
)

func Epoch_to_iso(epoch_time int) string {
	t := time.Unix(int64(epoch_time), 0)
	return t.Local().String()
}

func Iso_to_epoch(iso_time string) (int, error) {
	t, err := time.Parse("2006-01-02 15:04:05 -0700", iso_time)
	if err != nil {
		return 0, err
	}
	return int(t.Unix()), nil
}
