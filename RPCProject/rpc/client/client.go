package main

import (
	"context"
	"gomod/rpc/server/proto"

	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:7000",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := proto.NewOrderServiceClient(conn)
	defer conn.Close()


	// Contact the server and print out its response.

	ctx,_ := context.WithTimeout(context.Background(), time.Second)

	req := &proto.QueryId{Id: 5}
	order, err :=c.SelectById(ctx,req)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(order.Id)
	fmt.Println(order.FileUrl)
	fmt.Println(order.Status)
	fmt.Println(order.Amount)
	fmt.Println(order.UserName)
	fmt.Println(order.OrderNo)

}
