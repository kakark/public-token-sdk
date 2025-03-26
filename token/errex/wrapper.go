package errex

import (
	"github.com/pkg/errors"
)

// WrapErr 使用第二个 error 参数作为附加消息使用
func WrapErr(cause error, err error) error {
	var msg string
	if err != nil {
		msg = err.Error()
	}
	return errors.Wrap(cause, msg)
}

// Wrap 使用第二个参数作为附加消息使用
func Wrap(cause error, msg string) error {
	return errors.Wrap(cause, msg)
}

// Wrapf 使用第二个参数以及后续参数作为附加消息使用
func Wrapf(err error, format string, args ...interface{}) error {
	return errors.Wrapf(err, format, args...)
}

// WithStack 在当前 err 的基础上添加调用栈信息
func WithStack(err error) error {
	return errors.WithStack(err)
}
