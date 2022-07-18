package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"MeliMutant/cmd/api/di"
)

func main() {
	handler, err := di.Initialize()
	if err != nil {
		panic("fatal err: " + err.Error())
	}
	lambda.Start(handler.Handle)
}
