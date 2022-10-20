package utils

import "testing"

func Test_makeEndpoint(t *testing.T) {
	type args struct {
		baseUrl string
		path    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"make endpoint",
			args{"https://example.com", "/api"},
			"https://example.com/api",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeEndpoint(tt.args.baseUrl, tt.args.path); got != tt.want {
				t.Errorf("makeEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
