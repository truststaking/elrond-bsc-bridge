package testHelpers

import (
	"time"

	logger "github.com/ElrondNetwork/elrond-go-logger"
)

func SetTestLogLevel() {
	_ = logger.SetLogLevel("*:" + logger.LogNone.String())
}

type TimerStub struct {
	TimeNowUnix int64
	WasStarted  bool
}

func (s *TimerStub) After(time.Duration) <-chan time.Time {
	return time.After(0 * time.Millisecond)
}

func (s *TimerStub) NowUnix() int64 {
	return s.TimeNowUnix
}

func (s *TimerStub) Start() {
	s.WasStarted = true
}

func (s *TimerStub) Close() error {
	return nil
}
