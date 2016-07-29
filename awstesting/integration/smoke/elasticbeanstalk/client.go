// +build integration

//Package elasticbeanstalk provides gucumber integration tests support.
package elasticbeanstalk

import (
	"github.com/catenoid-company/aws-sdk-go/awstesting/integration/smoke"
	"github.com/catenoid-company/aws-sdk-go/service/elasticbeanstalk"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@elasticbeanstalk", func() {
		World["client"] = elasticbeanstalk.New(smoke.Session)
	})
}
