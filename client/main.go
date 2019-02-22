package main

import (
	"log"

	"github.com/midoblgsm/gogrpcrestswagger-boilerplate/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := api.NewObjectClient(conn)
	//   response, err := c.SayHello(context.Background(), &api.PingMessage{Greeting: "foo"})
	response, err := c.CreateObject(context.Background(), &api.CreateObjectRequest{ObjectID: "123", ObjectDescription: "Nice object"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.GetStatus())
}
