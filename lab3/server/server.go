package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	pb "lab3/gRPC"
	"lab3/utils/dto"
	"lab3/utils/services"
	"net"
	"sync"
)

type normalizationServer struct {
	pb.UnimplementedNormalizationServer
	mu sync.Mutex
}

func (s *normalizationServer) SendUnnormalizedStudent(stream pb.Normalization_SendUnnormalizedStudentServer) error {
	for {
		unnormalizedStudent, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Response{
				Code:    "200",
				Message: "ok",
			})
		}
		services.FailOnError(err, "server grpc error")

		us := dto.UnnormalizedStudent{
			Id:             int(unnormalizedStudent.Id),
			Name:           unnormalizedStudent.Name,
			Emails:         unnormalizedStudent.Emails,
			Courses:        unnormalizedStudent.Courses,
			BirthDate:      unnormalizedStudent.BirthDate,
			Teachers:       unnormalizedStudent.Teachers,
			Faculty:        unnormalizedStudent.Faculty,
			Specialization: unnormalizedStudent.Specialization,
		}
		fmt.Println(us)
		services.NormalizeStudent(us)
	}
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9876))
	services.FailOnError(err, "failed to listen")

	grpcServer := grpc.NewServer(grpc.Creds(credentials.NewTLS(services.GetServerCerts())))
	pb.RegisterNormalizationServer(grpcServer, &normalizationServer{})
	fmt.Println("Listener started")
	grpcServer.Serve(lis)
}
