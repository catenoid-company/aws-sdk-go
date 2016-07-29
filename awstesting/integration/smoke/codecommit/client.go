// +build integration

//Package codecommit provides gucumber integration tests support.
package codecommit

import (
	"github.com/catenoid-company/aws-sdk-go/awstesting/integration/smoke"
	"github.com/catenoid-company/aws-sdk-go/service/codecommit"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@codecommit", func() {
		World["client"] = codecommit.New(smoke.Session)
	})
}
