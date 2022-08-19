package main

import (
	"os"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Object struct {
	Pk string `json:"pk"`
	Sk string `json:"sk"`
	Name string `json:"name"`
}

type Object2 struct {
	Pk string `json:"pk"`
	Sk string `json:"sk"`
}

func getHello (event Object2) (string, error) {
	TABLE_NAME := os.Getenv("GREETINGS_TABLE")


	sess,err := session.NewSession(&aws.Config{})
	if err != nil {
		return "", err
	}

	svc := dynamodb.New(sess)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: aws.String(event.Pk),
			},
			"sk": {
				S: aws.String(event.Sk),
			},
		},
		
	})

	if err != nil {
		return "", err
	}

	if result.Item == nil {
		return "", nil
	}

	object := Object{}

	err1 := dynamodbattribute.UnmarshalMap(result.Item, &object)


	if err1 != nil {
		return "", err
	}

	str := "Hello " + object.Name + "!"
	return str, nil

}


func main () {
	lambda.Start(getHello)
}