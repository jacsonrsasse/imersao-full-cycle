package server

import (
	"log"
	"net"

	"github.com/jacsonrsasse/imersao-full-cycle/codebank/infra/grpc/pb"
	"github.com/jacsonrsasse/imersao-full-cycle/codebank/infra/grpc/service"
	"github.com/jacsonrsasse/imersao-full-cycle/codebank/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	ProcessTransactionUseCase usecase.UseCaseTransaction
}

func NewGRPCServer() GRPCServer {
	return GRPCServer{}
}

func (g GRPCServer) Serve() {
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("could not listen tcp port")
	}

	transactionService := service.NewTransactionService()
	transactionService.ProcessTransactionUseCase = g.ProcessTransactionUseCase

	grpcServer := grpc.NewServer()
	pb.RegisterPaymentServiceServer(grpcServer, transactionService)
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}