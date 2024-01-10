package main

import (
	"learn-microservices-mail/config"
	"learn-microservices-mail/service"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading godotenv: ", err)
	}
	config.InitConfig()

	ms := service.NewMailService()

	k, err := service.NewKafkaService(ms)
	if err != nil {
		log.Fatal("Error starting kafka service: ", err)
	}

	err = k.ListenTopics()
	if err != nil {
		log.Fatal("Error listen kafka: ", err)
	}

}
