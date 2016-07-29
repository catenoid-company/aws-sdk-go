// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package elb

import (
	"github.com/catenoid-company/aws-sdk-go/aws"
	"github.com/catenoid-company/aws-sdk-go/aws/client"
	"github.com/catenoid-company/aws-sdk-go/aws/client/metadata"
	"github.com/catenoid-company/aws-sdk-go/aws/request"
	"github.com/catenoid-company/aws-sdk-go/aws/signer/v4"
	"github.com/catenoid-company/aws-sdk-go/private/protocol/query"
)

// Elastic Load Balancing distributes incoming traffic across your EC2 instances.
//
// For information about the features of Elastic Load Balancing, see What Is
// Elastic Load Balancing? (http://docs.aws.amazon.com/ElasticLoadBalancing/latest/DeveloperGuide/elastic-load-balancing.html)
// in the Elastic Load Balancing Developer Guide.
//
// For information about the AWS regions supported by Elastic Load Balancing,
// see Regions and Endpoints - Elastic Load Balancing (http://docs.aws.amazon.com/general/latest/gr/rande.html#elb_region)
// in the Amazon Web Services General Reference.
//
// All Elastic Load Balancing operations are idempotent, which means that they
// complete at most one time. If you repeat an operation, it succeeds with a
// 200 OK response code.
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type ELB struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "elasticloadbalancing"

// New creates a new instance of the ELB client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a ELB client from just a session.
//     svc := elb.New(mySession)
//
//     // Create a ELB client with additional configuration
//     svc := elb.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *ELB {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *ELB {
	svc := &ELB{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2012-06-01",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a ELB operation and runs any
// custom request initialization.
func (c *ELB) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
