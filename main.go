package bugsnag

import (
	"fmt"
	"github.com/buzdyk/bugsnag-monitor/bugsnag"
	"github.com/buzdyk/bugsnag-monitor/trend"
	"os"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide organizationID")
		return
	}

	// Retrieve the first argument (excluding the program name)
	organizationID := os.Args[1]

	projects, err := bugsnag.GetProjects(organizationID)
	if err != nil {
		panic(err)
	}
	sum := 0
	trendval := 0

	var wg sync.WaitGroup

	for _, project := range projects {
		wg.Add(1)

		go func(projectID string) {
			defer wg.Done()
			trends, err := bugsnag.GetTrends(projectID, "30d", "2h")
			if err != nil {
				panic(err)
			}

			for _, item := range trends {
				sum += item.EventsCount
			}

			trendval += trend.TwoWeeks(trends)
		}(project.ID)
	}

	wg.Wait()

	fmt.Println("Events over last 30 days:", sum)
	fmt.Println("Trend over 2 last weeks:", trendval, " events")
}
