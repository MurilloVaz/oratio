package logger

import (
	"os"

	"github.com/MurilloVaz/oratio/extensions/configuration"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

var standardFields log.Fields

func getHostname() string {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return name
}

func InvalidArg(expected interface{}, got interface{}) {
	log.WithFields(standardFields).WithFields(log.Fields{"expected": expected, "got": got}).Info("Invalid argument")
}

func ArgNil(argName string) {
	log.WithFields(standardFields).WithFields(log.Fields{"argName": argName}).Info("Argument is nil")
}

func Unexpected(err error) {
	log.WithFields(standardFields).WithFields(log.Fields{"stacktrace": errors.Wrap(err, "").Error()}).Error("Unexpected has ocurred")
}

func Database(err error) {
	log.WithFields(standardFields).WithFields(log.Fields{"stacktrace": errors.Wrap(err, "").Error()}).Warning("Command failed in database")
}

func Unreachable(resource string, err error) {
	log.WithFields(standardFields).WithFields(log.Fields{"stacktrace": errors.Wrap(err, "").Error(), "resource": resource}).Warning("Cannot reach resource")
}

func SetLogger() {
	standardFields = log.Fields{
		"hostname": getHostname(),
		"appname":  configuration.Global.AppName,
	}
}
