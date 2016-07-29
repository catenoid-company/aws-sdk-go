// +build integration

//Package directoryservice provides gucumber integration tests support.
package directoryservice

import (
	"github.com/catenoid-company/aws-sdk-go/awstesting/integration/smoke"
	"github.com/catenoid-company/aws-sdk-go/service/directoryservice"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@directoryservice", func() {
		World["client"] = directoryservice.New(smoke.Session)
	})
}
