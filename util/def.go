package util

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	Config *viper.Viper
	Logger *zap.Logger
	P      Parms
)

type Parms struct {
	Conf    string
	Help    bool
	LogFile string
}

const KeyRequestId = "requestId"
