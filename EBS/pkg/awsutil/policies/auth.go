package policies

import (
	"github.com/aws/aws-lambda-go/events"
)

const ALLOW = "Allow"
const DENY = "Deny"

func GenerateAuthPolicy(
	effect string,
	resource string,
) (
	policy events.APIGatewayCustomAuthorizerPolicy,
) {
	policy = events.APIGatewayCustomAuthorizerPolicy{
		Version: "2012-10-17",
		Statement: []events.IAMPolicyStatement{
			{
				Action:   []string{"execute-api:Invoke"},
				Effect:   effect,
				Resource: []string{resource},
			},
		},
	}

	return
}
