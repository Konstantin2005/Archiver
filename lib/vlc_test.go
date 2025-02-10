package lib

import (
	"reflect"
	"testing"
)

func Test_prepareText(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "Base Test",
			str:  "My name is Ted",
			want: "!my name is !ted",
		},
		{
			name: "Test2",
			str:  "MDDD",
			want: "!m!d!d!d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareText(tt.str); got != tt.want {
				t.Errorf("prepareText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeBin(t *testing.T) {

	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "Base Test",
			str:  "!ted",
			want: "001000100110100101",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeBin(tt.str); got != tt.want {
				t.Errorf("encodeBin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitByChunks(t *testing.T) {
	type args struct {
		bStr      string
		ChunkSize int
	}
	tests := []struct {
		name string
		args args
		want BinaryChunks
	}{
		{
			name: "Base Test",
			args: args{
				bStr:      "001000100110100101",
				ChunkSize: 8,
			},
			want: BinaryChunks{"00100010", "01101001", "01000000"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitByChunks(tt.args.bStr, tt.args.ChunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitByChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryChunks_ToHex(t *testing.T) {
	tests := []struct {
		name string
		bcs  BinaryChunks
		want HexChunks
	}{
		{
			name: "Base Test",
			bcs:  BinaryChunks{"0101111", "10000000"},
			want: HexChunks{"2F", "80"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bcs.ToHex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "Base Test",
			args: "My name is Ted",
			want: "20 30 3C 18 77 4A E4 4D 28",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexChunks_ToString(t *testing.T) {
	tests := []struct {
		name string
		hcs  HexChunks
		want string
	}{
		{
			name: "Base Test",
			hcs:  HexChunks{"2F", "80"},
			want: "2F 80",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hcs.ToString(); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
