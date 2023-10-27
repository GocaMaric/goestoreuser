package bd

import (
	"goestoreuser/models"
	"goestoreuser/secretm"
	"os"
)

var SecretModel models.SecretRDSJson
var err error

func ReadSecret() error {
	SecretModel, err := secretm.GetSecret(os.Getenv("SecretName"))
}
