// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package codecommit

import (
	"github.com/catenoid-company/aws-sdk-go/aws"
	"github.com/catenoid-company/aws-sdk-go/aws/client"
	"github.com/catenoid-company/aws-sdk-go/aws/client/metadata"
	"github.com/catenoid-company/aws-sdk-go/aws/request"
	"github.com/catenoid-company/aws-sdk-go/aws/signer/v4"
	"github.com/catenoid-company/aws-sdk-go/private/protocol/jsonrpc"
)

// This is the AWS CodeCommit API Reference. This reference provides descriptions
// of the operations and data types for AWS CodeCommit API.
//
// You can use the AWS CodeCommit API to work with the following objects:
//
//  Repositories, by calling the following:  BatchGetRepositories, which returns
// information about one or more repositories associated with your AWS account
// CreateRepository, which creates an AWS CodeCommit repository DeleteRepository,
// which deletes an AWS CodeCommit repository GetRepository, which returns information
// about a specified repository ListRepositories, which lists all AWS CodeCommit
// repositories associated with your AWS account UpdateRepositoryDescription,
// which sets or updates the description of the repository UpdateRepositoryName,
// which changes the name of the repository. If you change the name of a repository,
// no other users of that repository will be able to access it until you send
// them the new HTTPS or SSH URL to use.  Branches, by calling the following:
//  CreateBranch, which creates a new branch in a specified repository GetBranch,
// which returns information about a specified branch ListBranches, which lists
// all branches for a specified repository UpdateDefaultBranch, which changes
// the default branch for a repository  Information about committed code in
// a repository, by calling the following:  GetCommit, which returns information
// about a commit, including commit messages and committer information.  Triggers,
// by calling the following:  GetRepositoryTriggers, which returns information
// about triggers configured for a repository PutRepositoryTriggers, which replaces
// all triggers for a repository and can be used to create or delete triggers
// TestRepositoryTriggers, which tests the functionality of a repository trigger
// by sending data to the trigger target    For information about how to use
// AWS CodeCommit, see the AWS CodeCommit User Guide (http://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html).
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type CodeCommit struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "codecommit"

// New creates a new instance of the CodeCommit client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a CodeCommit client from just a session.
//     svc := codecommit.New(mySession)
//
//     // Create a CodeCommit client with additional configuration
//     svc := codecommit.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *CodeCommit {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *CodeCommit {
	svc := &CodeCommit{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2015-04-13",
				JSONVersion:   "1.1",
				TargetPrefix:  "CodeCommit_20150413",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a CodeCommit operation and runs any
// custom request initialization.
func (c *CodeCommit) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
