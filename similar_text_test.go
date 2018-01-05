package php

import "testing"

func TestSimilarText(t *testing.T) {
	type args struct {
		first  string
		second string
	}
	tests := []struct {
		name        string
		args        args
		wantCount   int
		wantPercent float32
	}{
		{
			"test1",
			args{
				"Hello world",
				"Hallo world",
			},
			10,
			90.90909,
		},
		{
			"test2",
			args{
				"",
				"Hello world",
			},
			0,
			0,
		},
		{
			"test3",
			args{
				"",
				"",
			},
			0,
			0,
		},
		{
			"test4",
			args{
				"",
				"",
			},
			0,
			0,
		},
		{
			"test5",
			args{
				"abc",
				"ABC",
			},
			0,
			0,
		},
		{
			"test6",
			args{
				"我爱你中国",
				"我恨你美国",
			},
			3,
			60,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCount, gotPercent := SimilarText(tt.args.first, tt.args.second)
			if gotCount != tt.wantCount {
				t.Errorf("SimilarText() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
			if gotPercent != tt.wantPercent {
				t.Errorf("SimilarText() gotPercent = %v, want %v", gotPercent, tt.wantPercent)
			}
		})
	}
}
