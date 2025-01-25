package lib

import "testing"

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
