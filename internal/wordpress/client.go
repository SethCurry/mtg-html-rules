package wordpress

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// AuthProvider defines an interface for adding authentication to HTTP requests.
type AuthProvider interface {
	AddAuth(req *http.Request)
}

// BasicAuthProvider is an AuthProvider that uses HTTP basic authentication.
type BasicAuthProvider struct {
	Username string
	Password string
}

// AddAuth adds HTTP basic authentication to the request.
func (p *BasicAuthProvider) AddAuth(req *http.Request) {
	req.SetBasicAuth(p.Username, p.Password)
}

type Client struct {
	BaseURL string
	Auth    AuthProvider
}

func NewClient(baseURL string, auth AuthProvider) *Client {
	return &Client{
		BaseURL: baseURL,
		Auth:    auth,
	}
}

type UpdatePostRequest struct {
	Title   *string `json:"title,omitempty"`
	Content *string `json:"content,omitempty"`
}

func (c *Client) UpdatePost(postID int, request UpdatePostRequest) error {
	body, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/wp/v2/posts/%d", c.BaseURL, postID), bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	if c.Auth != nil {
		c.Auth.AddAuth(req)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update post: got status code %d", resp.StatusCode)
	}

	return nil
}
