package helper

import (
	"fmt"

	"github.com/kaankoken/helper/pkg"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

// DebugModule -> Dependency Injection for Debug logger
var DebugModule = fx.Options(fx.Provide(SetLoggerFormat))

// DLogger -> Dependency Injection Data Model for Debug logger
type DLogger struct {
	Tag string
}

/*
SetLoggerFormat -> Debug logger formatter initialization

[tag] -> takes {Tag} argument as string to idenfity or differentiate the logs

[return] -> DLogger with initialized logger
*/
func SetLoggerFormat(tag *pkg.Tag) *DLogger {
	log.SetFormatter(&log.TextFormatter{})

	return &DLogger{Tag: tag.T}
}

/*
Error -> Debug logger error logger without callback

	-> Checks whether error {nil} or {not}

[err] -> take parameter as error

[return] -> returns Tag with {error} or {nil} if error does not {exist}
*/
func (logger DLogger) Error(err error) error {
	if err != nil {
		log.Error(logger.Tag, err.Error())

		return fmt.Errorf("%s", logger.Tag+err.Error())
	}

	return nil
}

/*
ErrorWithCallback -> Debug logger error logger with callback

	-> Checks whether error {nil} or {not}

[err] -> take parameter as error
[f] -> callback method needs to be called if error {exist}

[return] -> returns logger.Tag with {error} or {nil} if error does not {exist}
*/
func (logger DLogger) ErrorWithCallback(err error, f func()) error {
	if err != nil {
		f()
		log.Error(logger.Tag, err.Error())

		return fmt.Errorf("%s", logger.Tag+err.Error())
	}

	return nil
}

/*
Info -> Debug logger info logger without callback

[msg] -> take string message as parameter

[return] -> returns logger.Tag with {msg}
*/
func (logger DLogger) Info(msg string) string {
	log.Info(logger.Tag, msg)

	return fmt.Sprintf(logger.Tag + msg)
}

/*
InfoWithCallback -> Debug logger info logger without callback

[msg] -> take string message as parameter
[f] -> callback method needs to be called

[return] -> returns logger.Tag with {msg}
*/
func (logger DLogger) InfoWithCallback(msg string, f func()) string {
	f()
	log.Info(logger.Tag, msg)

	return fmt.Sprintf(logger.Tag + msg)
}
