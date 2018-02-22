package aws

import (
	"fmt"

	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AmazonKuberCluster struct {
	KuberClusterId     uint
	NodeSpotPrice      string
	NodeMinCount       int
	NodeMaxCount       int
	NodeImage          string
	MasterInstanceType string
	MasterImage        string
}

type CreateAmazonCluster struct {
	Master *CreateAmazonMaster `json:"master"`
	Node   *CreateAmazonNode   `json:"node"`
}

type CreateAmazonMaster struct {
	InstanceType string `json:"instanceType"`
	Image        string `json:"image"`
}

type CreateAmazonNode struct {
	SpotPrice string `json:"spotPrice"`
	MinCount  int    `json:"minCount"`
	MaxCount  int    `json:"maxCount"`
	Image     string `json:"image"`
}

type UpdateAmazonCluster struct {
	*UpdateAmazonNode `json:"node"`
}

type UpdateAmazonNode struct {
	MinCount int `json:"minCount"`
	MaxCount int `json:"maxCount"`
}

func InstancesByRegions(states []string, regions []string) {
	instanceCriteria := " "
	for _, state := range states {
		instanceCriteria += "[" + state + "]"
	}
	if len(regions) == 0 {
		var err error
		regions, err = fetchRegion()
		if err != nil {
			log.Fatalf("error: %v\n", err)
		}
	}
	for _, region := range regions {
		sess := NewSession(region)
		ec2Svc := ec2.New(sess)
		params := &ec2.DescribeInstancesInput{
			Filters: []*ec2.Filter{
				&ec2.Filter{
					Name:   aws.String("instance-state-name"),
					Values: aws.StringSlice(states),
				},
			},
		}
		result, err := ec2Svc.DescribeInstances(params)
		if err != nil {
			fmt.Println("Error", err)
		} else {
			fmt.Printf("\n\n\nFetching instace details  for region: %s with criteria: %s**\n ", region, instanceCriteria)
			if len(result.Reservations) == 0 {
				fmt.Printf("There is no instance for the for region %s with the matching Criteria:%s  \n", region, instanceCriteria)
			}
			for _, reservation := range result.Reservations {

				fmt.Println("printing instance details.....")
				for _, instance := range reservation.Instances {
					fmt.Println("instance id " + *instance.InstanceId)
					fmt.Println("current State " + *instance.State.Name)
				}
			}
			fmt.Printf("done for region %s **** \n", region)
		}
	}
}

func GetVolumes(region string) *ec2.DescribeVolumesOutput {
	svc := ec2.New(NewSession(region))
	input := &ec2.DescribeVolumesInput{}
	result, err := svc.DescribeVolumes(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			log.Fatalf("error: %v\n", err.Error())
		}
		log.Fatalf("error: %v\n", err)
	}
	return result
}

func GetVolumesForInstance(region string, instanceId string) *ec2.DescribeVolumesOutput {
	svc := ec2.New(NewSession(region))
	input := &ec2.DescribeVolumesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("attachment.instance-id"),
				Values: []*string{
					aws.String(instanceId),
				},
			},
			{
				Name: aws.String("attachment.delete-on-termination"),
				Values: []*string{
					aws.String("true"),
				},
			},
		},
	}
	result, err := svc.DescribeVolumes(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			log.Fatalf("error: %v\n", err.Error())
		}
		log.Fatalf("error: %v\n", err)
	}
	return result
}

func InstancesByTag(tagName string, tagValue, region string) *ec2.DescribeInstancesOutput {
	svc := ec2.New(NewSession(region))
	fmt.Printf("listing instances with tag name %v and value %v in: %v\n", tagName, tagValue, region)
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:" + tagName),
				Values: []*string{
					aws.String(strings.Join([]string{"*", tagValue, "*"}, "")),
				},
			},
		},
	}
	resp, err := svc.DescribeInstances(params)
	if err != nil {
		fmt.Println("There was an error listing instances in", region, err.Error())
		log.Fatal(err.Error())
	}
	return resp
}

func KuberInstances(region string) *ec2.DescribeInstancesOutput {
	svc := ec2.New(NewSession(region))
	fmt.Printf("Listing Kuber instances in: %v\n", region)
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:" + "kuber-role"),
				Values: []*string{
					aws.String(strings.Join([]string{"*"}, "")),
				},
			},
		},
	}
	resp, err := svc.DescribeInstances(params)
	if err != nil {
		fmt.Println("There was an error listing instances in", region, err.Error())
		log.Fatal(err.Error())
	}
	return resp
}

