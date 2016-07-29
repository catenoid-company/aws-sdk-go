// Package unit performs initialization and validation for unit tests
package unit

import (
	"github.com/catenoid-company/aws-sdk-go/aws"
	"github.com/catenoid-company/aws-sdk-go/aws/credentials"
	"github.com/catenoid-company/aws-sdk-go/aws/session"
)

// Session is a shared session for unit tests to use.
var Session = session.New(aws.NewConfig().
	WithCredentials(credentials.NewStaticCredentials("AKID", "SECRET", "SESSION")).
	WithRegion("mock-region"))
