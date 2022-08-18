package put

import (
	"os"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Object struct {
	pk string `json:"pk"`
	sk string `json:"sk"`
	name string `json:"name"`
}

func saveHello(event Object) (string, error) {
	TABLE_NAME := os.Getenv("GREETINGS_TABLE")

	sess,err := session.NewSession(&aws.Config{})
	if err != nil {
		return "", err
	}

	svc := dynamodb.New(sess)
	object := Object{event.pk, event.sk, event.name}
	item, err := dynamodbattribute.MarshalMap(object)
	if err != nil {
		return "", err
	}
	input := &dynamodb.PutItemInput{
		Item: item,
		TableName: aws.String(TABLE_NAME),
	}
	_, err = svc.PutItem(input)
	if err != nil {
		return "", err
	}
	return "Item saved", nil
}




func main () {
	lambda.Start(saveHello)
	
}
