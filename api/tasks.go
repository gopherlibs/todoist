package api

import (
	"encoding/json"
	"fmt"
	"io"
)

type Task struct {
	UserID         string   `json:"user_id"`
	ID             string   `json:"id"`
	ProjectID      string   `json:"project_id"`
	SectionID      string   `json:"section_id"`
	ParentID       string   `json:"parent_id"`
	AddedByUID     string   `json:"added_by_uid"`
	AssignedByUID  string   `json:"assigned_by_uid"`
	ResponsibleUID string   `json:"responsible_uid"`
	Labels         []string `json:"labels"`
	Deadline       struct {
		Property1 string `json:"property1"`
		Property2 string `json:"property2"`
	} `json:"deadline"`
	Duration struct {
		Property1 int `json:"property1"`
		Property2 int `json:"property2"`
	} `json:"duration"`
	Checked        bool   `json:"checked"`
	IsDeleted      bool   `json:"is_deleted"`
	AddedAt        string `json:"added_at"`
	CompletedAt    string `json:"completed_at"`
	CompletedByUID string `json:"completed_by_uid"`
	UpdatedAt      string `json:"updated_at"`
	// Due - I'm not sure how this works yet
	Priority    int    `json:"priority"`
	ChildOrder  int    `json:"child_order"`
	Content     string `json:"content"`
	Description string `json:"description"`
	NoteCount   int    `json:"note_count"`
	DayOrder    int    `json:"day_order"`
	IsCollapsed bool   `json:"is_collapsed"`
}

// /api/v1/tasks
type getTasksResponse struct {
	Results    []Task `json:"results"`
	NextCursor string `json:"next_cursor"`
}

func (c *client) Tasks(projectID string) (*getTasksResponse, error) {

	u, err := c.baseURL.Parse("/api/v1/tasks?limit=200")
	if err != nil {
		return nil, err
	}

	params := u.Query()

	// Optionally add a project ID.
	if projectID != "" {
		params.Set("project_id", projectID)
	}

	u.RawQuery = params.Encode()

	resp, err := c.get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tasks *getTasksResponse
	err = json.Unmarshal(body, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (c *client) TaskClose(id string) (int, error) {

	url, err := c.baseURL.Parse(fmt.Sprintf("/api/v1/tasks/%s/close", id))
	if err != nil {
		return 400, err
	}

	resp, err := c.post(url, nil)
	if err != nil {
		return 400, err
	}

	return resp.StatusCode, nil
}

func (c *client) TaskDelete(id string) (int, error) {

	url, err := c.baseURL.Parse(fmt.Sprintf("/api/v1/tasks/%s", id))
	if err != nil {
		return 400, err
	}

	resp, err := c.delete(url)
	if err != nil {
		return 400, err
	}

	return resp.StatusCode, nil
}
