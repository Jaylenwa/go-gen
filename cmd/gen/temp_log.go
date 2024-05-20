package gen

const TempLog = `
package log

import "github.com/sirupsen/logrus"

func InitLog() *logrus.Logger {
	return logrus.New()
}
`
