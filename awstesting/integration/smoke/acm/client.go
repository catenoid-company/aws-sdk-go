// +build integration

//Package acm provides gucumber integration tests support.
package acm

import (
	"github.com/catenoid-company/aws-sdk-go/awstesting/integration/smoke"
	"github.com/catenoid-company/aws-sdk-go/service/acm"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@acm", func() {
		World["client"] = acm.New(smoke.Session)
	})
}
