package php

import (
	"testing"
)

func TestRound(t *testing.T) {
	type args struct {
		val       float64
		precision int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"test1",
			args{
				3.4,
				0,
			},
			3,
		},
		{
			"test2",
			args{
				3.5,
				0,
			},
			4,
		},
		{
			"test3",
			args{
				3.6,
				0,
			},
			4,
		},
		{
			"test4",
			args{
				1.95583,
				2,
			},
			1.96,
		},
		{
			"test5",
			args{
				1241757,
				-3,
			},
			1242000,
		},
		{
			"test6",
			args{
				5.045,
				2,
			},
			5.05,
		},
		{
			"test6",
			args{
				5.055,
				2,
			},
			5.06,
		},
		{
			"test7",
			args{
				1.55,
				1,
			},
			1.6,
		},
		{
			"test8",
			args{
				1.54,
				1,
			},
			1.5,
		},
		{
			"test9",
			args{
				-1.55,
				1,
			},
			-1.6,
		},
		{
			"test10",
			args{
				-1.54,
				1,
			},
			-1.5,
		},
		{
			"test11",
			args{
				0,
				1,
			},
			0,
		},
		{
			"test12",
			args{
				0.12345678901234567890123456789,
				15,
			},
			0.12345678901235,
		},
		{
			"test13",
			args{
				12345678901234567890,
				-15,
			},
			1.23457e+19,
		},
		{
			"test14",
			args{
				0,
				-1,
			},
			0,
		},
		{
			"test15",
			args{
				-1.55,
				-2,
			},
			0,
		},
		{
			"test16",
			args{
				-1.55,
				-1,
			},
			0,
		},
		{
			"test17",
			args{
				-1.55,
				0,
			},
			-2,
		},
		{
			"test18",
			args{
				-0.05,
				0,
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Round(tt.args.val, tt.args.precision); got != tt.want {
				t.Errorf("Round() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbs(t *testing.T) {
	if Abs(-1) != 1 {
		t.Fail()
	}
}

func TestMtRand(t *testing.T) {
	v := MtRand(10, 0)
	if v != 0 {
		t.Fail()
	}
	v = MtRand(-10, 10)
	if v < -10 || v > 10 {
		t.Fail()
	}
}

func TestBaseConvert(t *testing.T) {
	if _, err := BaseConvert("a", 10, -1); err == nil {
		t.Fail()
	}
	if v, err := BaseConvert("a37334", 16, 2); err != nil || v != "101000110111001100110100" {
		t.Fail()
	}
}

func TestBindec(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"1",
			args{"110011"},
			"51",
			false,
		},
		{
			"2",
			args{"000110011"},
			"51",
			false,
		},
		{
			"3",
			args{"111"},
			"7",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Bindec(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bindec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Bindec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecbin(t *testing.T) {
	type args struct {
		number int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{12},
			"1100",
		},
		{
			"2",
			args{26},
			"11010",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decbin(tt.args.number); got != tt.want {
				t.Errorf("Decbin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCeil(t *testing.T) {
	if Ceil(4.3) != 5 {
		t.Fail()
	}
	if Ceil(9.999) != 10 {
		t.Fail()
	}
	if Ceil(-3.14) != -3 {
		t.Fail()
	}
}

func TestFloor(t *testing.T) {
	if Floor(4.3) != 4 {
		t.Fail()
	}
	if Floor(9.999) != 9 {
		t.Fail()
	}
	if Floor(-3.14) != -4 {
		t.Fail()
	}
}

func TestPi(t *testing.T) {
	if Pi() < 3.14 || Pi() > 3.15 {
		t.Fail()
	}
}

func TestDechex(t *testing.T) {
	type args struct {
		number int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{10},
			"a",
		},
		{
			"2",
			args{47},
			"2f",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Dechex(tt.args.number); got != tt.want {
				t.Errorf("Dechex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexdec(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			"1",
			args{"ee"},
			int64(238),
			false,
		},
		{
			"2",
			args{"a0"},
			int64(160),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Hexdec(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hexdec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Hexdec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecoct(t *testing.T) {
	type args struct {
		number int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{15},
			"17",
		},
		{
			"2",
			args{264},
			"410",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decoct(tt.args.number); got != tt.want {
				t.Errorf("Decoct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOctdec(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			"1",
			args{"77"},
			int64(63),
			false,
		},
		{
			"2",
			args{"45"},
			int64(37),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Octdec(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("Octdec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Octdec() = %v, want %v", got, tt.want)
			}
		})
	}
}
