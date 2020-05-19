package secrets

import (
	"fmt"

	"EBS/m/v2/pkg/awsutil"
	"EBS/m/v2/pkg/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

var svc = secretsmanager.New(awsutil.Session)

func GetSecret(name string) (secret string, err error) {
	resp, err := svc.GetSecretValue(&secretsmanager.GetSecretValueInput{
		SecretId: &name,
	})
	if err != nil {
		return
	}
	if resp.SecretString == nil {
		err = fmt.Errorf("%s has no secret value", secret)
		return
	}
	secret = *resp.SecretString

	return
}

func CreateSecret(name string, value string) (secret *secretsmanager.CreateSecretOutput, err error) {
	secret, err = svc.CreateSecret(&secretsmanager.CreateSecretInput{
		Name:         &name,
		SecretString: aws.String(utils.ToJSON(value)),
	})

	return
}

func UpdateSecret(name string, value string) (secret *secretsmanager.UpdateSecretOutput, err error) {
	secret, err = svc.UpdateSecret(&secretsmanager.UpdateSecretInput{
		SecretId:     &name,
		SecretString: aws.String(utils.ToJSON(value)),
	})

	return
}
