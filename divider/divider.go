package divider

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func Div(a, b int) (int, error) {
	logger, _ := zap.NewDevelopment()
	if b == 0 {
		logger.Error("cannot divide by zero")
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}
