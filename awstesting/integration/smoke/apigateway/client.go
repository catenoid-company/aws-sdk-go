// +build integration

//Package apigateway provides gucumber integration tests support.
package apigateway

import (
	"github.com/catenoid-company/aws-sdk-go/awstesting/integration/smoke"
	"github.com/catenoid-company/aws-sdk-go/service/apigateway"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@apigateway", func() {
		World["client"] = apigateway.New(smoke.Session)
	})
}
