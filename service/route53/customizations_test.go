package route53_test

import (
	"testing"

	"github.com/catenoid-company/aws-sdk-go/aws"
	"github.com/catenoid-company/aws-sdk-go/awstesting"
	"github.com/catenoid-company/aws-sdk-go/awstesting/unit"
	"github.com/catenoid-company/aws-sdk-go/service/route53"
)

func TestBuildCorrectURI(t *testing.T) {
	svc := route53.New(unit.Session)
	svc.Handlers.Validate.Clear()
	req, _ := svc.GetHostedZoneRequest(&route53.GetHostedZoneInput{
		Id: aws.String("/hostedzone/ABCDEFG"),
	})

	req.Build()

	awstesting.Match(t, `\/hostedzone\/ABCDEFG$`, req.HTTPRequest.URL.String())
}
