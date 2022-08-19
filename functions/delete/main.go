package main

import (
	"os"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Object struct {
	Pk string `json:"pk"`
	Sk string `json:"sk"`
	Name string `json:"name"`
}

func deleteHello (event Object) (string, error) {
	TABLE_NAME := os.Getenv("GREETINGS_TABLE")
	sess,err := session.NewSession(&aws.Config{})
	if err != nil {
		return "", err
	}


	svc := dynamodb.New(sess)


	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: aws.String(event.Pk),
			},
			"sk": {
				S: aws.String(event.Sk),
			},
		},
	}

	_, err = svc.DeleteItem(input)
	if err != nil {
		return "", err
	}

	return "Hello " + event.Name, nil
}

func main() {
	lambda.Start(deleteHello)
}