package secrets

import (
	"encoding/json"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

// Secrets - Application secrets
type Secrets struct {
	DBHost     string `json:"db_host"`
	DBPort     string `json:"db_port"`
	DBUser     string `json:"db_user"`
	DBPass     string `json:"db_pass"`
	DBDatabase string `json:"db_database"`
	JWTSecret  string `json:"jwt_secret"`
}

// LoadSecrets - first attempts to load secrets from secrets.json file
// if no secrets.json file is found, will attempt to load from aws
// secrets manager
func LoadSecrets(secretName, region string) (Secrets, error) {
	secrets := Secrets{}
	configJSON, err := ioutil.ReadFile("configs/secrets.json")
	if err != nil {
		return secrets, err
	}
	err = json.Unmarshal([]byte(configJSON), &secrets)
	if err != nil {
		// No local config file found, load from AWS
		return loadAWSSecrets(secretName, region), nil
	}
	return secrets, nil
}

// loadAWSSecrets - load secrets from aws secrets manager
func loadAWSSecrets(secretName, region string) Secrets {
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
		panic(err)
	}

	err = json.Unmarshal([]byte(*output.SecretString), &secrets)
	if err != nil {
		panic(err)
	}

	return secrets
}
