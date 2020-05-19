package awsutil

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodbstreams"
	"github.com/aws/aws-sdk-go/service/firehose"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/redshift"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/ssm"
)

var Session *session.Session

func init() {
	Session = session.Must(session.NewSession(&aws.Config{
		S3DisableContentMD5Validation: aws.Bool(true),
		S3ForcePathStyle:              aws.Bool(true),
		EndpointResolver:              endpoints.ResolverFunc(serviceProxy),
	}))
}

func serviceProxy(
	service string,
	region string,
	optFns ...func(*endpoints.Options),
) (
	resolver endpoints.ResolvedEndpoint,
	err error,
) {
	resolver, err = endpoints.DefaultResolver().EndpointFor(service, region, optFns...)

	if os.Getenv("ENVIRONMENT") == "dev" {
		if service := serviceAddress(service); service != "" {
			resolver = endpoints.ResolvedEndpoint{URL: service}
		}
	}

	return
}

func serviceAddress(service string) string {
	host := localstackEnv("HOST")
	if host == "" {
		return ""
	}

	services := map[string]string{
		apigateway.EndpointsID:      "APIGW",
		kinesis.EndpointsID:         "KINESIS",
		dynamodb.EndpointsID:        "DYNAMO",
		dynamodbstreams.EndpointsID: "DYNAMO_STREAMS",
		s3.EndpointsID:              "S3",
		firehose.EndpointsID:        "FIREHOSE",
		lambda.EndpointsID:          "LAMBDA",
		sns.EndpointsID:             "SNS",
		sqs.EndpointsID:             "SQS",
		redshift.EndpointsID:        "REDSHIFT",
		ses.EndpointsID:             "SES",
		route53.EndpointsID:         "ROUTE53",
		cloudformation.EndpointsID:  "CLF",
		ssm.EndpointsID:             "SSM",
		secretsmanager.EndpointsID:  "SECRETS",
	}

	port := localstackEnv(services[service])
	if port == "" {
		return ""
	}

	return fmt.Sprintf("http://%s:%s", host, port)
}

func localstackEnv(v string) string {
	return os.Getenv("LOCALSTACK_" + v)
}
