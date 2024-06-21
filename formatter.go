package main

import "time"

func FormatTimestamp(timestamp int64, layout string) string {
	if layout == "" {
		layout = "2006-01-02 15:04:05"
	}

	return time.Unix(timestamp, 0).Format(layout)
}
