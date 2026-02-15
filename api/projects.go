package api

import (
	"encoding/json"
	"io"
)

type project struct {
	ID             string `json:"id"`
	CanAssignTasks bool   `json:"can_assign_tasks"`
	ChildOrder     int    `json:"child_order"`
	Color          string `json:"color"`
	CreatorUID     string `json:"creator_uid"`
	CreatedAt      string `json:"created_at"`
	IsArchived     bool   `json:"is_archived"`
	IsDeleted      bool   `json:"is_deleted"`
	IsFavorite     bool   `json:"is_favorite"`
	IsFrozen       bool   `json:"is_frozen"`
	Name           string `json:"name"`
	UpdatedAt      string `json:"updated_at"`
	ViewStyle      string `json:"view_style"`
	DefaultOrder   int    `json:"default_order"`
	Description    string `json:"description"`
	PublicKey      string `json:"public_key"`
	// access
	Role         string `json:"role"`
	ParentID     string `json:"parent_id"`
	InboxProject bool   `json:"inbox_project"`
	IsCollapsed  bool   `json:"is_collapsed"`
	IsShared     bool   `json:"is_shared"`
}

// /api/v1/projects
type getProjectsResponse struct {
	Results    []project `json:"results"`
	NextCursor string    `json:"next_cursor"`
}

func (c *Client) Projects() (*getProjectsResponse, error) {

	url, err := c.baseURL.Parse("/api/v1/projects")
	if err != nil {
		return nil, err
	}

	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var projects *getProjectsResponse
	err = json.Unmarshal(body, &projects)
	if err != nil {
		return nil, err
	}

	return projects, nil

}
