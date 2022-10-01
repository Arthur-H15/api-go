package kafka

import (
	"encoding/json"
	"arthur/api-go/infra/kafka"
	route2 "arthur/api-go/aplicacao/router"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
	"time"
)

// Produce is responsible to publish the positions of each request
// Example of a json request:
//{"clientId":"1","routeId":"1"}
//{"clientId":"2","routeId":"2"}
//{"clientId":"3","routeId":"3"}
func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route2.NewRoute()
	
	// json.Unmarshal(msg.Value, &route)
	json.Unmarshal(msg.Value, &route)
	log.Println("aqui",route)
	route.LoadPositions()
	
	positions, err := route.ExportJsonPositions()
	
	if err != nil {
		log.Println(err.Error())
	}
	
	for _, p := range positions {
		
		
		kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}