func GetLoadBalancersByTag(tagName string, tagValue, region string) []*string {
	svc := elb.New(NewSession(region))
	fmt.Printf("listing load balancers with tag name %v and value %v in: %v\n", tagName, tagValue, region)
	input := &elb.DescribeLoadBalancersInput{
	/*
		LoadBalancerNames: []*string{
			aws.String(name),
		},
	*/
	}
	resp, err := svc.DescribeLoadBalancers(input)
	if err != nil {
		fmt.Println("DescribeLoadBalancers error in", region, err.Error())
		log.Fatal(err.Error())
	}
	var names []*string
	for _, elb := range resp.LoadBalancerDescriptions {
		names = append(names, elb.LoadBalancerName)
	}
	input2 := &elb.DescribeTagsInput{
		LoadBalancerNames: names,
	}
	tags, err := svc.DescribeTags(input2)
	if err != nil {
		fmt.Println("There was an error listing instances in", region, err.Error())
		log.Fatal(err.Error())
	}
	var elbs []*string
	for _, desc := range tags.TagDescriptions {
		for _, tag := range desc.Tags {
			if *tag.Key == tagName {
				if *tag.Value == tagValue {
					elbs = append(elbs, desc.LoadBalancerName)
				}
			}
		}
	}
	return elbs
}

func RegisterInstanceToLoadBalancer(instanceId *string, loadBalancerName *string, region string) *elb.RegisterInstancesWithLoadBalancerOutput {
	svc := elb.New(NewSession(region))
	fmt.Printf("Register instance to load balancers %v and value %v in: %v\n", instanceId, loadBalancerName, region)
	input := &elb.RegisterInstancesWithLoadBalancerInput{
		Instances: []*elb.Instance{
			{
				InstanceId: instanceId,
			},
		},
		LoadBalancerName: loadBalancerName,
	}
	result, err := svc.RegisterInstancesWithLoadBalancer(input)
	if err != nil {
		fmt.Println("RegisterInstanceToLoadBalancer error in", region, err.Error())
		log.Fatal(err.Error())
	}

	return result
}

func ScaleWorkers(desiredWorkers int64, region string) *autoscaling.Group {
	svc := autoscaling.New(NewSession(region))

	asgn := "kuber.node"

	input := &autoscaling.DescribeAutoScalingGroupsInput{
		AutoScalingGroupNames: []*string{
			aws.String(asgn),
		},
	}

	result, err := svc.DescribeAutoScalingGroups(input)
	if err != nil {
		fmt.Println("ScaleWorkers error in", region, err.Error())
		log.Fatal(err.Error())
	}

	g := result.AutoScalingGroups[0]

	if &desiredWorkers != g.DesiredCapacity {
		input2 := &autoscaling.UpdateAutoScalingGroupInput{
			AutoScalingGroupName: aws.String(asgn),
			DesiredCapacity:      &desiredWorkers,
		}

		_, err := svc.UpdateAutoScalingGroup(input2)
		if err != nil {
			fmt.Println("ScaleWorkers error in", region, err.Error())
			log.Fatal(err.Error())
		}

	}

	return g

}

func TagResource(resourceId string, tagName string, tagValue, region string) {
	sess := NewSession(region)
	svc := ec2.New(sess, &aws.Config{Region: aws.String(region)})
	input := &ec2.CreateTagsInput{
		Resources: []*string{
			aws.String(resourceId),
		},
		Tags: []*ec2.Tag{
			{
				Key:   aws.String(tagName),
				Value: aws.String(tagValue),
			},
		},
	}

	_, err := svc.CreateTags(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
}

func ListS3(bucket string, region string) {
	svc := s3.New(NewSession(region))
	i := 0
	err := svc.ListObjectsPages(&s3.ListObjectsInput{
		Bucket: &bucket,
	}, func(p *s3.ListObjectsOutput, last bool) (shouldContinue bool) {
		fmt.Println("Page,", i)
		i++
		for _, obj := range p.Contents {
			fmt.Println("Object:", *obj.Key)
		}
		return true
	})
	if err != nil {
		fmt.Println("failed to list objects", err)
		return
	}
}

func NewSession(region string) *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		//		Credentials: credentials.NewSharedCredentials("", "kuber"),
	})
	if err != nil {
		fmt.Println("error", err)
	}
	// In addition, checking if your credentials have been found is fairly easy.
	_, err2 := sess.Config.Credentials.Get()
	if err2 != nil {
		fmt.Println("error", err2)
	}
	return sess
}

func fetchRegion() ([]string, error) {
	awsSession := session.Must(session.NewSession(&aws.Config{}))
	svc := ec2.New(awsSession)
	awsRegions, err := svc.DescribeRegions(&ec2.DescribeRegionsInput{})
	if err != nil {
		return nil, err
	}
	regions := make([]string, 0, len(awsRegions.Regions))
	for _, region := range awsRegions.Regions {
		regions = append(regions, *region.RegionName)
	}
	return regions, nil
}
