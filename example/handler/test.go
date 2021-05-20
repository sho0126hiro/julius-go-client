package handler

import (
	"context"
	"fmt"
	"sho0126hiro/julius-go-client/client"
)

func init() {
	client.RegisterHandler(testMessage, test)
}

const testMessage = "Some Message (language: ja only(maybe))"

func test(_ context.Context, result *client.Result) error {
	fmt.Println("Some Message")
	fmt.Println("result: ", result)
	return nil
}
