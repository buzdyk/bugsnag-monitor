package bugsnag

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Project struct {
	ID                           string    `json:"id"`
	OrganizationID               string    `json:"organization_id"`
	Slug                         string    `json:"slug"`
	Name                         string    `json:"name"`
	APIKey                       string    `json:"api_key"`
	Type                         string    `json:"type"`
	IsFullView                   bool      `json:"is_full_view"`
	ReleaseStages                []string  `json:"release_stages"`
	Language                     string    `json:"language"`
	CreatedAt                    time.Time `json:"created_at"`
	UpdatedAt                    time.Time `json:"updated_at"`
	ErrorsURL                    string    `json:"errors_url"`
	EventsURL                    string    `json:"events_url"`
	URL                          string    `json:"url"`
	HTMLURL                      string    `json:"html_url"`
	OpenErrorCount               int       `json:"open_error_count"`
	ForReviewErrorCount          int       `json:"for_review_error_count"`
	CollaboratorsCount           int       `json:"collaborators_count"`
	TeamsCount                   int       `json:"teams_count"`
	GlobalGrouping               []string  `json:"global_grouping"`
	LocationGrouping             []string  `json:"location_grouping"`
	DiscardedAppVersions         []string  `json:"discarded_app_versions"`
	DiscardedErrors              []string  `json:"discarded_errors"`
	CustomEventFieldsUsed        int       `json:"custom_event_fields_used"`
	ResolveOnDeploy              bool      `json:"resolve_on_deploy"`
	PerformanceDisplayType       string    `json:"performance_display_type"`
	DefaultPerformancePercentile string    `json:"default_performance_percentile"`
}

type Trend struct {
	From        time.Time `json:"from"`
	To          time.Time `json:"to"`
	EventsCount int       `json:"events_count"`
}

func GetProjects(organizationID string) ([]Project, error) {
	endpoint := fmt.Sprintf("/organizations/%s/projects", organizationID)
	body, err := request(endpoint)

	if err != nil {
		return nil, err
	}

	var result []Project
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New(fmt.Sprintf("Error unmarshaling response: %s", err))
	}

	return result, nil
}

func GetTrends(projectID string, since string, resolution string) ([]Trend, error) {
	endpoint := fmt.Sprintf("/projects/%s/trend?resolution=%s&filters[event.since][][value]=%s&filters[event.since][][type]=eq", projectID, resolution, since)
	body, err := request(endpoint)

	if err != nil {
		return nil, err
	}

	var result []Trend
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New(fmt.Sprintf("Error unmarshaling response: %s", err))
	}

	return result, nil
}

func request(endpoint string) ([]byte, error) {
	authToken := "xxx"

	url := "https://api.bugsnag.com" + endpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error creating request:%s", err))
	}

	req.Header.Add("Authorization", "token "+authToken)
	req.Header.Add("X-Version", "2")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error making request:%s", err))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Unexpected status code: %d", resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error reading response body: %s", err))
	}

	return body, nil
}
