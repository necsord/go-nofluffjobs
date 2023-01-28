package nofluffjobs

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	contentType = "application/json"
)

type Client struct {
	baseUrl    *url.URL
	httpClient *http.Client
	logger     log.Logger
}

func NewClient(baseUri string, httpClient *http.Client, logger *log.Logger) (*Client, error) {
	parsedUrl, err := url.Parse(baseUri)
	if err != nil {
		return nil, fmt.Errorf("client creation failed: %w", err)
	}

	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 5 * time.Second,
		}
	}
	if logger == nil {
		logger = log.Default()
	}

	return &Client{
		baseUrl:    parsedUrl,
		httpClient: httpClient,
	}, nil
}

func (c *Client) SearchPosting(ctx context.Context, reqQParams SearchPostingQuery, reqBody SearchPostingRequest) (*SearchPostingResponse, error) {
	qParams, err := query.Values(reqQParams)
	if err != nil {
		return nil, fmt.Errorf("could not process query params: %w", err)
	}

	endpoint := fmt.Sprintf("%s/search/posting?%s", c.baseUrl.String(), qParams.Encode())
	req, err := c.prepareRequest(ctx, endpoint, http.MethodPost, reqBody)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer c.closeBody(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, ErrorResponse{Response: res}
	}

	postings := &SearchPostingResponse{}
	return postings, parseResponse(res, postings)
}

func (c *Client) prepareRequest(ctx context.Context, endpoint, httpMethod string, body any) (*http.Request, error) {
	var bufReqBody io.Reader
	if body != nil {
		reqBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("could not encode request body: %w", err)
		}
		bufReqBody = bytes.NewBuffer(reqBody)
	}

	req, err := http.NewRequestWithContext(ctx, httpMethod, endpoint, bufReqBody)
	if err != nil {
		return nil, fmt.Errorf("request creation failed: %w", err)
	}
	req.Header.Set("Content-Type", contentType)

	return req, nil
}

func (c *Client) closeBody(body io.ReadCloser) {
	err := body.Close()
	if err != nil {
		c.logger.Printf("error occurred when closing response body: %w\n", err)
	}
}

func parseResponse(res *http.Response, v any) error {
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("could not read response body (status code: %d): %w", res.StatusCode, err)
	}

	err = json.Unmarshal(resBody, v)
	if err != nil {
		return &ErrorResponse{Response: res, Err: fmt.Errorf("could not decode response: %w", err)}
	}

	return nil
}
