package helper

import (
	"fmt"
	"os"

	"github.com/kaankoken/helper/pkg"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ReleaseModule -> Dependency Injection for Release logger
var ReleaseModule = fx.Options(
	fx.Provide(InitializeLogger, InitializeLoggerPtr),
	fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
		return &RLogger{Logger: logger}
	}),
)

// RLogger -> Dependency Injection Data Model for Release logger
type RLogger struct {
	Logger *zap.Logger
	Tag    string
}

/*
InitializeLogger -> Initialize zap logger for release mode

	-> Generate zap logger for both {command line} & logger to {text file}

[return] -> return zap.logger
*/
func InitializeLogger() *zap.Logger {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	logFile, _ := os.OpenFile("text.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	writer := zapcore.AddSync(logFile)
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)

	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

/*
InitializeLoggerPtr -> Generate RLogger

[logger] -> takes argument as zap.logger
[tag] -> takes {Tag} argument as string to idenfity or differentiate the logs

[return] -> returns {Rlogger} that generated with {zap.logger}
*/
func InitializeLoggerPtr(logger *zap.Logger, tag *pkg.Tag) *RLogger {
	return &RLogger{Logger: logger, Tag: tag.T}
}

/*
Error -> Release logger error logger without callback

	-> Checks whether error {nil} or {not}

[err] -> take parameter as error

[return] -> returns logger.Tag with {error} or {nil} if error does not {exist}
*/
func (logger RLogger) Error(err error) error {
	if err != nil {
		logger.Logger.Error(logger.Tag + err.Error())

		return fmt.Errorf(logger.Tag + err.Error())
	}

	return nil
}

/*
ErrorWithCallback -> Release logger error logger with callback

	-> Checks whether error {nil} or {not}

[err] -> take parameter as error
[f] -> callback method needs to be called if error {exist}

[return] -> returns logger.Tag with {error} or {nil} if error does not {exist}
*/
func (logger RLogger) ErrorWithCallback(err error, f func()) error {
	if err != nil {
		f()
		logger.Logger.Error(logger.Tag + err.Error())

		return fmt.Errorf(logger.Tag + err.Error())
	}

	return nil
}

/*
Info -> Release logger info logger without callback

[msg] -> take string message as parameter

[return] -> returns logger.Tag with {msg}
*/
func (logger RLogger) Info(msg string) string {
	logger.Logger.Info(logger.Tag + msg)

	return fmt.Sprintf(logger.Tag + msg)
}

/*
InfoWithCallback -> Release logger info logger without callback

[msg] -> take string message as parameter
[f] -> callback method needs to be called

[return] -> returns logger.Tag with {msg}
*/
func (logger *RLogger) InfoWithCallback(msg string, f func()) string {
	f()
	logger.Logger.Info(logger.Tag + msg)

	return fmt.Sprintf(logger.Tag + msg)
}

/*
LogEvent -> Release logger for logging fx event

[event] -> take argument as fx.event
*/
func (logger RLogger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		logger.Logger.Debug("OnStart hook executing: ",
			zap.String("callee", e.FunctionName),
			zap.String("caller", e.CallerName),
		)
	case *fxevent.OnStartExecuted:
		if e.Err != nil {
			logger.Logger.Debug("OnStart hook failed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.Error(e.Err),
			)
		} else {
			logger.Logger.Debug("OnStart hook executed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.String("runtime", e.Runtime.String()),
			)
		}
	case *fxevent.OnStopExecuting:
		logger.Logger.Debug("OnStop hook executing: ",
			zap.String("callee", e.FunctionName),
			zap.String("caller", e.CallerName),
		)
	case *fxevent.OnStopExecuted:
		if e.Err != nil {
			logger.Logger.Debug("OnStop hook failed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.Error(e.Err),
			)
		} else {
			logger.Logger.Debug("OnStop hook executed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.String("runtime", e.Runtime.String()),
			)
		}
	case *fxevent.Supplied:
		logger.Logger.Debug("supplied: ", zap.String("type", e.TypeName), zap.Error(e.Err))
	case *fxevent.Provided:
		for _, rtype := range e.OutputTypeNames {
			logger.Logger.Debug("provided: " + e.ConstructorName + " => " + rtype)
		}
	case *fxevent.Decorated:
		for _, rtype := range e.OutputTypeNames {
			logger.Logger.Debug("decorated: ",
				zap.String("decorator", e.DecoratorName),
				zap.String("type", rtype),
			)
		}
	case *fxevent.Invoking:
		logger.Logger.Debug("invoking: " + e.FunctionName)
	case *fxevent.Started:
		if e.Err == nil {
			logger.Logger.Debug("started")
		}
	case *fxevent.LoggerInitialized:
		if e.Err == nil {
			logger.Logger.Debug("initialized: custom fxevent.Logger -> " + e.ConstructorName)
		}
	}
}
