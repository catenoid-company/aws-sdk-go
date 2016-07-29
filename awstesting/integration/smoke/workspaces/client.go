// +build integration

//Package workspaces provides gucumber integration tests support.
package workspaces

import (
	"github.com/catenoid-company/aws-sdk-go/awstesting/integration/smoke"
	"github.com/catenoid-company/aws-sdk-go/service/workspaces"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@workspaces", func() {
		World["client"] = workspaces.New(smoke.Session)
	})
}
