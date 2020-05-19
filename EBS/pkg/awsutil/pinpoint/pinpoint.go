package pinpoint

import (
	"github.com/aws/aws-sdk-go/service/pinpoint"

	"EBS/m/v2/pkg/awsutil"
)

var svc = pinpoint.New(awsutil.Session)

type Pinpoint struct {
	ApplicationId string
}

func New(ApplicationId string) *Pinpoint {
	return &Pinpoint{ApplicationId: ApplicationId}
}
