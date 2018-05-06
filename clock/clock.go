package clock

import (
	"fmt"
)

type Clock struct {
	hour, minute int
}

func (clock Clock) String() string {
	return fmt.Sprintf("%.2d:%.2d", clock.hour, clock.minute)
}

func (clock Clock) Add(minutes int) Clock {
	hours, minutes := divmod(minutes, 60)

	clock.hour += hours
	clock.minute += minutes

	if clock.minute < 0 {
		clockHours, clockMinutes := divmod(clock.minute, 60)
		clock.minute = 60 + clockMinutes
		clock.hour -= clockHours + 1
	}
	if clock.hour < 0 {
		clock.hour = 24 + clock.hour%24
	}

	clock.hour += clock.minute / 60
	clock.minute %= 60
	clock.hour %= 24

	return clock
}

func (clock Clock) Subtract(minutes int) Clock {
	return clock.Add(-minutes)
}

func New(hour, minute int) Clock {
	return Clock{0, 0}.Add(hour*60 + minute)
}

func divmod(dividend, divisor int) (quotient, remainder int) {
	quotient = dividend / divisor
	remainder = dividend % divisor
	return
}
