package secrets

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/joho/godotenv"
)

// LoadSecrets - first attempts to load secrets from .env file
// if no .env file is found, will attempt to load from aws
// secrets manager
func LoadSecrets(secretName, region string) map[string]string {
	var secrets map[string]string
	var err error
	// Load local .env file
	secrets, err = godotenv.Read()
	if err != nil {
		log.Fatal(".env file not found. Trying aws")
		// No .env file found, load from AWS
		return loadAWSSecrets(secretName, region)
	}
	return secrets
}

// loadAWSSecrets - load secrets from aws secrets manager
func loadAWSSecrets(secretName, region string) map[string]string {
	var secrets map[string]string

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
		panic(err)
	}

	err = json.Unmarshal([]byte(*output.SecretString), &secrets)
	if err != nil {
		panic(err)
	}

	return secrets
}
