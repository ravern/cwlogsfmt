package cwlogfmt

import "github.com/sirupsen/logrus"

// CloudWatchLogsFormatter provides formatting similar to the default
// CloudWatch Logs format (from AWS Lambda).
//
//     // Examples of the format
//     START RequestId: 66389135-fd00-11e7-a1f9-8945479469b0 Version: $LATEST
//     PANIC Message: A serious crash
//     FATAL Message: A crash
//     ERROR Message: An error
//     WARN Message: A warning
//     INFO Message: Some information
//     DEBUG Message: Some more information
//     END RequestId: 66389135-fd00-11e7-a1f9-8945479469b0
//     REPORT RequestId: 66389135-fd00-11e7-a1f9-8945479469b0 Duration: 0.96 ms Billed Duration: 100 ms Memory Size: 128 MB Max Memory Used: 24 MB
//
// It does not log entry.Buffer.
type CloudWatchLogsFormatter struct{}

// Format formats the given log entry.
func (f *CloudWatchLogsFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return nil, nil
}
