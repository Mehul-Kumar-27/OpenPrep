package config

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/mattn/go-colorable"
)

var Log *logrus.Logger

type Fields logrus.Fields

// CustomFormatter is a custom log formatter for Logrus that includes colored output
type CustomFormatter struct {
	logrus.TextFormatter
}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Use JSON formatter for the actual log entry
	formatter := &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z07:00",
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := filepath.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
		PrettyPrint: true,
	}
	jsonBytes, err := formatter.Format(entry)
	if err != nil {
		return nil, err
	}

	// Apply color for console output
	if entry.Logger.Out == os.Stdout || entry.Logger.Out == os.Stderr {
		return f.TextFormatter.Format(entry)
	}

	return jsonBytes, nil
}

// InitLogger initializes the logger with the specified configuration
func InitLogger(logLevel string, logFile string, enableConsole bool) {
	Log = logrus.New()

	// Set formatter
	Log.SetFormatter(&CustomFormatter{
		TextFormatter: logrus.TextFormatter{
			TimestampFormat: "2006-01-02T15:04:05.000Z07:00",
			FullTimestamp:   true,
			ForceColors:     true,
			DisableColors:  false,
		},
	})

	Log.Info("Logger initialized... ")
	// Enable caller information
	Log.SetReportCaller(true)

	// Set output
	var output io.Writer
	if enableConsole && logFile != "" {
		output = io.MultiWriter(colorable.NewColorableStdout(), openLogFile(logFile))
	} else if enableConsole {
		output = colorable.NewColorableStdout()
	} else if logFile != "" {
		output = openLogFile(logFile)
	} else {
		output = io.Discard
	}
	Log.SetOutput(output)

	// Set Gin mode
	if Log.GetLevel() == logrus.DebugLevel {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

// openLogFile opens or creates a log file
func openLogFile(logFile string) *os.File {
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file: %v", err))
	}
	return file
}

// LogMiddleware creates a Gin middleware for logging HTTP requests
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Log request details
		param := gin.LogFormatterParams{
			Request: c.Request,
			Keys:    c.Keys,
		}

		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()

		param.BodySize = c.Writer.Size()

		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path

		Log.WithFields(logrus.Fields{
			"clientIP":  param.ClientIP,
			"method":    param.Method,
			"path":      param.Path,
			"status":    param.StatusCode,
			"latency":   param.Latency,
			"bodySize":  param.BodySize,
			"userAgent": c.Request.UserAgent(),
		}).Info("HTTP request")
	}
}

// Convenience functions for logging
func Info(args ...interface{}) {
	Log.Info(args...)
}

func Infof(format string, args ...interface{}) {
	Log.Infof(format, args...)
}

func Warn(args ...interface{}) {
	Log.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	Log.Warnf(format, args...)
}

func Error(args ...interface{}) {
	Log.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	Log.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	Log.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	Log.Fatalf(format, args...)
}

func WithFields(fields Fields) *logrus.Entry {
	return Log.WithFields(logrus.Fields(fields))
}
