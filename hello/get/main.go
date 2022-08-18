package get

import (
	"encoding/json"
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

func getHello (event Object) (string, error) {
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
				S: aws.String(event.pk),
			},
			"sk": {
				S: aws.String(event.sk),
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

	item := dynamodbattribute.UnmarshalMap(result.Item, &object)
	
	data,err := json.Marshal(item)

	if err != nil {
		return "", err
	}

	return string(data), nil

}


func main () {
	lambda.Start(getHello)
}