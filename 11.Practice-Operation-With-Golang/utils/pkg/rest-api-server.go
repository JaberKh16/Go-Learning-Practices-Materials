package restapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	BaseURLV1 = "https://api.facest.io/v1"
)

type Client struct {
	BaseURL string 
	apiKey string
	HTTPClient *http.Client
}

type successResponse struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
	Message string `json:"message"`
}

type errorResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

type FacesResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data FacesList `json:"data"`
}

type FacesList struct {
	Count int `json:"count"`
	PagesCount int `json:"pages_count"`
	Faces []Face `json:"faces"`
}

type Face struct {
	FaceToken string `json:"face_token"`
	FaceID string `json:"face_id"`
	FaceImages [] FaceImage `json:"face_images"`
	CreatedAt time.Timer `json:"created_at"`
}

type FaceImage struct {
	ImageToken string `json:"image_token"`
	ImageURL string `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
}

type FacesListOptions struct {
	Limit int `json:"limit"`
	Page int `json:"page"`
}

func NewClient(apiKey string) *Client {
	return &Client{
		BaseURL: BaseURLV1,
		apiKey: appKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) GetFaces(ctx context.Context, options *FacesListOptions) (*FacesList, error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	// setup a request
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/faces?limit=%d&page=%d", c.BaseURL, limit, page), nil)
	if err != nil {
		return  nil, err
	}


	// setup a request context
	request = request.WithContext(ctx)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))


	// send the request
	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return  nil, err
	}

	defer response.Body.Close()

	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusBadRequest {
		var errResponse errorResponse
		if err := json.NewDecoder(response.Body).Decode(&errResponse); err != nil {
			return  nil, fmt.Errorf("unknown error, status code: %d", response.StatusCode)
		}
		return  nil, errors.New(errResponse.Message)
	}

	var fullResponse FacesResponse
	if err := json.NewDecoder(request.Body).Decode(&fullResponse); err != nil {
		return nil, err
	}
	return  &fullResponse.Data, nil
}

func (c *Client) sendRequest(ctx context.Context, w http.ResponseWriter, r *http.Request, v interface{}) error {
	// setup a request context
	request = request.WithContext(ctx)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))


	// send the request
	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return  nil, err
	}

	defer response.Body.Close()

	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusBadRequest {
		var errResponse errorResponse
		if err := json.NewDecoder(response.Body).Decode(&errResponse); err != nil {
			return  nil, fmt.Errorf("unknown error, status code: %d", response.StatusCode)
		}
		return  nil, errors.New(errResponse.Message)
	}

	var fullResponse FacesResponse
	if err := json.NewDecoder(request.Body).Decode(&fullResponse); err != nil {
		return nil, err
	}
	return  &fullResponse.Data, nil
}