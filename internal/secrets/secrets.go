package secrets

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/go-playground/validator/v10"
)

// Loads application secrets from local FILE or AWS
func LoadSecrets() (Secrets, error) {
	secretType := os.Getenv("SECRET_TYPE")
	if (secretType == "" || secretType != "AWS") && (secretType == "" || secretType != "FILE") {
		return Secrets{}, errors.New("environment variable SECRET_TYPE must be set as either 'AWS' or 'FILE'")
	}
	secretName := os.Getenv("SECRET_NAME")
	if secretName == "" {
		return Secrets{}, errors.New("environment variable SECRET_NAME must be set as either AWS secret name or path to config FILE")
	}
	switch secretType {
	case "FILE":
		return LoadSecretsFromFile(secretName)
	case "AWS":
		return LoadSecretsFromAWS(secretName)
	default:
		return Secrets{}, errors.New("environment variable SECRET_TYPE value is unknown")
	}
}

// Load secrets from local config file
func LoadSecretsFromFile(secretName string) (Secrets, error) {
	secrets := Secrets{}

	// Read local secrets file
	configJSON, err := ioutil.ReadFile(secretName)
	if err != nil {
		return secrets, err
	}

	// Marshal secrets file into secrets struct
	err = json.Unmarshal([]byte(configJSON), &secrets)
	if err != nil {
		return secrets, err
	}
	err = validator.New().Struct(secrets)
	if err == nil {
		return secrets, nil
	}
	validationErrors := err.(validator.ValidationErrors)
	if validationErrors != nil {
		return Secrets{}, err
	}

	return secrets, nil
}

// Load secrets from AWS secrets manager
func LoadSecretsFromAWS(secretName string) (Secrets, error) {
	secrets := Secrets{}

	// Create a new AWS session
	sess := session.Must(session.NewSession())

	// Get secrets manager svc
	svc := secretsmanager.New(sess)

	// Set the name of the secret that we want to get
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	// Get the secret from AWS
	output, err := svc.GetSecretValue(input)
	if err != nil {
		return Secrets{}, err
	}

	// Marshal AWS  secrets into secrets struct
	err = json.Unmarshal([]byte(*output.SecretString), &secrets)
	if err != nil {
		return Secrets{}, err
	}

	err = validator.New().Struct(secrets)
	validationErrors := err.(validator.ValidationErrors)
	if validationErrors != nil {
		return Secrets{}, err
	}
	// values not valid, deal with errors here

	return secrets, nil
}
