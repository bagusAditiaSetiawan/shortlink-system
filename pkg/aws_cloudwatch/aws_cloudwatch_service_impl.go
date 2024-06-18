package aws_cloudwatch

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/gofiber/fiber/v2/log"
	"os"
	"shortlink-system/pkg/helper"
	"time"
)

type AwsCloudWatchServiceImpl struct {
	Logging *cloudwatchlogs.CloudWatchLogs
	IsSend  bool
}

func NewCloudWatchLogsService(sess *session.Session) *cloudwatchlogs.CloudWatchLogs {
	return cloudwatchlogs.New(sess)
}

func NewAwsCloudWatchServiceImpl(logging *cloudwatchlogs.CloudWatchLogs, isSend bool) *AwsCloudWatchServiceImpl {
	return &AwsCloudWatchServiceImpl{
		Logging: logging,
		IsSend:  isSend,
	}
}

func (service AwsCloudWatchServiceImpl) sendLog(flag string, a ...interface{}) bool {
	message := fmt.Sprintf("[%s] ", flag)
	for _, item := range a {
		message += fmt.Sprint(item) + " "
	}
	log.Info(message)
	if !service.IsSend {
		return false
	}

	group := os.Getenv("AWS_CLOUDWATCH_GROUP")
	stream := os.Getenv("AWS_CLOUDWATCH_STREAM")

	_, err := service.Logging.PutLogEvents(&cloudwatchlogs.PutLogEventsInput{
		LogGroupName:  aws.String(group),
		LogStreamName: aws.String(stream),
		LogEvents: []*cloudwatchlogs.InputLogEvent{
			{
				Message:   aws.String(message),
				Timestamp: aws.Int64(time.Now().UnixNano() / int64(time.Millisecond)),
			},
		},
	})
	helper.IfErrorHandler(err)
	return true
}

func (service *AwsCloudWatchServiceImpl) SendLogInfo(a ...interface{}) bool {
	return service.sendLog("info", a...)
}
