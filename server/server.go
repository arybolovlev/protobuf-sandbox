package server

import (
	"context"
	"log"
	"net"

	"github.com/google/uuid"
	"google.golang.org/grpc"

	jobv1 "github.com/arybolovlev/protobuf-sandbox/api/job"
	"github.com/arybolovlev/protobuf-sandbox/proto/job"
)

type server struct {
	job.JobServer
}

func Run() {
	ln, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalf("Falied to listen: %v", err)
	}

	srv := &server{}
	gs := grpc.NewServer()
	job.RegisterJobServer(gs, srv)

	log.Printf("server listening at %v", ln.Addr())

	if err := gs.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("Tot ziens!")
}

func (s *server) Create(ctx context.Context, req *job.JobRequest) (*job.JobResponse, error) {
	log.Println("New job request received:", req.Name)

	log.Println("Creating a new job")
	j := jobv1.Job{
		ID:   uuid.New().String(),
		Name: req.Name,
	}

	log.Println("New job was created", j)

	resp := &job.JobResponse{
		Name: j.ID,
		Id:   j.Name,
	}

	return resp, nil
}
