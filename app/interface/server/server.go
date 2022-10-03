package server

import (
	"fmt"
	"swagger/app/interface/container"
	"swagger/app/shared/config"
	"swagger/app/transport"

	"github.com/gin-gonic/gin"
)

func StartServer(container container.Container) {

	app := gin.Default()

	transport := transport.SetupTransport(container)
	SetUpRouter(transport, app)

	fmt.Println(app.Run(config.Server.PORTHTTP))
}
