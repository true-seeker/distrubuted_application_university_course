package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "lab3/gRPC"
	"lab3/utils/config"
	"lab3/utils/services"
	"log"
	"time"
)

func SendUnnormalizedStudent(client pb.NormalizationClient) {
	unnormalizedStudents := services.ReadSqlite("../db.db")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.SendUnnormalizedStudent(ctx)
	services.FailOnError(err, "client.RecordRoute failed")

	for _, unnormalizedStudent := range unnormalizedStudents {
		var us = pb.UnnormalizedStudent{
			Id:             int32(unnormalizedStudent.Id),
			Name:           unnormalizedStudent.Name,
			Emails:         unnormalizedStudent.Emails,
			Courses:        unnormalizedStudent.Courses,
			BirthDate:      unnormalizedStudent.BirthDate,
			Teachers:       unnormalizedStudent.Teachers,
			Faculty:        unnormalizedStudent.Faculty,
			Specialization: unnormalizedStudent.Specialization,
		}
		if err := stream.Send(&us); err != nil {
			log.Fatalf("client.RecordRoute: stream.Send(%v) failed: %v", &us, err)
		}

	}
	reply, err := stream.CloseAndRecv()
	services.FailOnError(err, "client.RecordRoute failed")

	log.Printf("Route summary: %v", reply)
}

func main() {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s",
		config.GetProperty("gRPC", "server_address"),
		config.GetProperty("gRPC", "server_port")),
		grpc.WithTransportCredentials(credentials.NewTLS(services.GetClientCerts())))
	services.FailOnError(err, "failed to dial")

	defer conn.Close()
	client := pb.NewNormalizationClient(conn)

	SendUnnormalizedStudent(client)
}
