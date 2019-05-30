package hook

import "github.com/sirupsen/logrus"

// NoopFormatter is a no-op logrus formatter
type NoopFormatter struct{}

// Format implements logrus formatter interface
func (*NoopFormatter) Format(*logrus.Entry) ([]byte, error) {
	return nil, nil
}
