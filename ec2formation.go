package main

import (
	"fmt"

	"github.com/awslabs/goformation/cloudformation"
)

func main() {

	// Create a new CloudFormation template

	t := cloudformation.NewTemplate()

	//Define instance Paramaters

	/*
		t.Parameters["InstanceType"] = &instanceParams{

			Description:   "EC2 Instance Type",
			Type:          "String",
			Default:       "t2.micro",
			AllowedValues: []string{"t2.nano", "t2.micro", "t1.micro", "t2.small", "t2.medium"},
		}
	*/

	t.Parameters["InstanceType"] = &Paramaters{

		Description:   "EC2 Instance Type",
		Type:          "String",
		Default:       "t2.nano",
		AllowedValues: []string{"t2.nano", "t2.micro", "t1.micro", "t2.small", "t2.medium"},
	}

	t.Parameters["KeyName"] = &Paramaters{

		Description: "Name of an existing EC2 Keypair to enable SSH access to the instance",
		Type:        "AWS::EC2::KeyPair::KeyName",
		ConstraintDescription: "Must be the name of existing EC2 Keypar",
	}

	// An an example SNS Topic
	t.Resources["MySNSTopic"] = &cloudformation.AWSSNSTopic{
		DisplayName: "test-sns-topic-display-name",
		TopicName:   "test-sns-topic-name",
		Subscription: []cloudformation.AWSSNSTopic_Subscription{
			cloudformation.AWSSNSTopic_Subscription{
				Endpoint: "test-sns-topic-subscription-endpoint",
				Protocol: "test-sns-topic-subscription-protocol",
			},
		},
	}

	// ...and a Route 53 Hosted Zone too
	t.Resources["MyRoute53HostedZone"] = &cloudformation.AWSRoute53HostedZone{
		Name: "example.com",
	}

	// Let's see the JSON
	j, err := t.JSON()
	if err != nil {
		fmt.Printf("Failed to generate JSON: %s\n", err)
	} else {
		fmt.Printf("%s\n", string(j))
	}

	y, err := t.YAML()
	if err != nil {
		fmt.Printf("Failed to generate YAML: %s\n", err)
	} else {
		fmt.Printf("%s\n", string(y))
	}

}
