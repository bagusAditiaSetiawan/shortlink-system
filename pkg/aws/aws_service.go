package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
)

func NewAwsSessionService() *session.Session {
	region := aws.String(os.Getenv("AWS_S3_REGION"))
	secret := os.Getenv("AWS_S3_SECRET")
	key := os.Getenv("AWS_S3_KEY")
	cred := credentials.NewStaticCredentials(key, secret, "")
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      region,
		Credentials: cred,
	}))
	return sess
}
