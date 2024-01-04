package sna

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// SignIn - Get a new token for user
func (c *Client) SignIn() (*AuthResponse, error) {
	if c.Auth.Username == "" || c.Auth.Password == "" {
		return nil, fmt.Errorf("define username and password")
	}
	authString := "username=" + c.Auth.Username + "&password=" + c.Auth.Password

	req, err := http.NewRequest("POST", getTokenUrl(c), strings.NewReader(authString))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}
	ar := AuthResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}

// SignOut - Revoke the token for a user
func (c *Client) SignOut(authToken *string) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/signout", c.HostURL), strings.NewReader(string("")))
	if err != nil {
		return err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return err
	}

	if string(body) != "Signed out user" {
		return errors.New(string(body))
	}

	return nil
}
