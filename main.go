package main;
import ( 
	"fmt"
	"arthur/api-go/infra/kafka"
	kafka2 "arthur/api-go/aplicacao/kafka"
	"github.com/joho/godotenv"
	"log"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

)
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}
func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()
	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)
	}
}
