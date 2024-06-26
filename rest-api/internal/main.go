package main

import (
	"log"

	"go_rest_api/package/swagger/server/restapi"

	"go_rest_api/package/swagger/server/restapi/operations"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
)

func main() {

	// Initialize Swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewHelloAPIAPI(swaggerSpec)
	server := restapi.NewServer(api)

	defer func() {
		if err := server.Shutdown(); err != nil {
			// error handle
			log.Fatalln(err)
		}
	}()

	server.Port = 8080

	api.CheckHealthHandler = operations.CheckHealthHandlerFunc(Health)

	api.GetLetsGoTeamHandler = operations.GetLetsGoTeamHandlerFunc(GetHelloUser)

	// Start server which listening
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

// Health route returns OK
func Health(operations.CheckHealthParams) middleware.Responder {
	return operations.NewCheckHealthOK().WithPayload("OK")
}

// GetHelloUser returns Hello + your name
func GetHelloUser(team operations.GetLetsGoTeamParams) middleware.Responder {
	return operations.NewGetLetsGoTeamOK().WithPayload("Let's Go " + team.Team + "!")
}
