package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/kbsantiago/srs_assessment/pkg/customer/api/customerpb"

	config "github.com/kbsantiago/srs_assessment/pkg/customer/config"
	repo "github.com/kbsantiago/srs_assessment/pkg/customer/repository"
	service "github.com/kbsantiago/srs_assessment/pkg/customer/usecase"
)

var con = config.Config{}
var customerRepo = repo.NewMongoRepository(con.Database, con.Server)
var customerService = service.NewService(customerRepo)

//Server struct
type server struct{}

func (s *server) CreateCustomer(ctx context.Context, req *customerpb.CreateCustomerRequest) (*customerpb.CreateCustomerResponse, error) {
	// customer := req.GetCustomer()

	// data := entity.Customer{
	// 	Nome:     "Teste",
	// 	Cpf:      "92018068256",
	// 	Endereco: "Rua Curt Hering, 80",
	// 	Dividas: []entity.Bill{
	// 		{Valor: 1000, Data: time.Now()},
	// 		{Valor: 2000, Data: time.Now()},
	// 	},
	// }

	//res := customerService.Create(data)
	return nil, nil
}

func (s *server) FindByDocument(ctx context.Context, req *customerpb.FindByDocumentRequest) (*customerpb.FindByDocumentResponse, error) {
	// var result, err = customerService.FindByDocument(req.Document)

	// return &customerpb.CreateCustomerResponse {
	// 	Customer: &customerpb.Customer{
	// 		Id:       oid.Hex(),
	// 		Cpf:      result.GetCpf(),
	// 		Nome:     result.GetNome(),
	// 		Endereco: result.GetEndereco(),
	// 		Dividas:
	// 	},
	// }, nil

	return nil, nil
}

func main() {
	fmt.Println("API Server initialized...")

	lis, err := net.Listen("tcp", "0.0.0.0:50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	customerpb.RegisterCustomerServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
