// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package kinesisiface provides an interface for the Amazon Kinesis.
package kinesisiface

import (
	"github.com/catenoid-company/aws-sdk-go/aws/request"
	"github.com/catenoid-company/aws-sdk-go/service/kinesis"
)

// KinesisAPI is the interface type for kinesis.Kinesis.
type KinesisAPI interface {
	AddTagsToStreamRequest(*kinesis.AddTagsToStreamInput) (*request.Request, *kinesis.AddTagsToStreamOutput)

	AddTagsToStream(*kinesis.AddTagsToStreamInput) (*kinesis.AddTagsToStreamOutput, error)

	CreateStreamRequest(*kinesis.CreateStreamInput) (*request.Request, *kinesis.CreateStreamOutput)

	CreateStream(*kinesis.CreateStreamInput) (*kinesis.CreateStreamOutput, error)

	DecreaseStreamRetentionPeriodRequest(*kinesis.DecreaseStreamRetentionPeriodInput) (*request.Request, *kinesis.DecreaseStreamRetentionPeriodOutput)

	DecreaseStreamRetentionPeriod(*kinesis.DecreaseStreamRetentionPeriodInput) (*kinesis.DecreaseStreamRetentionPeriodOutput, error)

	DeleteStreamRequest(*kinesis.DeleteStreamInput) (*request.Request, *kinesis.DeleteStreamOutput)

	DeleteStream(*kinesis.DeleteStreamInput) (*kinesis.DeleteStreamOutput, error)

	DescribeStreamRequest(*kinesis.DescribeStreamInput) (*request.Request, *kinesis.DescribeStreamOutput)

	DescribeStream(*kinesis.DescribeStreamInput) (*kinesis.DescribeStreamOutput, error)

	DescribeStreamPages(*kinesis.DescribeStreamInput, func(*kinesis.DescribeStreamOutput, bool) bool) error

	DisableEnhancedMonitoringRequest(*kinesis.DisableEnhancedMonitoringInput) (*request.Request, *kinesis.EnhancedMonitoringOutput)

	DisableEnhancedMonitoring(*kinesis.DisableEnhancedMonitoringInput) (*kinesis.EnhancedMonitoringOutput, error)

	EnableEnhancedMonitoringRequest(*kinesis.EnableEnhancedMonitoringInput) (*request.Request, *kinesis.EnhancedMonitoringOutput)

	EnableEnhancedMonitoring(*kinesis.EnableEnhancedMonitoringInput) (*kinesis.EnhancedMonitoringOutput, error)

	GetRecordsRequest(*kinesis.GetRecordsInput) (*request.Request, *kinesis.GetRecordsOutput)

	GetRecords(*kinesis.GetRecordsInput) (*kinesis.GetRecordsOutput, error)

	GetShardIteratorRequest(*kinesis.GetShardIteratorInput) (*request.Request, *kinesis.GetShardIteratorOutput)

	GetShardIterator(*kinesis.GetShardIteratorInput) (*kinesis.GetShardIteratorOutput, error)

	IncreaseStreamRetentionPeriodRequest(*kinesis.IncreaseStreamRetentionPeriodInput) (*request.Request, *kinesis.IncreaseStreamRetentionPeriodOutput)

	IncreaseStreamRetentionPeriod(*kinesis.IncreaseStreamRetentionPeriodInput) (*kinesis.IncreaseStreamRetentionPeriodOutput, error)

	ListStreamsRequest(*kinesis.ListStreamsInput) (*request.Request, *kinesis.ListStreamsOutput)

	ListStreams(*kinesis.ListStreamsInput) (*kinesis.ListStreamsOutput, error)

	ListStreamsPages(*kinesis.ListStreamsInput, func(*kinesis.ListStreamsOutput, bool) bool) error

	ListTagsForStreamRequest(*kinesis.ListTagsForStreamInput) (*request.Request, *kinesis.ListTagsForStreamOutput)

	ListTagsForStream(*kinesis.ListTagsForStreamInput) (*kinesis.ListTagsForStreamOutput, error)

	MergeShardsRequest(*kinesis.MergeShardsInput) (*request.Request, *kinesis.MergeShardsOutput)

	MergeShards(*kinesis.MergeShardsInput) (*kinesis.MergeShardsOutput, error)

	PutRecordRequest(*kinesis.PutRecordInput) (*request.Request, *kinesis.PutRecordOutput)

	PutRecord(*kinesis.PutRecordInput) (*kinesis.PutRecordOutput, error)

	PutRecordsRequest(*kinesis.PutRecordsInput) (*request.Request, *kinesis.PutRecordsOutput)

	PutRecords(*kinesis.PutRecordsInput) (*kinesis.PutRecordsOutput, error)

	RemoveTagsFromStreamRequest(*kinesis.RemoveTagsFromStreamInput) (*request.Request, *kinesis.RemoveTagsFromStreamOutput)

	RemoveTagsFromStream(*kinesis.RemoveTagsFromStreamInput) (*kinesis.RemoveTagsFromStreamOutput, error)

	SplitShardRequest(*kinesis.SplitShardInput) (*request.Request, *kinesis.SplitShardOutput)

	SplitShard(*kinesis.SplitShardInput) (*kinesis.SplitShardOutput, error)
}

var _ KinesisAPI = (*kinesis.Kinesis)(nil)
