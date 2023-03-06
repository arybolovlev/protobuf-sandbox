package client

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	jobv1 "github.com/arybolovlev/protobuf-sandbox/api/job"
	"github.com/arybolovlev/protobuf-sandbox/proto/job"
)

type Jobs struct {
	Job []jobv1.Job `hcl:"job,block"`
}

func Run(file string) {
	// Connect to the gRPC server.
	conn, err := grpc.Dial("localhost:8090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot connect: %v", err)
	}
	defer conn.Close()
	c := job.NewJobClient(conn)

	// Read HCL Job spec.
	var jobs Jobs
	err = DecodeFile(file, &jobs)
	if err != nil {
		log.Fatalf("Failed to load configuration file: %s", err)
	}

	// Walk through all jobs in spec and send them one by one to the server.
	for _, j := range jobs.Job {
		req := &job.JobRequest{
			Name: j.Name,
		}
		resp, err := c.Create(context.Background(), req)
		if err != nil {
			log.Fatalln("Failed to create a Job")
		}
		log.Println("New job sucessfully created:", resp)
	}
}
