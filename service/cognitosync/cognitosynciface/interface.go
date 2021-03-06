// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cognitosynciface provides an interface for the Amazon Cognito Sync.
package cognitosynciface

import (
	"github.com/catenoid-company/aws-sdk-go/aws/request"
	"github.com/catenoid-company/aws-sdk-go/service/cognitosync"
)

// CognitoSyncAPI is the interface type for cognitosync.CognitoSync.
type CognitoSyncAPI interface {
	BulkPublishRequest(*cognitosync.BulkPublishInput) (*request.Request, *cognitosync.BulkPublishOutput)

	BulkPublish(*cognitosync.BulkPublishInput) (*cognitosync.BulkPublishOutput, error)

	DeleteDatasetRequest(*cognitosync.DeleteDatasetInput) (*request.Request, *cognitosync.DeleteDatasetOutput)

	DeleteDataset(*cognitosync.DeleteDatasetInput) (*cognitosync.DeleteDatasetOutput, error)

	DescribeDatasetRequest(*cognitosync.DescribeDatasetInput) (*request.Request, *cognitosync.DescribeDatasetOutput)

	DescribeDataset(*cognitosync.DescribeDatasetInput) (*cognitosync.DescribeDatasetOutput, error)

	DescribeIdentityPoolUsageRequest(*cognitosync.DescribeIdentityPoolUsageInput) (*request.Request, *cognitosync.DescribeIdentityPoolUsageOutput)

	DescribeIdentityPoolUsage(*cognitosync.DescribeIdentityPoolUsageInput) (*cognitosync.DescribeIdentityPoolUsageOutput, error)

	DescribeIdentityUsageRequest(*cognitosync.DescribeIdentityUsageInput) (*request.Request, *cognitosync.DescribeIdentityUsageOutput)

	DescribeIdentityUsage(*cognitosync.DescribeIdentityUsageInput) (*cognitosync.DescribeIdentityUsageOutput, error)

	GetBulkPublishDetailsRequest(*cognitosync.GetBulkPublishDetailsInput) (*request.Request, *cognitosync.GetBulkPublishDetailsOutput)

	GetBulkPublishDetails(*cognitosync.GetBulkPublishDetailsInput) (*cognitosync.GetBulkPublishDetailsOutput, error)

	GetCognitoEventsRequest(*cognitosync.GetCognitoEventsInput) (*request.Request, *cognitosync.GetCognitoEventsOutput)

	GetCognitoEvents(*cognitosync.GetCognitoEventsInput) (*cognitosync.GetCognitoEventsOutput, error)

	GetIdentityPoolConfigurationRequest(*cognitosync.GetIdentityPoolConfigurationInput) (*request.Request, *cognitosync.GetIdentityPoolConfigurationOutput)

	GetIdentityPoolConfiguration(*cognitosync.GetIdentityPoolConfigurationInput) (*cognitosync.GetIdentityPoolConfigurationOutput, error)

	ListDatasetsRequest(*cognitosync.ListDatasetsInput) (*request.Request, *cognitosync.ListDatasetsOutput)

	ListDatasets(*cognitosync.ListDatasetsInput) (*cognitosync.ListDatasetsOutput, error)

	ListIdentityPoolUsageRequest(*cognitosync.ListIdentityPoolUsageInput) (*request.Request, *cognitosync.ListIdentityPoolUsageOutput)

	ListIdentityPoolUsage(*cognitosync.ListIdentityPoolUsageInput) (*cognitosync.ListIdentityPoolUsageOutput, error)

	ListRecordsRequest(*cognitosync.ListRecordsInput) (*request.Request, *cognitosync.ListRecordsOutput)

	ListRecords(*cognitosync.ListRecordsInput) (*cognitosync.ListRecordsOutput, error)

	RegisterDeviceRequest(*cognitosync.RegisterDeviceInput) (*request.Request, *cognitosync.RegisterDeviceOutput)

	RegisterDevice(*cognitosync.RegisterDeviceInput) (*cognitosync.RegisterDeviceOutput, error)

	SetCognitoEventsRequest(*cognitosync.SetCognitoEventsInput) (*request.Request, *cognitosync.SetCognitoEventsOutput)

	SetCognitoEvents(*cognitosync.SetCognitoEventsInput) (*cognitosync.SetCognitoEventsOutput, error)

	SetIdentityPoolConfigurationRequest(*cognitosync.SetIdentityPoolConfigurationInput) (*request.Request, *cognitosync.SetIdentityPoolConfigurationOutput)

	SetIdentityPoolConfiguration(*cognitosync.SetIdentityPoolConfigurationInput) (*cognitosync.SetIdentityPoolConfigurationOutput, error)

	SubscribeToDatasetRequest(*cognitosync.SubscribeToDatasetInput) (*request.Request, *cognitosync.SubscribeToDatasetOutput)

	SubscribeToDataset(*cognitosync.SubscribeToDatasetInput) (*cognitosync.SubscribeToDatasetOutput, error)

	UnsubscribeFromDatasetRequest(*cognitosync.UnsubscribeFromDatasetInput) (*request.Request, *cognitosync.UnsubscribeFromDatasetOutput)

	UnsubscribeFromDataset(*cognitosync.UnsubscribeFromDatasetInput) (*cognitosync.UnsubscribeFromDatasetOutput, error)

	UpdateRecordsRequest(*cognitosync.UpdateRecordsInput) (*request.Request, *cognitosync.UpdateRecordsOutput)

	UpdateRecords(*cognitosync.UpdateRecordsInput) (*cognitosync.UpdateRecordsOutput, error)
}

var _ CognitoSyncAPI = (*cognitosync.CognitoSync)(nil)
