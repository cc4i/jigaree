package logging

import "github.com/sirupsen/logrus"

var Lx = logrus.New()

func init() {
	Lx.SetFormatter(&logrus.JSONFormatter{})
	Lx.SetReportCaller(true)
	Lx.SetLevel(logrus.DebugLevel)
}
