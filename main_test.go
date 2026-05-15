package main

import (
	"testing"
	"time"
)

func TestRelativeTime(t *testing.T) {
	cases := []struct {
		name string
		d    time.Duration
		want string
	}{
		{"just now <1m", 30 * time.Second, "just now"},
		{"1 minute", 1 * time.Minute, "1 minute ago"},
		{"singular vs plural at boundary", 2 * time.Minute, "2 minutes ago"},
		{"59 minutes still in minutes", 59 * time.Minute, "59 minutes ago"},
		{"1 hour boundary", 1 * time.Hour, "1 hour ago"},
		{"23 hours", 23 * time.Hour, "23 hours ago"},
		{"1 day", 24 * time.Hour, "1 day ago"},
		// edge cases worth thinking about:
		// zero duration? negative duration (clock skew, future tasks)?
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := relativeTime(tc.d)
			if got != tc.want {
				t.Errorf("relativeTime(%v) = %q, want %q", tc.d, got, tc.want)
			}
		})
	}
}
