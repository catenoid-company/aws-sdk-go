// +build integration

//Package sts provides gucumber integration tests support.
package sts

import (
	"github.com/catenoid-company/aws-sdk-go/awstesting/integration/smoke"
	"github.com/catenoid-company/aws-sdk-go/service/sts"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@sts", func() {
		World["client"] = sts.New(smoke.Session)
	})
}
