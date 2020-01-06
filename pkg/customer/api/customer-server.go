package api

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

//FindByDocument func
func (s *server) FindByDocument(ctx context.Context, req *customerpb.FindByDocumentRequest) (*customerpb.FindByDocumentResponse, error) {
	var result, err = customerService.FindByDocument(req.Document)
	if err != nil {
		fmt.Printf("Error... %v", err)
	}

	var bill = &customerpb.Bill{}
	var bills []*customerpb.Bill

	for i, billResult := range result.Dividas {
		bill.Data = billResult.Data.String()
		bill.Valor = billResult.Valor
		bills[i] = bill
	}

	return &customerpb.FindByDocumentResponse{
		Customer: &customerpb.Customer{
			Cpf:      result.Cpf,
			Nome:     result.Nome,
			Endereco: result.Endereco,
			Dividas:  bills,
		},
	}, nil
}

//Initialize func
func Initialize() {
	fmt.Println("API Server initialized...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	customerpb.RegisterCustomerServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
