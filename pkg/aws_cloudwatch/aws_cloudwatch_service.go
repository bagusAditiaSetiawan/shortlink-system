package aws_cloudwatch

type AwsCloudWatchService interface {
	SendLogInfo(a ...interface{}) bool
	SendLog(flag string, a ...interface{}) bool
}
