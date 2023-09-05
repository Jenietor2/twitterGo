package secretmanager

import (
	"curso_go/twitterGo/awsgo"
	"curso_go/twitterGo/models"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(secretName string) (models.Secret, error) {
	var secret models.Secret
	fmt.Println("» Get secret " + secretName)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		return secret, err
	}

	json.Unmarshal([]byte(*key.SecretString), &secret)
	println("» Read secret OK" + secretName)
	return secret, nil
}
