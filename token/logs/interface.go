package logs

type ILogger interface {
	IEntry

	// WithField creates an entry from the standard logger and adds a field to
	// it. If you want multiple fields, use `WithFields`.
	//
	// Note that it doesn't log until you call Debug, Info, Warn, Fatal
	// or Panic on the Entry it returns.
	WithField(key string, value interface{}) ILogger

	// WithFields creates an entry from the standard logger and adds multiple
	// fields to it. This is simply a helper for `WithField`, invoking it
	// once for each field.
	//
	// Note that it doesn't log until you call Debug, Info, Warn, Fatal
	// or Panic on the Entry it returns.
	WithFields(fields map[string]interface{}) ILogger

	// WithError creates an entry from the standard logger and adds an error to it,
	// using the value defined in ErrorKey as key.
	WithError(err error) ILogger

	// AddField add field to existed logger
	AddField(key string, value interface{})

	// AddField add field to existed logger
	AddFields(fields map[string]interface{})
}

type IEntry interface {
	Trace(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})

	Tracef(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
}
