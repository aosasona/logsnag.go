package utils

const (
	GET = iota
	POST
)

type Request struct {
	Url     string
	Headers map[string]string
}

type Header struct {
	Key   string
	Value string
}

type Body interface{}

type RequestConfig struct {
	Body    Body
	Headers []Header
}

type SendConfig struct {
	Url    string
	method int
	Body   Body
}
