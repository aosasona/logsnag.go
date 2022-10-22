package logsnag

import (
	"github.com/aosasona/logsnag.go/utils"
)

var endpoint = "https://api.logsnag.com/v1"

type Logsnag struct {
	token       string
	projectName string
}

type PublishData struct {
	channel     string
	event       string
	description string
	icon        rune
	notify      bool
	tags        map[string]interface{}
}

func New(token string, projectName string) *Logsnag {
	return &Logsnag{token, projectName}
}

func (l *Logsnag) Publish(data PublishData) (map[string]interface{}, error) {
	req := utils.Request{}
	r := req.New(endpoint, []utils.Header{{"Authorization", "Bearer " + l.token}, {"Content-Type", "application/json"}})
	if data.channel == "" {
		data.channel = "default"
	}

	response, err := r.Post("/log", utils.RequestConfig{Body: data})
	if err != nil {
		return nil, err
	}

	return response, nil
}
