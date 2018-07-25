package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Main Function
func main() {

	ec2Meta := ec2metadata.New(session.New(), aws.NewConfig())

	region, err := ec2Meta.Region()
	if err != nil {
		panic(err)
	}
	instanceIDDoc, err := ec2Meta.GetInstanceIdentityDocument()
	if err != nil {
		panic(err)
	}

	instanceID := instanceIDDoc.InstanceID

	svc := ec2.New(session.New(&aws.Config{Region: aws.String(region)}))

	instances, err := svc.DescribeInstances(&ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("instance-id"),
				Values: []*string{
					aws.String(instanceID),
				},
			},
		},
	})

	if err != nil {
		panic(err)
	}

	for _, r := range instances.Reservations {
		for _, i := range r.Instances {
			fmt.Printf("%s", aws.StringValue(i.PrivateDnsName))
		}
	}

}
