package handlers

import (
	"context"
	"curso_go/twitterGo/models"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func TwitterHandlers(ctx context.Context, request events.APIGatewayProxyRequest) models.ResAPI {
	fmt.Println("Voy a procesar " + ctx.Value(models.Key("path")).(string) + " « " + ctx.Value(models.Key("method")).(string))

	var r models.ResAPI
	r.Status = 400

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {

		}
	}

	r.Message="Method Invalid"
	return r
}
