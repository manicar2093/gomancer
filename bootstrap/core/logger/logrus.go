package logger

import (
	"fmt"
	"github.com/manicar2093/gomancer/bootstrap/core/stages"
	"io"
	internalLog "log"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	createBlackBox = func(stage string) (*os.File, error) {
		file, err := os.OpenFile(fmt.Sprintf("black_box_%s.log", stage), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		return file, err
	}
	initializers = map[string]func(l *logrus.Logger) error{
		stages.Test: func(l *logrus.Logger) error {
			file, err := createBlackBox(stages.Test)
			if err != nil {
				return err
			}
			loggerOutput := io.MultiWriter(file, os.Stdout)

			l.SetOutput(loggerOutput)
			internalLog.SetOutput(loggerOutput)

			l.SetFormatter(&logrus.JSONFormatter{
				PrettyPrint: true,
			})
			l.SetReportCaller(true)
			l.SetLevel(logrus.DebugLevel)
			return nil
		},
		stages.Dev: func(l *logrus.Logger) error {
			file, err := createBlackBox(stages.Dev)
			if err != nil {
				return err
			}
			loggerOutput := io.MultiWriter(file, os.Stdout)
			l.SetOutput(loggerOutput)
			internalLog.SetOutput(loggerOutput)
			l.SetFormatter(&logrus.JSONFormatter{
				PrettyPrint: true,
			})
			l.SetReportCaller(true)
			l.SetLevel(logrus.DebugLevel)
			return nil
		},
		stages.Prod: func(l *logrus.Logger) error {
			l.SetOutput(os.Stdout)
			internalLog.SetOutput(os.Stdout)
			l.SetFormatter(&logrus.JSONFormatter{})
			l.SetReportCaller(true)
			return nil
		},
	}
)

// Config check configured environment to set configurations to logrus.
// Depending on environment outputs will be (considering output from golang log package will be catched to).
// A log file called black_box_<environment>.log will be created
func Config() {
	logger := logrus.StandardLogger()
	currentStage, err := stages.GetCurrentStage()
	if err != nil {
		panic(err)
	}
	if err := initializers[currentStage](logger); err != nil {
		panic(err)
	}
}
