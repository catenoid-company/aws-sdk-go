// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package autoscaling_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/catenoid-company/aws-sdk-go/aws"
	"github.com/catenoid-company/aws-sdk-go/aws/session"
	"github.com/catenoid-company/aws-sdk-go/service/autoscaling"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleAutoScaling_AttachInstances() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.AttachInstancesInput{
		AutoScalingGroupName: aws.String("ResourceName"), // Required
		InstanceIds: []*string{
			aws.String("XmlStringMaxLen19"), // Required
			// More values...
		},
	}
	resp, err := svc.AttachInstances(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_AttachLoadBalancers() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.AttachLoadBalancersInput{
		AutoScalingGroupName: aws.String("ResourceName"),
		LoadBalancerNames: []*string{
			aws.String("XmlStringMaxLen255"), // Required
			// More values...
		},
	}
	resp, err := svc.AttachLoadBalancers(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_CompleteLifecycleAction() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.CompleteLifecycleActionInput{
		AutoScalingGroupName:  aws.String("ResourceName"),          // Required
		LifecycleActionResult: aws.String("LifecycleActionResult"), // Required
		LifecycleHookName:     aws.String("AsciiStringMaxLen255"),  // Required
		InstanceId:            aws.String("XmlStringMaxLen19"),
		LifecycleActionToken:  aws.String("LifecycleActionToken"),
	}
	resp, err := svc.CompleteLifecycleAction(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_CreateAutoScalingGroup() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.CreateAutoScalingGroupInput{
		AutoScalingGroupName: aws.String("XmlStringMaxLen255"), // Required
		MaxSize:              aws.Int64(1),                     // Required
		MinSize:              aws.Int64(1),                     // Required
		AvailabilityZones: []*string{
			aws.String("XmlStringMaxLen255"), // Required
			// More values...
		},
		DefaultCooldown:         aws.Int64(1),
		DesiredCapacity:         aws.Int64(1),
		HealthCheckGracePeriod:  aws.Int64(1),
		HealthCheckType:         aws.String("XmlStringMaxLen32"),
		InstanceId:              aws.String("XmlStringMaxLen19"),
		LaunchConfigurationName: aws.String("ResourceName"),
		LoadBalancerNames: []*string{
			aws.String("XmlStringMaxLen255"), // Required
			// More values...
		},
		NewInstancesProtectedFromScaleIn: aws.Bool(true),
		PlacementGroup:                   aws.String("XmlStringMaxLen255"),
		Tags: []*autoscaling.Tag{
			{ // Required
				Key:               aws.String("TagKey"), // Required
				PropagateAtLaunch: aws.Bool(true),
				ResourceId:        aws.String("XmlString"),
				ResourceType:      aws.String("XmlString"),
				Value:             aws.String("TagValue"),
			},
			// More values...
		},
		TerminationPolicies: []*string{
			aws.String("XmlStringMaxLen1600"), // Required
			// More values...
		},
		VPCZoneIdentifier: aws.String("XmlStringMaxLen255"),
	}
	resp, err := svc.CreateAutoScalingGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_CreateLaunchConfiguration() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.CreateLaunchConfigurationInput{
		LaunchConfigurationName:  aws.String("XmlStringMaxLen255"), // Required
		AssociatePublicIpAddress: aws.Bool(true),
		BlockDeviceMappings: []*autoscaling.BlockDeviceMapping{
			{ // Required
				DeviceName: aws.String("XmlStringMaxLen255"), // Required
				Ebs: &autoscaling.Ebs{
					DeleteOnTermination: aws.Bool(true),
					Encrypted:           aws.Bool(true),
					Iops:                aws.Int64(1),
					SnapshotId:          aws.String("XmlStringMaxLen255"),
					VolumeSize:          aws.Int64(1),
					VolumeType:          aws.String("BlockDeviceEbsVolumeType"),
				},
				NoDevice:    aws.Bool(true),
				VirtualName: aws.String("XmlStringMaxLen255"),
			},
			// More values...
		},
		ClassicLinkVPCId: aws.String("XmlStringMaxLen255"),
		ClassicLinkVPCSecurityGroups: []*string{
			aws.String("XmlStringMaxLen255"), // Required
			// More values...
		},
		EbsOptimized:       aws.Bool(true),
		IamInstanceProfile: aws.String("XmlStringMaxLen1600"),
		ImageId:            aws.String("XmlStringMaxLen255"),
		InstanceId:         aws.String("XmlStringMaxLen19"),
		InstanceMonitoring: &autoscaling.InstanceMonitoring{
			Enabled: aws.Bool(true),
		},
		InstanceType:     aws.String("XmlStringMaxLen255"),
		KernelId:         aws.String("XmlStringMaxLen255"),
		KeyName:          aws.String("XmlStringMaxLen255"),
		PlacementTenancy: aws.String("XmlStringMaxLen64"),
		RamdiskId:        aws.String("XmlStringMaxLen255"),
		SecurityGroups: []*string{
			aws.String("XmlString"), // Required
			// More values...
		},
		SpotPrice: aws.String("SpotPrice"),
		UserData:  aws.String("XmlStringUserData"),
	}
	resp, err := svc.CreateLaunchConfiguration(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_CreateOrUpdateTags() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.CreateOrUpdateTagsInput{
		Tags: []*autoscaling.Tag{ // Required
			{ // Required
				Key:               aws.String("TagKey"), // Required
				PropagateAtLaunch: aws.Bool(true),
				ResourceId:        aws.String("XmlString"),
				ResourceType:      aws.String("XmlString"),
				Value:             aws.String("TagValue"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateOrUpdateTags(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_DeleteAutoScalingGroup() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.DeleteAutoScalingGroupInput{
		AutoScalingGroupName: aws.String("ResourceName"), // Required
		ForceDelete:          aws.Bool(true),
	}
	resp, err := svc.DeleteAutoScalingGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_DeleteLaunchConfiguration() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.DeleteLaunchConfigurationInput{
		LaunchConfigurationName: aws.String("ResourceName"), // Required
	}
	resp, err := svc.DeleteLaunchConfiguration(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_DeleteLifecycleHook() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.DeleteLifecycleHookInput{
		AutoScalingGroupName: aws.String("ResourceName"),         // Required
		LifecycleHookName:    aws.String("AsciiStringMaxLen255"), // Required
	}
	resp, err := svc.DeleteLifecycleHook(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_DeleteNotificationConfiguration() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.DeleteNotificationConfigurationInput{
		AutoScalingGroupName: aws.String("ResourceName"), // Required
		TopicARN:             aws.String("ResourceName"), // Required
	}
	resp, err := svc.DeleteNotificationConfiguration(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_DeletePolicy() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.DeletePolicyInput{
		PolicyName:           aws.String("ResourceName"), // Required
		AutoScalingGroupName: aws.String("ResourceName"),
	}
	resp, err := svc.DeletePolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_DeleteScheduledAction() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.DeleteScheduledActionInput{
		ScheduledActionName:  aws.String("ResourceName"), // Required
		AutoScalingGroupName: aws.String("ResourceName"),
	}
	resp, err := svc.DeleteScheduledAction(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_DeleteTags() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.DeleteTagsInput{
		Tags: []*autoscaling.Tag{ // Required
			{ // Required
				Key:               aws.String("TagKey"), // Required
				PropagateAtLaunch: aws.Bool(true),
				ResourceId:        aws.String("XmlString"),
				ResourceType:      aws.String("XmlString"),
				Value:             aws.String("TagValue"),
			},
			// More values...
		},
	}
	resp, err := svc.DeleteTags(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_DescribeAccountLimits() {
	svc := autoscaling.New(session.New())

	var params *autoscaling.DescribeAccountLimitsInput
	resp, err := svc.DescribeAccountLimits(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_DescribeAdjustmentTypes() {
	svc := autoscaling.New(session.New())

	var params *autoscaling.DescribeAdjustmentTypesInput
	resp, err := svc.DescribeAdjustmentTypes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_DescribeAutoScalingGroups() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.DescribeAutoScalingGroupsInput{
		AutoScalingGroupNames: []*string{
			aws.String("ResourceName"), // Required
			// More values...
		},
		MaxRecords: aws.Int64(1),
		NextToken:  aws.String("XmlString"),
	}
	resp, err := svc.DescribeAutoScalingGroups(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_DescribeAutoScalingInstances() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.DescribeAutoScalingInstancesInput{
		InstanceIds: []*string{
			aws.String("XmlStringMaxLen19"), // Required
			// More values...
		},
		MaxRecords: aws.Int64(1),
		NextToken:  aws.String("XmlString"),
	}
	resp, err := svc.DescribeAutoScalingInstances(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_DescribeAutoScalingNotificationTypes() {
	svc := autoscaling.New(session.New())

	var params *autoscaling.DescribeAutoScalingNotificationTypesInput
	resp, err := svc.DescribeAutoScalingNotificationTypes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_DescribeLaunchConfigurations() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.DescribeLaunchConfigurationsInput{
		LaunchConfigurationNames: []*string{
			aws.String("ResourceName"), // Required
			// More values...
		},
		MaxRecords: aws.Int64(1),
		NextToken:  aws.String("XmlString"),
	}
	resp, err := svc.DescribeLaunchConfigurations(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_DescribeLifecycleHookTypes() {
	svc := autoscaling.New(session.New())

	var params *autoscaling.DescribeLifecycleHookTypesInput
	resp, err := svc.DescribeLifecycleHookTypes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_DescribeLifecycleHooks() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.DescribeLifecycleHooksInput{
		AutoScalingGroupName: aws.String("ResourceName"), // Required
		LifecycleHookNames: []*string{
			aws.String("AsciiStringMaxLen255"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeLifecycleHooks(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_DescribeLoadBalancers() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.DescribeLoadBalancersInput{
		AutoScalingGroupName: aws.String("ResourceName"), // Required
		MaxRecords:           aws.Int64(1),
		NextToken:            aws.String("XmlString"),
	}
	resp, err := svc.DescribeLoadBalancers(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_DescribeMetricCollectionTypes() {
	svc := autoscaling.New(session.New())

	var params *autoscaling.DescribeMetricCollectionTypesInput
	resp, err := svc.DescribeMetricCollectionTypes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_DescribeNotificationConfigurations() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.DescribeNotificationConfigurationsInput{
		AutoScalingGroupNames: []*string{
			aws.String("ResourceName"), // Required
			// More values...
		},
		MaxRecords: aws.Int64(1),
		NextToken:  aws.String("XmlString"),
	}
	resp, err := svc.DescribeNotificationConfigurations(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_DescribePolicies() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.DescribePoliciesInput{
		AutoScalingGroupName: aws.String("ResourceName"),
		MaxRecords:           aws.Int64(1),
		NextToken:            aws.String("XmlString"),
		PolicyNames: []*string{
			aws.String("ResourceName"), // Required
			// More values...
		},
		PolicyTypes: []*string{
			aws.String("XmlStringMaxLen64"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribePolicies(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_DescribeScalingActivities() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.DescribeScalingActivitiesInput{
		ActivityIds: []*string{
			aws.String("XmlString"), // Required
			// More values...
		},
		AutoScalingGroupName: aws.String("ResourceName"),
		MaxRecords:           aws.Int64(1),
		NextToken:            aws.String("XmlString"),
	}
	resp, err := svc.DescribeScalingActivities(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_DescribeScalingProcessTypes() {
	svc := autoscaling.New(session.New())

	var params *autoscaling.DescribeScalingProcessTypesInput
	resp, err := svc.DescribeScalingProcessTypes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_DescribeScheduledActions() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.DescribeScheduledActionsInput{
		AutoScalingGroupName: aws.String("ResourceName"),
		EndTime:              aws.Time(time.Now()),
		MaxRecords:           aws.Int64(1),
		NextToken:            aws.String("XmlString"),
		ScheduledActionNames: []*string{
			aws.String("ResourceName"), // Required
			// More values...
		},
		StartTime: aws.Time(time.Now()),
	}
	resp, err := svc.DescribeScheduledActions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_DescribeTags() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.DescribeTagsInput{
		Filters: []*autoscaling.Filter{
			{ // Required
				Name: aws.String("XmlString"),
				Values: []*string{
					aws.String("XmlString"), // Required
					// More values...
				},
			},
			// More values...
		},
		MaxRecords: aws.Int64(1),
		NextToken:  aws.String("XmlString"),
	}
	resp, err := svc.DescribeTags(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_DescribeTerminationPolicyTypes() {
	svc := autoscaling.New(session.New())

	var params *autoscaling.DescribeTerminationPolicyTypesInput
	resp, err := svc.DescribeTerminationPolicyTypes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_DetachInstances() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.DetachInstancesInput{
		AutoScalingGroupName:           aws.String("ResourceName"), // Required
		ShouldDecrementDesiredCapacity: aws.Bool(true),             // Required
		InstanceIds: []*string{
			aws.String("XmlStringMaxLen19"), // Required
			// More values...
		},
	}
	resp, err := svc.DetachInstances(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_DetachLoadBalancers() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.DetachLoadBalancersInput{
		AutoScalingGroupName: aws.String("ResourceName"),
		LoadBalancerNames: []*string{
			aws.String("XmlStringMaxLen255"), // Required
			// More values...
		},
	}
	resp, err := svc.DetachLoadBalancers(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_DisableMetricsCollection() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.DisableMetricsCollectionInput{
		AutoScalingGroupName: aws.String("ResourceName"), // Required
		Metrics: []*string{
			aws.String("XmlStringMaxLen255"), // Required
			// More values...
		},
	}
	resp, err := svc.DisableMetricsCollection(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_EnableMetricsCollection() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.EnableMetricsCollectionInput{
		AutoScalingGroupName: aws.String("ResourceName"),       // Required
		Granularity:          aws.String("XmlStringMaxLen255"), // Required
		Metrics: []*string{
			aws.String("XmlStringMaxLen255"), // Required
			// More values...
		},
	}
	resp, err := svc.EnableMetricsCollection(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_EnterStandby() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.EnterStandbyInput{
		AutoScalingGroupName:           aws.String("ResourceName"), // Required
		ShouldDecrementDesiredCapacity: aws.Bool(true),             // Required
		InstanceIds: []*string{
			aws.String("XmlStringMaxLen19"), // Required
			// More values...
		},
	}
	resp, err := svc.EnterStandby(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_ExecutePolicy() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.ExecutePolicyInput{
		PolicyName:           aws.String("ResourceName"), // Required
		AutoScalingGroupName: aws.String("ResourceName"),
		BreachThreshold:      aws.Float64(1.0),
		HonorCooldown:        aws.Bool(true),
		MetricValue:          aws.Float64(1.0),
	}
	resp, err := svc.ExecutePolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_ExitStandby() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.ExitStandbyInput{
		AutoScalingGroupName: aws.String("ResourceName"), // Required
		InstanceIds: []*string{
			aws.String("XmlStringMaxLen19"), // Required
			// More values...
		},
	}
	resp, err := svc.ExitStandby(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_PutLifecycleHook() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.PutLifecycleHookInput{
		AutoScalingGroupName:  aws.String("ResourceName"),         // Required
		LifecycleHookName:     aws.String("AsciiStringMaxLen255"), // Required
		DefaultResult:         aws.String("LifecycleActionResult"),
		HeartbeatTimeout:      aws.Int64(1),
		LifecycleTransition:   aws.String("LifecycleTransition"),
		NotificationMetadata:  aws.String("XmlStringMaxLen1023"),
		NotificationTargetARN: aws.String("NotificationTargetResourceName"),
		RoleARN:               aws.String("ResourceName"),
	}
	resp, err := svc.PutLifecycleHook(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_PutNotificationConfiguration() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.PutNotificationConfigurationInput{
		AutoScalingGroupName: aws.String("ResourceName"), // Required
		NotificationTypes: []*string{ // Required
			aws.String("XmlStringMaxLen255"), // Required
			// More values...
		},
		TopicARN: aws.String("ResourceName"), // Required
	}
	resp, err := svc.PutNotificationConfiguration(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_PutScalingPolicy() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.PutScalingPolicyInput{
		AdjustmentType:          aws.String("XmlStringMaxLen255"), // Required
		AutoScalingGroupName:    aws.String("ResourceName"),       // Required
		PolicyName:              aws.String("XmlStringMaxLen255"), // Required
		Cooldown:                aws.Int64(1),
		EstimatedInstanceWarmup: aws.Int64(1),
		MetricAggregationType:   aws.String("XmlStringMaxLen32"),
		MinAdjustmentMagnitude:  aws.Int64(1),
		MinAdjustmentStep:       aws.Int64(1),
		PolicyType:              aws.String("XmlStringMaxLen64"),
		ScalingAdjustment:       aws.Int64(1),
		StepAdjustments: []*autoscaling.StepAdjustment{
			{ // Required
				ScalingAdjustment:        aws.Int64(1), // Required
				MetricIntervalLowerBound: aws.Float64(1.0),
				MetricIntervalUpperBound: aws.Float64(1.0),
			},
			// More values...
		},
	}
	resp, err := svc.PutScalingPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_PutScheduledUpdateGroupAction() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.PutScheduledUpdateGroupActionInput{
		AutoScalingGroupName: aws.String("ResourceName"),       // Required
		ScheduledActionName:  aws.String("XmlStringMaxLen255"), // Required
		DesiredCapacity:      aws.Int64(1),
		EndTime:              aws.Time(time.Now()),
		MaxSize:              aws.Int64(1),
		MinSize:              aws.Int64(1),
		Recurrence:           aws.String("XmlStringMaxLen255"),
		StartTime:            aws.Time(time.Now()),
		Time:                 aws.Time(time.Now()),
	}
	resp, err := svc.PutScheduledUpdateGroupAction(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_RecordLifecycleActionHeartbeat() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.RecordLifecycleActionHeartbeatInput{
		AutoScalingGroupName: aws.String("ResourceName"),         // Required
		LifecycleHookName:    aws.String("AsciiStringMaxLen255"), // Required
		InstanceId:           aws.String("XmlStringMaxLen19"),
		LifecycleActionToken: aws.String("LifecycleActionToken"),
	}
	resp, err := svc.RecordLifecycleActionHeartbeat(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_ResumeProcesses() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.ScalingProcessQuery{
		AutoScalingGroupName: aws.String("ResourceName"), // Required
		ScalingProcesses: []*string{
			aws.String("XmlStringMaxLen255"), // Required
			// More values...
		},
	}
	resp, err := svc.ResumeProcesses(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_SetDesiredCapacity() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.SetDesiredCapacityInput{
		AutoScalingGroupName: aws.String("ResourceName"), // Required
		DesiredCapacity:      aws.Int64(1),               // Required
		HonorCooldown:        aws.Bool(true),
	}
	resp, err := svc.SetDesiredCapacity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_SetInstanceHealth() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.SetInstanceHealthInput{
		HealthStatus:             aws.String("XmlStringMaxLen32"), // Required
		InstanceId:               aws.String("XmlStringMaxLen19"), // Required
		ShouldRespectGracePeriod: aws.Bool(true),
	}
	resp, err := svc.SetInstanceHealth(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_SetInstanceProtection() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.SetInstanceProtectionInput{
		AutoScalingGroupName: aws.String("ResourceName"), // Required
		InstanceIds: []*string{ // Required
			aws.String("XmlStringMaxLen19"), // Required
			// More values...
		},
		ProtectedFromScaleIn: aws.Bool(true), // Required
	}
	resp, err := svc.SetInstanceProtection(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_SuspendProcesses() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.ScalingProcessQuery{
		AutoScalingGroupName: aws.String("ResourceName"), // Required
		ScalingProcesses: []*string{
			aws.String("XmlStringMaxLen255"), // Required
			// More values...
		},
	}
	resp, err := svc.SuspendProcesses(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_TerminateInstanceInAutoScalingGroup() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.TerminateInstanceInAutoScalingGroupInput{
		InstanceId:                     aws.String("XmlStringMaxLen19"), // Required
		ShouldDecrementDesiredCapacity: aws.Bool(true),                  // Required
	}
	resp, err := svc.TerminateInstanceInAutoScalingGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAutoScaling_UpdateAutoScalingGroup() {
	svc := autoscaling.New(session.New())

	params := &autoscaling.UpdateAutoScalingGroupInput{
		AutoScalingGroupName: aws.String("ResourceName"), // Required
		AvailabilityZones: []*string{
			aws.String("XmlStringMaxLen255"), // Required
			// More values...
		},
		DefaultCooldown:         aws.Int64(1),
		DesiredCapacity:         aws.Int64(1),
		HealthCheckGracePeriod:  aws.Int64(1),
		HealthCheckType:         aws.String("XmlStringMaxLen32"),
		LaunchConfigurationName: aws.String("ResourceName"),
		MaxSize:                 aws.Int64(1),
		MinSize:                 aws.Int64(1),
		NewInstancesProtectedFromScaleIn: aws.Bool(true),
		PlacementGroup:                   aws.String("XmlStringMaxLen255"),
		TerminationPolicies: []*string{
			aws.String("XmlStringMaxLen1600"), // Required
			// More values...
		},
		VPCZoneIdentifier: aws.String("XmlStringMaxLen255"),
	}
	resp, err := svc.UpdateAutoScalingGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
