package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"

	"MeliMutant/cmd/api/di"
)

func main() {
	handler, err := di.Initialize()
	if err != nil {
		log.Println(err.Error())

		panic("fatal err: " + err.Error())
	}
	lambda.Start(handler.Handle)
}
