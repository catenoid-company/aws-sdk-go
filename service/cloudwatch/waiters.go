// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cloudwatch

import (
	"github.com/catenoid-company/aws-sdk-go/private/waiter"
)

func (c *CloudWatch) WaitUntilAlarmExists(input *DescribeAlarmsInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeAlarms",
		Delay:       5,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "path",
				Argument: "length(MetricAlarms[]) > `0`",
				Expected: true,
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}
