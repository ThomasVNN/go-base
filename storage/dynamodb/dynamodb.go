package dynamodb


import (
	"github.com/ThomasVNN/go-base/storage/local"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

)

func Connection() (*dynamodb.DynamoDB, error) {
	region := local.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return nil, err
	}
	dynaClient := dynamodb.New(awsSession)

	return dynaClient, nil
}
