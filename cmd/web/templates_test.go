package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2022, 01, 21, 2, 28, 0, 0, time.UTC),
			want: "21 Jan 2022 at 02:28",
		},
		{
			name: "Emplty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2022, 01, 21, 4, 22, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "21 Jan 2022 at 03:22",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			if hd != tt.want {
				t.Errorf("want %q; got %q", tt.want, hd)
			}
		})
	}

}
