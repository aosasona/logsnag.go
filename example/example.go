package main

import (
	"log"

	"github.com/aosasona/logsnag.go"
)

func main() {
	l := logsnag.New("YOUR_TOKEN", "YOUR_PROJECT_NAME")

	// Publish an event
	data, err := l.Publish(&logsnag.PublishData{Channel: "test", Event: "test-event", Description: "Hello world!", Icon: "🎉", Notify: true})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Published : %v", data)

	// Create an insight
	data, err = l.Insight(&logsnag.InsightData{Title: "test-insight", Value: "Hello world!", Icon: "🟢"})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Insight : %v", data)
}
