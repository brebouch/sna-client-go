package sna

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// HostURL - Default Hashicups URL
const HostURL string = "http://localhost:19090"

// Client -
type Client struct {
	HostURL    string
	Tenant     int
	HTTPClient *http.Client
	Token      string
	Auth       AuthStruct
	Cookies    []http.Cookie
}

// AuthStruct -
type AuthStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	TokenTtlInSeconds                  int          `json:"tokenTtlInSeconds"`
	SessionMaxIdleTimeInSeconds        int          `json:"sessionMaxIdleTimeInSeconds"`
	PasswordExpiresAt                  string       `json:"passwordExpiresAt"`
	AuthResponse                       AuthResponse `json:"authResponse"`
	LastSuccessfulLoginDate            string       `json:"lastSuccessfulLoginDate"`
	LastSuccessfulLoginLocation        string       `json:"lastSuccessfulLoginLocation"`
	NumFailedLoginsSinceLastSuccessful int          `json:"numFailedLoginsSinceLastSuccessful"`
	LastLoginInfoPopupEnabled          bool         `json:"lastLoginInfoPopupEnabled"`
	PasswordReset                      bool         `json:"passwordReset"`
}

type AuthRole struct {
	Name   string `json:"name"`
	Source string `json:"source"`
}

// AuthResponse -
type AuthResponse struct {
	UserID   int        `json:"userId"`
	Username string     `json:"username"`
	FullName string     `json:"fullName"`
	Roles    []AuthRole `json:"roles"`
}

func getUrl(c *Client, path string) string {
	u := "https://" + c.HostURL + "/tenants/" + strconv.Itoa(c.Tenant) + path
	return u
}

func getTenantUrl(c *Client) string {
	u := "https://" + c.HostURL + "/tenants"
	return u
}

func getTokenUrl(c *Client) string {
	u := "https://" + c.HostURL + "/token/v2/authenticate"
	return u
}

// NewClient -
func NewClient(host, username, password, tenant *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default Hashicups URL
		HostURL: HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	// If username or password not provided, return empty client
	if username == nil || password == nil {
		return &c, nil
	}

	c.Auth = AuthStruct{
		Username: *username,
		Password: *password,
	}

	_, err := c.SignIn()
	if err != nil {
		return nil, err
	}

	t, e := c.GetTenant(*tenant)
	if e != nil {
		return nil, e
	}

	c.Tenant = *t

	return &c, nil
}

func (c *Client) doRequest(req *http.Request, authToken *string) ([]byte, error) {
	token := c.Token

	if authToken != nil {
		token = *authToken
	}

	req.Header.Set("Authorization", token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
