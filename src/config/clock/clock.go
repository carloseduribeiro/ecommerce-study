package clock

import (
	"github.com/benbjohnson/clock"
)

func init() {
	SetSystemClock()
}

var Time clock.Clock

func SetSystemClock() {
	Time = clock.New()
}

func GetClockMock() *clock.Mock {
	mock := clock.NewMock()
	Time = mock
	return mock
}
