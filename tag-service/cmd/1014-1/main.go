package main

import (
	"flag"
	pb "github.com/go-projects/tag-service/proto"
	"github.com/go-projects/tag-service/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

var grpcPort string
var httpPort string

func init() {
	flag.StringVar(&grpcPort, "grpc_port", "50051", "gRPC启动端口号")
	flag.StringVar(&httpPort, "http_port", "50052", "HTTP启动端口号")
	flag.Parse()
}

func main()  {
	errs := make(chan error)
	go func() {
		err := RunHttpServer(httpPort)
		if err != nil {
			errs <- err
		}
	}()
	go func() {
		err := RunGrpcServer(grpcPort)
		if err != nil {
			errs <- err
		}
	}()

	select {
	case err := <-errs:
		log.Fatalf("Run Server err: %v", err)
	}
}

func RunHttpServer(port string) error {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`pong`))
	})

	return http.ListenAndServe(":"+port, serveMux)
}

func RunGrpcServer(port string) error {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	return s.Serve(lis)
}