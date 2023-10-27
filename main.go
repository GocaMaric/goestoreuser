package main

import (
	"context"
	"errors"
	"fmt"
	"goestoreuser/awsgo"
	"goestoreuser/bd"
	"goestoreuser/models"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(ExecuteLambda)

}

func ExecuteLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {

	awsgo.InitializeAWS()

	if !ValidParaments() {
		fmt.Println("Error finding lost paraments. Check ´SecretManager´")
		err := errors.New("error with finding Secret Name")
		return event, err
	}

	var data models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			data.UserEmail = att
			fmt.Print("Email =" + data.UserEmail)
		case "sub":
			data.UserUUID = att
			fmt.Println("Sub " + data.UserUUID)

		}
	}

	err := bd.ReadSecret()
	if err != nil {
		fmt.Println(" Error with Secret" + err.Error())
		return event, err
	}

}

func ValidParaments() bool {
	var getParament bool
	_, getParament = os.LookupEnv("SecretName")
	return getParament
}
