// Client used for testing
package main

import (
	"context"
	"flag"
	"log"
	"time"

	"final/api"

	"google.golang.org/grpc"
)

var addr = flag.String("addr", "localhost:50052", "the address to connect to")

func main() {
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() {
		if e := conn.Close(); e != nil {
			log.Printf("failed to close connection: %s", e)
		}
	}()
	c := api.NewAccountClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.CreateUser(ctx, &api.UserInfo{Name: "yuxin"})
	if err != nil {
		log.Printf("failed to create the user: %v", err)
		return
	}
	log.Printf("user created! ")
}
