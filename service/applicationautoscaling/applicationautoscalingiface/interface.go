// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package applicationautoscalingiface provides an interface for the Application Auto Scaling.
package applicationautoscalingiface

import (
	"github.com/catenoid-company/aws-sdk-go/aws/request"
	"github.com/catenoid-company/aws-sdk-go/service/applicationautoscaling"
)

// ApplicationAutoScalingAPI is the interface type for applicationautoscaling.ApplicationAutoScaling.
type ApplicationAutoScalingAPI interface {
	DeleteScalingPolicyRequest(*applicationautoscaling.DeleteScalingPolicyInput) (*request.Request, *applicationautoscaling.DeleteScalingPolicyOutput)

	DeleteScalingPolicy(*applicationautoscaling.DeleteScalingPolicyInput) (*applicationautoscaling.DeleteScalingPolicyOutput, error)

	DeregisterScalableTargetRequest(*applicationautoscaling.DeregisterScalableTargetInput) (*request.Request, *applicationautoscaling.DeregisterScalableTargetOutput)

	DeregisterScalableTarget(*applicationautoscaling.DeregisterScalableTargetInput) (*applicationautoscaling.DeregisterScalableTargetOutput, error)

	DescribeScalableTargetsRequest(*applicationautoscaling.DescribeScalableTargetsInput) (*request.Request, *applicationautoscaling.DescribeScalableTargetsOutput)

	DescribeScalableTargets(*applicationautoscaling.DescribeScalableTargetsInput) (*applicationautoscaling.DescribeScalableTargetsOutput, error)

	DescribeScalableTargetsPages(*applicationautoscaling.DescribeScalableTargetsInput, func(*applicationautoscaling.DescribeScalableTargetsOutput, bool) bool) error

	DescribeScalingActivitiesRequest(*applicationautoscaling.DescribeScalingActivitiesInput) (*request.Request, *applicationautoscaling.DescribeScalingActivitiesOutput)

	DescribeScalingActivities(*applicationautoscaling.DescribeScalingActivitiesInput) (*applicationautoscaling.DescribeScalingActivitiesOutput, error)

	DescribeScalingActivitiesPages(*applicationautoscaling.DescribeScalingActivitiesInput, func(*applicationautoscaling.DescribeScalingActivitiesOutput, bool) bool) error

	DescribeScalingPoliciesRequest(*applicationautoscaling.DescribeScalingPoliciesInput) (*request.Request, *applicationautoscaling.DescribeScalingPoliciesOutput)

	DescribeScalingPolicies(*applicationautoscaling.DescribeScalingPoliciesInput) (*applicationautoscaling.DescribeScalingPoliciesOutput, error)

	DescribeScalingPoliciesPages(*applicationautoscaling.DescribeScalingPoliciesInput, func(*applicationautoscaling.DescribeScalingPoliciesOutput, bool) bool) error

	PutScalingPolicyRequest(*applicationautoscaling.PutScalingPolicyInput) (*request.Request, *applicationautoscaling.PutScalingPolicyOutput)

	PutScalingPolicy(*applicationautoscaling.PutScalingPolicyInput) (*applicationautoscaling.PutScalingPolicyOutput, error)

	RegisterScalableTargetRequest(*applicationautoscaling.RegisterScalableTargetInput) (*request.Request, *applicationautoscaling.RegisterScalableTargetOutput)

	RegisterScalableTarget(*applicationautoscaling.RegisterScalableTargetInput) (*applicationautoscaling.RegisterScalableTargetOutput, error)
}

var _ ApplicationAutoScalingAPI = (*applicationautoscaling.ApplicationAutoScaling)(nil)
