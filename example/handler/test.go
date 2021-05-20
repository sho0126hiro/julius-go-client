package handler

import (
	"context"
	"fmt"
	"github.com/sho0126hiro/julius-go-client/client"
)

func init() {
	client.RegisterHandler(filter, test)
}

const testMessage = "Some Message (language: ja only(maybe))"

func filter(_ context.Context, result *client.Result) bool {
	return result.Details[0].Word == testMessage
}

func test(_ context.Context, result *client.Result) error {
	fmt.Println("Some Message")
	fmt.Println("result: ", result)
	return nil
}
