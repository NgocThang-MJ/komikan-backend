package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	db "github.com/liquiddev99/komikan-backend/db/sqlc"
	"github.com/liquiddev99/komikan-backend/gapi"
	"github.com/liquiddev99/komikan-backend/pb"
	"github.com/liquiddev99/komikan-backend/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load env variables")
	}

	dbpool, err := pgxpool.New(context.Background(), config.DatabaseUrl)
	if err != nil {
		log.Fatal("Failed to connect to db")
	}
	defer dbpool.Close()

	query := db.New(dbpool)

	runGrpcServer(config, query)
}

func runGrpcServer(config util.Config, query *db.Queries) {
	server, err := gapi.NewServer(config, query)
	if err != nil {
		log.Fatal("Failed to create gRPC server")
	}

	grpcServer := grpc.NewServer()

	pb.RegisterKomikanServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("Failed to create listener server")
	}

	log.Printf("Start gRPC server at %s", listener.Addr().String())
	go func() {
		grpcServer.Serve(listener)
	}()
	// End of gRPC server

	// Create Gateway server
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:9090",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	err = pb.RegisterKomikanHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}).Handler(gwmux)

	gwServer := &http.Server{
		Addr:    ":8080",
		Handler: corsHandler,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8080")
	log.Fatalln(gwServer.ListenAndServe())
}
