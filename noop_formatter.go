package hook

import "github.com/sirupsen/logrus"

type NoopFormatter struct{}

func (*NoopFormatter) Format(*logrus.Entry) ([]byte, error) {
	return nil, nil
}
