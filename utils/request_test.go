package utils

import (
	"reflect"
	"testing"
)

func TestRequest_New(t *testing.T) {
	type fields struct {
		Url     string
		Headers map[string]string
	}
	type args struct {
		baseUrl string
		headers []Header
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{
		{
			"create new instance",
			fields{},
			args{"https://example.com", []Header{{"Content-Type", "application/json"}}},
			&Request{"https://example.com", map[string]string{"Content-Type": "application/json"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r Request
			if got := r.New(tt.args.baseUrl, tt.args.headers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_AppendHeader(t *testing.T) {
	type fields struct {
		Url     string
		Headers map[string]string
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{
		{"test header append",
			fields{"https://example.com", make(map[string]string)},
			args{"Content-Type", "application/json"},
			&Request{"https://example.com", map[string]string{"Content-Type": "application/json"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Request{
				Url:     tt.fields.Url,
				Headers: tt.fields.Headers,
			}
			if got := r.AppendHeader(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_Send(t *testing.T) {
	type fields struct {
		Url     string
		Headers map[string]string
	}
	type args struct {
		config SendConfig
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		{
			"send request",
			fields{"https://jsonplaceholder.typicode.com/", map[string]string{"Content-Type": "application/json"}},
			args{SendConfig{"/posts/9", GET, nil}},
			9,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Request{
				Url:     tt.fields.Url,
				Headers: tt.fields.Headers,
			}
			got, err := r.Send(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got["id"], tt.want) {
				t.Errorf("Send() got = %v, want %v", reflect.TypeOf(got["id"]), tt.want)
			}
		})
	}
}

func TestRequest_Get(t *testing.T) {
	type fields struct {
		Url     string
		Headers map[string]string
	}
	type args struct {
		endpoint string
		config   RequestConfig
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		{
			"get request",
			fields{"https://jsonplaceholder.typicode.com/", map[string]string{"Content-Type": "application/json"}},
			args{"/posts/9", RequestConfig{nil, []Header{}}},
			9,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Request{
				Url:     tt.fields.Url,
				Headers: tt.fields.Headers,
			}
			got, err := r.Get(tt.args.endpoint, tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got["id"], tt.want) {
				t.Errorf("Get() got = %v, want %v", got["id"], tt.want)
			}
		})
	}
}

func TestRequest_Post(t *testing.T) {
	type fields struct {
		Url     string
		Headers map[string]string
	}
	type args struct {
		endpoint string
		config   RequestConfig
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			"post request",
			fields{"https://jsonplaceholder.typicode.com/", map[string]string{"Content-Type": "application/json"}},
			args{"/posts", RequestConfig{map[string]interface{}{"title": "foo", "body": "bar", "userId": 1}, []Header{}}},
			map[string]interface{}{"title": "foo", "body": "bar"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Request{
				Url:     tt.fields.Url,
				Headers: tt.fields.Headers,
			}
			got, err := r.Post(tt.args.endpoint, tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("Post() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(map[string]interface{}{"title": got["title"], "body": got["body"]}, tt.want) {
				t.Errorf("Post() got = %v, want %v", got, tt.want)
			}
		})
	}
}
