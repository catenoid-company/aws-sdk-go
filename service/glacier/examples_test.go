// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package glacier_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/catenoid-company/aws-sdk-go/aws"
	"github.com/catenoid-company/aws-sdk-go/aws/session"
	"github.com/catenoid-company/aws-sdk-go/service/glacier"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleGlacier_AbortMultipartUpload() {
	svc := glacier.New(session.New())

	params := &glacier.AbortMultipartUploadInput{
		AccountId: aws.String("string"), // Required
		UploadId:  aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
	}
	resp, err := svc.AbortMultipartUpload(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGlacier_AbortVaultLock() {
	svc := glacier.New(session.New())

	params := &glacier.AbortVaultLockInput{
		AccountId: aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
	}
	resp, err := svc.AbortVaultLock(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGlacier_AddTagsToVault() {
	svc := glacier.New(session.New())

	params := &glacier.AddTagsToVaultInput{
		AccountId: aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
		Tags: map[string]*string{
			"Key": aws.String("TagValue"), // Required
			// More values...
		},
	}
	resp, err := svc.AddTagsToVault(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGlacier_CompleteMultipartUpload() {
	svc := glacier.New(session.New())

	params := &glacier.CompleteMultipartUploadInput{
		AccountId:   aws.String("string"), // Required
		UploadId:    aws.String("string"), // Required
		VaultName:   aws.String("string"), // Required
		ArchiveSize: aws.String("string"),
		Checksum:    aws.String("string"),
	}
	resp, err := svc.CompleteMultipartUpload(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGlacier_CompleteVaultLock() {
	svc := glacier.New(session.New())

	params := &glacier.CompleteVaultLockInput{
		AccountId: aws.String("string"), // Required
		LockId:    aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
	}
	resp, err := svc.CompleteVaultLock(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGlacier_CreateVault() {
	svc := glacier.New(session.New())

	params := &glacier.CreateVaultInput{
		AccountId: aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
	}
	resp, err := svc.CreateVault(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGlacier_DeleteArchive() {
	svc := glacier.New(session.New())

	params := &glacier.DeleteArchiveInput{
		AccountId: aws.String("string"), // Required
		ArchiveId: aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
	}
	resp, err := svc.DeleteArchive(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGlacier_DeleteVault() {
	svc := glacier.New(session.New())

	params := &glacier.DeleteVaultInput{
		AccountId: aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
	}
	resp, err := svc.DeleteVault(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGlacier_DeleteVaultAccessPolicy() {
	svc := glacier.New(session.New())

	params := &glacier.DeleteVaultAccessPolicyInput{
		AccountId: aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
	}
	resp, err := svc.DeleteVaultAccessPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGlacier_DeleteVaultNotifications() {
	svc := glacier.New(session.New())

	params := &glacier.DeleteVaultNotificationsInput{
		AccountId: aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
	}
	resp, err := svc.DeleteVaultNotifications(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGlacier_DescribeJob() {
	svc := glacier.New(session.New())

	params := &glacier.DescribeJobInput{
		AccountId: aws.String("string"), // Required
		JobId:     aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
	}
	resp, err := svc.DescribeJob(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGlacier_DescribeVault() {
	svc := glacier.New(session.New())

	params := &glacier.DescribeVaultInput{
		AccountId: aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
	}
	resp, err := svc.DescribeVault(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGlacier_GetDataRetrievalPolicy() {
	svc := glacier.New(session.New())

	params := &glacier.GetDataRetrievalPolicyInput{
		AccountId: aws.String("string"), // Required
	}
	resp, err := svc.GetDataRetrievalPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGlacier_GetJobOutput() {
	svc := glacier.New(session.New())

	params := &glacier.GetJobOutputInput{
		AccountId: aws.String("string"), // Required
		JobId:     aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
		Range:     aws.String("string"),
	}
	resp, err := svc.GetJobOutput(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGlacier_GetVaultAccessPolicy() {
	svc := glacier.New(session.New())

	params := &glacier.GetVaultAccessPolicyInput{
		AccountId: aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
	}
	resp, err := svc.GetVaultAccessPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGlacier_GetVaultLock() {
	svc := glacier.New(session.New())

	params := &glacier.GetVaultLockInput{
		AccountId: aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
	}
	resp, err := svc.GetVaultLock(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGlacier_GetVaultNotifications() {
	svc := glacier.New(session.New())

	params := &glacier.GetVaultNotificationsInput{
		AccountId: aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
	}
	resp, err := svc.GetVaultNotifications(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGlacier_InitiateJob() {
	svc := glacier.New(session.New())

	params := &glacier.InitiateJobInput{
		AccountId: aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
		JobParameters: &glacier.JobParameters{
			ArchiveId:   aws.String("string"),
			Description: aws.String("string"),
			Format:      aws.String("string"),
			InventoryRetrievalParameters: &glacier.InventoryRetrievalJobInput{
				EndDate:   aws.String("string"),
				Limit:     aws.String("string"),
				Marker:    aws.String("string"),
				StartDate: aws.String("string"),
			},
			RetrievalByteRange: aws.String("string"),
			SNSTopic:           aws.String("string"),
			Type:               aws.String("string"),
		},
	}
	resp, err := svc.InitiateJob(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGlacier_InitiateMultipartUpload() {
	svc := glacier.New(session.New())

	params := &glacier.InitiateMultipartUploadInput{
		AccountId:          aws.String("string"), // Required
		VaultName:          aws.String("string"), // Required
		ArchiveDescription: aws.String("string"),
		PartSize:           aws.String("string"),
	}
	resp, err := svc.InitiateMultipartUpload(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGlacier_InitiateVaultLock() {
	svc := glacier.New(session.New())

	params := &glacier.InitiateVaultLockInput{
		AccountId: aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
		Policy: &glacier.VaultLockPolicy{
			Policy: aws.String("string"),
		},
	}
	resp, err := svc.InitiateVaultLock(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGlacier_ListJobs() {
	svc := glacier.New(session.New())

	params := &glacier.ListJobsInput{
		AccountId:  aws.String("string"), // Required
		VaultName:  aws.String("string"), // Required
		Completed:  aws.String("string"),
		Limit:      aws.String("string"),
		Marker:     aws.String("string"),
		Statuscode: aws.String("string"),
	}
	resp, err := svc.ListJobs(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGlacier_ListMultipartUploads() {
	svc := glacier.New(session.New())

	params := &glacier.ListMultipartUploadsInput{
		AccountId: aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
		Limit:     aws.String("string"),
		Marker:    aws.String("string"),
	}
	resp, err := svc.ListMultipartUploads(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGlacier_ListParts() {
	svc := glacier.New(session.New())

	params := &glacier.ListPartsInput{
		AccountId: aws.String("string"), // Required
		UploadId:  aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
		Limit:     aws.String("string"),
		Marker:    aws.String("string"),
	}
	resp, err := svc.ListParts(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGlacier_ListTagsForVault() {
	svc := glacier.New(session.New())

	params := &glacier.ListTagsForVaultInput{
		AccountId: aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
	}
	resp, err := svc.ListTagsForVault(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGlacier_ListVaults() {
	svc := glacier.New(session.New())

	params := &glacier.ListVaultsInput{
		AccountId: aws.String("string"), // Required
		Limit:     aws.String("string"),
		Marker:    aws.String("string"),
	}
	resp, err := svc.ListVaults(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGlacier_RemoveTagsFromVault() {
	svc := glacier.New(session.New())

	params := &glacier.RemoveTagsFromVaultInput{
		AccountId: aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
		TagKeys: []*string{
			aws.String("string"), // Required
			// More values...
		},
	}
	resp, err := svc.RemoveTagsFromVault(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGlacier_SetDataRetrievalPolicy() {
	svc := glacier.New(session.New())

	params := &glacier.SetDataRetrievalPolicyInput{
		AccountId: aws.String("string"), // Required
		Policy: &glacier.DataRetrievalPolicy{
			Rules: []*glacier.DataRetrievalRule{
				{ // Required
					BytesPerHour: aws.Int64(1),
					Strategy:     aws.String("string"),
				},
				// More values...
			},
		},
	}
	resp, err := svc.SetDataRetrievalPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGlacier_SetVaultAccessPolicy() {
	svc := glacier.New(session.New())

	params := &glacier.SetVaultAccessPolicyInput{
		AccountId: aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
		Policy: &glacier.VaultAccessPolicy{
			Policy: aws.String("string"),
		},
	}
	resp, err := svc.SetVaultAccessPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGlacier_SetVaultNotifications() {
	svc := glacier.New(session.New())

	params := &glacier.SetVaultNotificationsInput{
		AccountId: aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
		VaultNotificationConfig: &glacier.VaultNotificationConfig{
			Events: []*string{
				aws.String("string"), // Required
				// More values...
			},
			SNSTopic: aws.String("string"),
		},
	}
	resp, err := svc.SetVaultNotifications(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGlacier_UploadArchive() {
	svc := glacier.New(session.New())

	params := &glacier.UploadArchiveInput{
		AccountId:          aws.String("string"), // Required
		VaultName:          aws.String("string"), // Required
		ArchiveDescription: aws.String("string"),
		Body:               bytes.NewReader([]byte("PAYLOAD")),
		Checksum:           aws.String("string"),
	}
	resp, err := svc.UploadArchive(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGlacier_UploadMultipartPart() {
	svc := glacier.New(session.New())

	params := &glacier.UploadMultipartPartInput{
		AccountId: aws.String("string"), // Required
		UploadId:  aws.String("string"), // Required
		VaultName: aws.String("string"), // Required
		Body:      bytes.NewReader([]byte("PAYLOAD")),
		Checksum:  aws.String("string"),
		Range:     aws.String("string"),
	}
	resp, err := svc.UploadMultipartPart(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
