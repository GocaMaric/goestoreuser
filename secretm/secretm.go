package secretm

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"goestoreuser/awsgo"
	"goestoreuser/models"
)

func GetSecret(numberSecret string) (models.SecretRDSJson, error) {
	var dataSecret models.SecretRDSJson
	fmt.Println(" > Secret number " + numberSecret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clue, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(numberSecret),
	})
	if err != nil {
		fmt.Println(err.Error())
		return dataSecret, err
	}

	json.Unmarshal([]byte(*clue.SecretString), &dataSecret)
	fmt.Println(" > Secret OK" + numberSecret)
	return dataSecret, nil
}
