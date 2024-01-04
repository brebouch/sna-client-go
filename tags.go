package sna

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// GetTags - Returns list of tags
func (c *Client) GetTags() ([]TagList, error) {
	req, err := http.NewRequest("GET", getUrl(c, "/tags"), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	tagData := TagListData{}
	err = json.Unmarshal(body, &tagData)
	if err != nil {
		return nil, err
	}

	return tagData.Data, nil
}

// GetAllTags - Returns all user's tag
func (c *Client) GetAllTags(authToken *string) (*[]TagList, error) {
	req, err := http.NewRequest("GET", getUrl(c, "/tags"), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	tagData := TagListData{}
	err = json.Unmarshal(body, &tagData)
	if err != nil {
		return nil, err
	}
	tags := tagData.Data

	return &tags, nil
}

// GetTag - Returns specific tag (no auth required)
func (c *Client) GetTag(tagID string) (*Tag, error) {
	req, err := http.NewRequest("GET", getUrl(c, fmt.Sprintf("/tags/%s", tagID)), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	tag := TagData{}
	err = json.Unmarshal(body, &tag)
	if err != nil {
		return nil, err
	}

	return &tag.Data, nil
}

// CreateTag - Create new tag
func (c *Client) CreateTag(tag Tag, authToken *string) (*Tag, error) {
	rb, err := json.Marshal(tag)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", getUrl(c, "/tags"), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	newTag := TagData{}
	err = json.Unmarshal(body, &newTag)
	if err != nil {
		return nil, err
	}

	return &newTag.Data, nil
}

// UpdateTag - Updates an tag
func (c *Client) UpdateTag(tagID string, tagItems []TagData, authToken *string) (*Tag, error) {
	rb, err := json.Marshal(tagItems)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", getUrl(c, fmt.Sprintf("/tags/%s", tagID)), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	tag := TagData{}
	err = json.Unmarshal(body, &tag)
	if err != nil {
		return nil, err
	}

	return &tag.Data, nil
}

// DeleteTag - Deletes an tag
func (c *Client) DeleteTag(tagID string, authToken *string) error {
	req, err := http.NewRequest("DELETE", getUrl(c, fmt.Sprintf("/tags/%s", tagID)), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return err
	}

	if string(body) != "" {
		return errors.New(string(body))
	}

	return nil
}
