package repositories_search

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	coursesDomain "search/domain_search"
)

type HTTPConfig struct {
	Host string
	Port string
}

type HTTP struct {
	baseURL func(courseID string) string
}

func NewHTTP(config HTTPConfig) HTTP {
	return HTTP{
		baseURL: func(courseID string) string {
			return fmt.Sprintf("http://%s:%s/courses/%s", config.Host, config.Port, courseID)
		},
	}
}

func (repository HTTP) GetCourseByID(ctx context.Context, id string) (coursesDomain.CourseDto, error) {
	resp, err := http.Get(repository.baseURL(id))
	if err != nil {
		return coursesDomain.CourseDto{}, fmt.Errorf("Error fetching course (%s): %w\n", id, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return coursesDomain.CourseDto{}, fmt.Errorf("Failed to fetch course (%s): received status code %d\n", id, resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return coursesDomain.CourseDto{}, fmt.Errorf("Error reading response body for course (%s): %w\n", id, err)
	}

	// Unmarshal the course details into the course struct
	var course coursesDomain.CourseDto
	if err := json.Unmarshal(body, &course); err != nil {
		return coursesDomain.CourseDto{}, fmt.Errorf("Error unmarshaling course data (%s): %w\n", id, err)
	}

	return course, nil
}
