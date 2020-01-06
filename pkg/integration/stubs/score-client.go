package stubs

import (
	"fmt"
	"log"

	"github.com/kbsantiago/srs_assessment/pkg/customer/api/customerpb"
	"google.golang.org/grpc"
)

//InitializeScoreClient func
func InitializeScoreClient() {
	fmt.Println("Hello I'm a client")
	cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer cc.Close()

	c := customerpb.NewCustomerServiceClient(cc)
	fmt.Printf("Created client: %f", c)
}
