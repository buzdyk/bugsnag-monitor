package trend

import (
	"github.com/buzdyk/bugsnag-monitor/bugsnag"
	"time"
)

// Private function that calculates trend differences based on provided start and end dates.
func getTrend(trends []bugsnag.Trend, period time.Duration) int {
	var lastPeriodCount, currentPeriodCount int
	now := time.Now()
	date1 := now.Add(-1 * period)
	date2 := now.Add(-2 * period)

	for _, trend := range trends {
		// Calculate counts for last week and this week based on 'From' date range
		if trend.From.After(date1) && trend.From.Before(now) {
			currentPeriodCount += trend.EventsCount
		} else if trend.From.After(date2) && trend.From.Before(date1) {
			lastPeriodCount += trend.EventsCount
		}
	}

	return currentPeriodCount - lastPeriodCount
}

func OneHour(trends []bugsnag.Trend) int {
	return getTrend(trends, time.Hour)
}
func ThreeHours(trends []bugsnag.Trend) int {
	return getTrend(trends, time.Hour*3)
}
func SixHours(trends []bugsnag.Trend) int {
	return getTrend(trends, time.Hour*6)
}
func TwelveHours(trends []bugsnag.Trend) int {
	return getTrend(trends, time.Hour*12)
}
func OneDay(trends []bugsnag.Trend) int {
	return getTrend(trends, time.Hour*24)
}
func OneWeek(trends []bugsnag.Trend) int {
	return getTrend(trends, time.Hour*24*7)
}
func TwoWeeks(trends []bugsnag.Trend) int { return getTrend(trends, time.Hour*24*7*2) }
func OneMonth(trends []bugsnag.Trend) int {
	return getTrend(trends, time.Hour*24*30)
}
