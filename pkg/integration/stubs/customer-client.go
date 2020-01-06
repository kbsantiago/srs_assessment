package stubs

import (
	"context"
	"fmt"
	"log"

	"github.com/kbsantiago/srs_assessment/pkg/customer/api/customerpb"
	"google.golang.org/grpc"
)

//InitializeCustomerClient func
func InitializeCustomerClient() {

	fmt.Println("Hello I'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer cc.Close()

	c := customerpb.NewCustomerServiceClient(cc)
	customerReq := &customerpb.FindByDocumentRequest{Document: "92018068256"}
	customerResp, err := c.FindByDocument(context.Background(), customerReq)
	if err != nil {
		fmt.Println("Cannot find customer - Error")
	}

	fmt.Printf("Customer founded: %v", customerResp)
}
