package main

import (
	"database/sql"
	"net"

	"github.com/isslerman/goexpert/14-GRPC/internal/database"
	"github.com/isslerman/goexpert/14-GRPC/internal/pb"
	"github.com/isslerman/goexpert/14-GRPC/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()
	// registrar o nosso serviço ao nosso servidor gRPC
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	// agora abrimos uma conexão TCP para falar com o gRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}
