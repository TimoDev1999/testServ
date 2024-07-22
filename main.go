package main

import (
	"database/sql"
	"fmt"
	"github.com/IBM/sarama"
	_ "github.com/lib/pq"
	"testServ/handlers"
	"time"
)

func main() {

	time.Sleep(10 * time.Second)

	connStr := "host=postgres port=5432 user=postgres password=Fyfcnfcbz11 dbname=testServ sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Failed to connect to the database: %v\n", err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("Failed to ping the database: %v\n", err)
		return
	}

	fmt.Println("Successfully connected to the database")

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true

	brokers := []string{"kafka:9092"}
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		fmt.Printf("Failed to initialize Kafka producer: %v\n", err)
		return
	}

	defer func() {
		if err := producer.Close(); err != nil {
			fmt.Printf("Failed to close Kafka producer: %v\n", err)
		}
	}()

	fmt.Println("Kafka producer initialized successfully")

	router := handlers.NewRouter()
	fmt.Println("listen on", router)
}
