package client

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/arybolovlev/protobuf-sandbox/proto/job"
)

func Run() {
	conn, err := grpc.Dial("localhost:8090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot connect: %v", err)
	}
	defer conn.Close()
	c := job.NewJobClient(conn)

	req := &job.JobRequest{
		Name: "name",
	}
	j, err := c.Create(context.Background(), req)
	if err != nil {
		log.Fatalln("Failed to create a Job")
	}
	log.Println("New job sucessfully created:", j)
}
