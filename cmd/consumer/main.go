package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/josecarlosmoura/gointensivo/internal/order/infra/database"
	"github.com/josecarlosmoura/gointensivo/internal/order/usecase"
	"github.com/josecarlosmoura/gointensivo/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"

	//sqlite3
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./orders.db")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	repository := database.NewOrderRepository(db)
	uc := usecase.CalculateFinalPriceUseCase{OrderRepository: repository}

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()
	out := make(chan amqp.Delivery)
	go rabbitmq.Consummer(ch, out)

	for msg := range out {
		var inputDTO usecase.OrderInputDTO
		err := json.Unmarshal(msg.Body, &inputDTO)
		if err != nil {
			panic(err)
		}

		outputDTo, err := uc.Execute(inputDTO)
		if err != nil {
			panic(err)
		}

		msg.Ack(false)
		fmt.Println(outputDTo)
		time.Sleep(500 * time.Millisecond)
	}
}
