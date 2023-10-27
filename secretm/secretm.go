package secretm

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"goestoreuser/awsgo"
	"goestoreuser/models"
)

func GetSecret(numberSecret string) (models.SecretRDSJson, error) {
	var dataSecret models.SecretRDSJson
	fmt.Println(" > Secret number " + numberSecret)

	svc := secretmanager.NewFromConfig(awsgo.Cfg)
	clue, err := svc.GetSecretValue(awsgo.Ctx, &secretmanager.GetSecretValueInput{
		SecretId: aws.String(numberSecret),
	})
	if err != nil {
		fmt.Println(err.Error())
		return dataSecret, err
	}

	json.Unmarshal([]byte(*clue.SecretString), &dataSecret)
	fmt.Println(" > Secret OK" + numberSecret)
	return dataSecret, nill
}
