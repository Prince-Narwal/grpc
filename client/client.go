package main

import (
	"context"
	proto "grpc/protoc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := proto.NewExampleClient(conn)

	req := &proto.HelloRequest{SomeString: "hello from client"}

	client.ServerReply(context.TODO(), req)
}

// WithTransportCredentials(insecure.NewCredentials()
