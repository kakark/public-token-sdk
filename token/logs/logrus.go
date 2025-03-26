package logs

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type logrusWrapper struct {
	*logrus.Entry
}

func init() {
	logrus.SetReportCaller(true)
}

func (l *logrusWrapper) WithField(key string, value interface{}) ILogger {
	return &logrusWrapper{Entry: l.Entry.WithField(key, value)}
}

func (l *logrusWrapper) WithFields(fields map[string]interface{}) ILogger {
	strFields := make(map[string]interface{})
	for k, v := range fields {
		strFields[k] = fmt.Sprintf("%v", v)
	}
	return &logrusWrapper{Entry: l.Entry.WithFields(strFields)}
}

func (l *logrusWrapper) AddField(key string, value interface{}) {
	l.Entry = l.Entry.WithField(key, value)
}

func (l *logrusWrapper) AddFields(fields map[string]interface{}) {
	strFields := make(map[string]interface{})
	for k, v := range fields {
		strFields[k] = fmt.Sprintf("%v", v)
	}
	l.Entry= l.Entry.WithFields(strFields)
}

func (l *logrusWrapper) WithError(err error) ILogger {
	return &logrusWrapper{Entry: l.Entry.WithError(err)}
}
