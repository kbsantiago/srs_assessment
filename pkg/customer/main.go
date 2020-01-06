package main

import (
	"fmt"

	server "github.com/kbsantiago/srs_assessment/pkg/customer/api"
	config "github.com/kbsantiago/srs_assessment/pkg/customer/config"
	repo "github.com/kbsantiago/srs_assessment/pkg/customer/repository"
)

var con = config.Config{}

func main() {
	con.Read()
	repository := repo.NewMongoRepository(con.Database, con.Server)
	i, err := repository.FindByDocument("92018068253")
	if err != nil {
		fmt.Println("Erro")
	}
	fmt.Printf("%v", i)
	server.Initialize()

}
