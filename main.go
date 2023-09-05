package main

import (
	"context"
	"curso_go/twitterGo/awsgo"
	"curso_go/twitterGo/bd"
	"curso_go/twitterGo/handlers"
	"curso_go/twitterGo/models"
	secretmanager "curso_go/twitterGo/secret_manager"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse

	awsgo.StartAWS()

	if !ValidateParameters() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers:    map[string]string{"Content-type": "application/json"},
			Body:       "Error en las variabpes de entorno. Deben incluir 'SecretName', 'BucketName', 'UrlPrefix'",
		}
	}

	secrectModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers:    map[string]string{"Content-type": "application/json"},
			Body:       "Error en la lectura del secret " + err.Error(),
		}

		return res, nil
	}

	path := strings.Replace(request.PathParameters["twitterGo"], os.Getenv("UrlPrefix"), "", -1)

	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), secrectModel.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), secrectModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), secrectModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), secrectModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtSing"), secrectModel.JWTSing)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucktName"))

	//Check database connection
	err = bd.ConecctionDB(awsgo.Ctx)
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers:    map[string]string{"Content-type": "application/json"},
			Body:       "Error conectando a la base de datos" + err.Error(),
		}

		return res, nil
	}

	resAPI := handlers.TwitterHandlers(awsgo.Ctx, request)
	if resAPI.CustomRespon == nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: resAPI.Status,
			Headers:    map[string]string{"Content-type": "application/json"},
			Body:       string(resAPI.Message),
		}

		return res, nil
	} else {
		return resAPI.CustomRespon, nil
	}
}

func ValidateParameters() bool {
	_, bringParameter := os.LookupEnv("SecretName")
	if !bringParameter {
		return bringParameter
	}

	_, bringParameter = os.LookupEnv("BucketName")
	if !bringParameter {
		return bringParameter
	}

	_, bringParameter = os.LookupEnv("UrlPrefix")
	if !bringParameter {
		return bringParameter
	}

	return bringParameter
}
