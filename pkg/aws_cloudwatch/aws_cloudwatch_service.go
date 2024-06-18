package aws_cloudwatch

type AwsCloudWatchService interface {
	SendLogInfo(a ...interface{}) bool
	sendLog(flag string, a ...interface{}) bool
}
