package main

import (
	"github.com/micro/go-micro/util/log"

	"github.com/micro/go-micro"
	"github.com/mytest/micro_demo/api/user/handler"
	"github.com/mytest/micro_demo/api/user/client"

	user "github.com/mytest/micro_demo/api/user/proto/user"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.abc.api.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the User srv client
		micro.WrapHandler(client.UserWrapper(service)),
	)

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
