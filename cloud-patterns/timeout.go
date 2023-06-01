/*
	Use case: when we produce a message to kafka and
*/

package main

import (
	"cloud-patterns/timeout"
	"context"
	"fmt"
	"time"
)

func mySlowRequest(input string) (string, error) {
	fmt.Println(input)
	time.Sleep(time.Minute)
	return "Hi! from slow function...", nil
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	populatedFunc := timeout.Timeout(mySlowRequest)
	res, err := populatedFunc(ctx, "This is start message of function passed from main...")

	fmt.Println(res, err)

}
