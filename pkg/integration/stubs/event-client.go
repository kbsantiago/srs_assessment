package stubs

import (
	"fmt"
	"log"

	"github.com/kbsantiago/srs_assessment/pkg/customer/api/customerpb"
	"google.golang.org/grpc"
)

//InitializeEventClient func
func InitializeEventClient() {
	fmt.Println("Hello I'm a client")
	cc, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer cc.Close()

	c := customerpb.NewCustomerServiceClient(cc)
	fmt.Printf("Created client: %f", c)
}
