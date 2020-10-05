package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"

	"tazapay.com/elearning/common/handlers"
	"tazapay.com/elearning/svc/user/controllers"
)

func main() {
	lambda.Start(handler)
}

// handler handle events invoking by AWS lambda
func handler(ctx context.Context, event interface{}) (interface{}, error) {
	return handlers.API(event, controllers.API)
}
