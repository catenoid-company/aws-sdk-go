// +build integration

//Package applicationdiscoveryservice provides gucumber integration tests support.
package applicationdiscoveryservice

import (
	"github.com/catenoid-company/aws-sdk-go/aws"
	"github.com/catenoid-company/aws-sdk-go/awstesting/integration/smoke"
	"github.com/catenoid-company/aws-sdk-go/service/applicationdiscoveryservice"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@applicationdiscoveryservice", func() {
		World["client"] = applicationdiscoveryservice.New(smoke.Session, &aws.Config{Region: aws.String("us-west-2")})
	})
}
