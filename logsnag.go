package logsnag

import (
	"fmt"

	"github.com/aosasona/logsnag.go/utils"
)

var endpoint = "https://api.logsnag.com/v1"

type Logsnag struct {
	Token       string
	ProjectName string
}

type PublishData struct {
	Project     string                 `json:"project"`
	Channel     string                 `json:"channel"`
	Event       string                 `json:"event"`
	Description string                 `json:"description,omitempty"`
	Icon        string                 `json:"icon,omitempty"`
	Tags        map[string]interface{} `json:"tags,omitempty"`
	Notify      bool                   `json:"notify"`
}

type InsightData struct {
	Project string `json:"project"`
	Title   string `json:"title"`
	Value   string `json:"value"`
	Icon    string `json:"icon,omitempty"`
}

func New(token string, projectName string) *Logsnag {
	return &Logsnag{token, projectName}
}

func (l *Logsnag) CreateDefaultRequest() *utils.Request {
	req := utils.Request{}
	req.New(endpoint, []utils.Header{{"Authorization", fmt.Sprintf("Bearer %s", l.Token)}, {"Content-Type", "application/json"}})
	return &req
}

func (l *Logsnag) Publish(data *PublishData) (map[string]interface{}, error) {

	r := l.CreateDefaultRequest()

	if data.Channel == "" {
		data.Channel = "default"
	}
	if data.Icon == "" {
		data.Icon = "🔔"
	}
	if data.Tags == nil {
		data.Tags = make(map[string]interface{})
	}

	data.Project = l.ProjectName

	mapData, err := utils.StructToMap(data)
	if err != nil {
		return nil, err
	}

	response, err := r.Post("/log", utils.RequestConfig{Body: mapData})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (l *Logsnag) Insight(data *InsightData) (map[string]interface{}, error) {

	r := l.CreateDefaultRequest()

	if data.Icon == "" {
		data.Icon = "🟠"
	}

	mapData := map[string]interface{}{
		"project": l.ProjectName,
		"title":   data.Title,
		"value":   data.Value,
		"icon":    data.Icon,
	}

	response, err := r.Post("/insight", utils.RequestConfig{Body: mapData})
	if err != nil {
		return nil, err
	}

	return response, nil
}
