package cwlogfmt_test

import (
	"bytes"
	"testing"

	"github.com/ravernkoh/cwlogfmt"
	"github.com/sirupsen/logrus"
)

func TestCloudWatchLogsFormatter_Format(t *testing.T) {
	tests := []struct {
		Formatter *cwlogfmt.CloudWatchLogsFormatter
		Fields    logrus.Fields
		Message   string
		Level     logrus.Level
		Expected  string
	}{
		{
			Formatter: &cwlogfmt.CloudWatchLogsFormatter{},
			Fields: logrus.Fields{
				"BoolField": true,
				"IntField":  1,
			},
			Message:  "Message",
			Level:    logrus.DebugLevel,
			Expected: "DEBUG Message: Message BoolField: true IntField: 1 \n",
		},
	}

	for i, test := range tests {
		b := new(bytes.Buffer)
		log := &logrus.Logger{
			Formatter: test.Formatter,
			Out:       b,
			Level:     logrus.DebugLevel,
		}
		entry := log.WithFields(test.Fields)

		// Don't perform fatal or panic
		switch test.Level {
		case logrus.DebugLevel:
			entry.Debug(test.Message)
		case logrus.InfoLevel:
			entry.Info(test.Message)
		case logrus.WarnLevel:
			entry.Warn(test.Message)
		case logrus.ErrorLevel:
			entry.Error(test.Message)
		}

		out := b.String()
		if out != test.Expected {
			t.Errorf("Test %d: Expected output to be %v, got %v", i+1, test.Expected, out)
		}
	}
}
