package video

import "testing"

func Test_parseDurationToSec(t *testing.T) {
	type args struct {
		duration string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "4 days 4 hours 1 second",
			args: args{duration: "P4DT4H1S"},
			want: 360001,
		},
		{
			name: "1 hour 1 minute 1 second",
			args: args{duration: "PT1H1M1S"},
			want: 3661,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseDurationToSec(tt.args.duration); got != tt.want {
				t.Errorf("parseDurationToSec() = %v, want %v", got, tt.want)
			}
		})
	}
}
