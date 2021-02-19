package s3

import (
	"bytes"
	"github.com/ThomasVNN/go-base/storage/local"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	awssession "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"io/ioutil"
	logg "log"
	"net/http"
)
func AddFileToS3(fileName string,files io.Reader) error {
	awsss, err := createAWSSession()
	logg.Print("\nAddFileToS3:",fileName)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(files)
	if err != nil {
		return  err
	}
	_, err = s3.New(awsss).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String("id-card-vib"),
		Key:                  aws.String(fileName),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(data),
		ContentLength:        aws.Int64(int64(len(data))),
		ContentType:          aws.String(http.DetectContentType(data)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	return nil
}
func createAWSSession() (*awssession.Session, error) {
	if local.Getenv("ENVIRONMENT") != "dev" {
		conf := aws.Config{
			Region:      aws.String(local.Getenv("AWS_REGION")),
		}
		return awssession.NewSession(&conf)
	}
	key := local.Getenv("aws_access_key_id")
	secret := local.Getenv("aws_secret_access_key")
	token := local.Getenv("aws_session_token")

	conf := aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(key, secret, token),
	}
	return awssession.NewSession(&conf)
}
