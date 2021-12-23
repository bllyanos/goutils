package strings

import (
	"fmt"
	"testing"
)

func TestRandomString(t *testing.T) {
	type args struct {
		charset RandomCharset
		len     int32
	}
	Seed(0)
	tests := []struct {
		name       string
		args       args
		wantLength int
		want       string
	}{
		{name: "Case 1", args: args{AlphaCharset, 10}, want: "cubyhizzka", wantLength: 10},
		{name: "Case 2", args: args{AlphaCharset, 5}, want: "kblee", wantLength: 5},
		{name: "Case 3", args: args{AlphaCharset, 20}, want: "ansocpmignckyrwxevws", wantLength: 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomString(tt.args.charset, tt.args.len); len(got) != tt.wantLength {
				t.Errorf("RandomString() = %v, wantLength %v", len(got), tt.wantLength)
			} else if got != tt.want {
				t.Errorf("RandomString() = %v, want %v", got, tt.want)
			} else {
				fmt.Printf("%v : len = %v got => %v\n", tt.name, tt.args.len, got)
			}
		})
	}
}
