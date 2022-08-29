package helper

import (
	"go.uber.org/fx"
)

const (
	dev string = "dev"
)

// LoggerModule -> Dependency Injection for General logger module
var LoggerModule = fx.Options(
	DebugModule,
	ReleaseModule,
	fx.Provide(Initialize),
)

// ILogHandler -> General interface that both RLogger & DLogger use
type ILogHandler interface {
	Error(err error) error
	ErrorWithCallback(err error, f func()) error
	Info(msg string) string
	InfoWithCallback(msg string, f func()) string
}

// LogHandler -> Dependency Injection Data Model for LoggerModule needs
type LogHandler struct {
	debug, release *ILogHandler
	flavor         string
}

// Flavor -> Data type for {flavor} that needed by LogHandler
type Flavor struct{ F string }

/*
Initialize -> Initialize General logger

[c] -> Passing config to seperate according to debug or release mode
[D] -> Passing DLogger to generate LogHandler
[R] -> Passing RLogger to generate LogHandler

[return] -> returns Generated LogHandler
*/
func Initialize(f *Flavor, D *DLogger, R *RLogger) *LogHandler {
	var d, r ILogHandler

	d = D
	r = R

	return &LogHandler{debug: &d, release: &r, flavor: f.F}
}

/*
Error -> LogHandler error logger without callback

	-> Checks whether error {nil} or {not}
	-> Calls {logger} according to {flavor}

[err] -> take parameter as error

[return] -> returns tag with {error} or {nil} if error does not {exist}
*/
func (l LogHandler) Error(err error) error {
	if l.flavor == dev {
		return (*l.debug).Error(err)
	}

	return (*l.release).Error(err)
}

/*
ErrorWithCallback -> LogHandler error logger with callback

	-> Checks whether error {nil} or {not}
	-> Calls {logger} according to {flavor}

[err] -> take parameter as error
[f] -> callback method needs to be called if error {exist}

[return] -> returns tag with {error} or {nil} if error does not {exist}
*/
func (l LogHandler) ErrorWithCallback(err error, f func()) error {
	if l.flavor == dev {
		return (*l.debug).ErrorWithCallback(err, f)
	}

	return (*l.release).ErrorWithCallback(err, f)
}

/*
Info -> LogHandler info logger without callback

	-> Calls {logger} according to {flavor}

[msg] -> take string message as parameter

[return] -> returns tag with {msg}
*/
func (l LogHandler) Info(msg string) string {
	if l.flavor == dev {
		return (*l.debug).Info(msg)
	}

	return (*l.release).Info(msg)
}

/*
InfoWithCallback -> LogHandler info logger without callback

	-> Calls {logger} according to {flavor}

[msg] -> take string message as parameter
[f] -> callback method needs to be called

[return] -> returns tag with {msg}
*/
func (l LogHandler) InfoWithCallback(msg string, f func()) string {
	if l.flavor == dev {
		return (*l.debug).InfoWithCallback(msg, f)
	}

	return (*l.release).InfoWithCallback(msg, f)
}
