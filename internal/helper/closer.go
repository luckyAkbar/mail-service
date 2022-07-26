package helper

import "github.com/sirupsen/logrus"

func WrapCloser(close func() error) {
	if err := close(); err != nil {
		logrus.Error(err)
	}
}
