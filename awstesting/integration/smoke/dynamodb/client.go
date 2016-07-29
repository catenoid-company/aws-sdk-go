// +build integration

//Package dynamodb provides gucumber integration tests support.
package dynamodb

import (
	"github.com/catenoid-company/aws-sdk-go/awstesting/integration/smoke"
	"github.com/catenoid-company/aws-sdk-go/service/dynamodb"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@dynamodb", func() {
		World["client"] = dynamodb.New(smoke.Session)
	})
}
