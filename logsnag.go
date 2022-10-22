package logsnag

import (
	"fmt"

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
	icon        string
	tags        map[string]interface{}
	notify      bool
}

type InsightData struct {
	title string
	value string
	icon  string
}

func New(token string, projectName string) *Logsnag {
	return &Logsnag{token, projectName}
}

func (l *Logsnag) CreateDefaultRequest() *utils.Request {
	req := utils.Request{}
	req.New(endpoint, []utils.Header{{"Authorization", fmt.Sprintf("Bearer %s", l.token)}, {"Content-Type", "application/json"}})
	return &req
}

func (l *Logsnag) Publish(data *PublishData) (map[string]interface{}, error) {

	r := l.CreateDefaultRequest()

	if data.channel == "" {
		data.channel = "default"
	}
	if data.icon == "" {
		data.icon = "🔔"
	}
	if data.tags == nil {
		data.tags = make(map[string]interface{})
	}

	mapData := map[string]interface{}{
		"project":     l.projectName,
		"channel":     data.channel,
		"event":       data.event,
		"description": data.description,
		"icon":        data.icon,
		"notify":      data.notify,
		"tags":        data.tags,
	}

	response, err := r.Post("/log", utils.RequestConfig{Body: mapData})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (l *Logsnag) Insight(data *InsightData) (map[string]interface{}, error) {

	r := l.CreateDefaultRequest()

	if data.icon == "" {
		data.icon = "🟠"
	}

	mapData := map[string]interface{}{
		"project": l.projectName,
		"title":   data.title,
		"value":   data.value,
		"icon":    data.icon,
	}

	response, err := r.Post("/insight", utils.RequestConfig{Body: mapData})
	if err != nil {
		return nil, err
	}

	return response, nil
}
