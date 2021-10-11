package ciscobcs

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

// Client is the main cisco bcs client for interacting with the library.  It can be created using NewClient
type Client struct {
	// BaseURL for BCS API.  Set to https://demo.api.csco-bcs.com/v2 using `ciscobcs.New()`, or set directly.
	BaseURL string

	//HTTP Client to use for making requests, allowing the user to supply their own if required.
	HTTPClient *http.Client

	//API Key for Cisco BCS.
	APIKey string

	// Services for accessing the various endpoints
	BulkService                      *BulkService
	ConfigurationBestPracticeService *ConfigurationBestPracticeService
	CrashPreventionService           *CrashPreventionService
	FeedbackService                  *FeedbackService
	InventoryService                 *InventoryService
	ProductAlertService              *ProductAlertService
	RiskMitigationService            *RiskMitigationService
	SoftwareTrackService             *SoftwareTrackService
	ContractService                  *ContractService
	CollectorsService                *CollectorsService
	CountService                     *CountService
	SyslogService                    *SyslogService

	lim *rate.Limiter
}

// BulkService
type BulkService struct {
	client *Client
}

// ConfigurationBestPracticeService
type ConfigurationBestPracticeService struct {
	client *Client
}

// CrashPreventionService
type CrashPreventionService struct {
	client *Client
}

// FeedbackService
type FeedbackService struct {
	client *Client
}

// InventoryService
type InventoryService struct {
	client *Client
}

// ProductAlertService
type ProductAlertService struct {
	client *Client
}

// RiskMitigationService
type RiskMitigationService struct {
	client *Client
}

// SoftwareTrackService
type SoftwareTrackService struct {
	client *Client
}

// ContractService
type ContractService struct {
	client *Client
}

// CollectorsService
type CollectorsService struct {
	client *Client
}

// CountService
type CountService struct {
	client *Client
}

// SyslogService
type SyslogService struct {
	client *Client
}

// NewClient is a helper function that returns an new cisco bcs client given an API Key.
// Optionally you can provide your own http client or use nil to use the default.  This is done to
// ensure you're aware of the decision you're making to not provide your own http client.
func NewClient(apikey string, client *http.Client) (*Client, error) {
	if apikey == "" {
		return nil, ErrMissingAPIKey
	}
	if client == nil {
		client = &http.Client{
			Timeout: 10 * time.Second,
		}
	}
	rl := rate.NewLimiter(150, 1) // this is not documented, so we'll limit to 150/s
	c := &Client{
		BaseURL:    "https://demo.api.csco-bcs.com/v2",
		HTTPClient: client,
		APIKey:     apikey,
		lim:        rl,
	}
	c.BulkService = &BulkService{client: c}
	c.ConfigurationBestPracticeService = &ConfigurationBestPracticeService{client: c}
	c.CrashPreventionService = &CrashPreventionService{client: c}
	c.FeedbackService = &FeedbackService{client: c}
	c.InventoryService = &InventoryService{client: c}
	c.ProductAlertService = &ProductAlertService{client: c}
	c.RiskMitigationService = &RiskMitigationService{client: c}
	c.SoftwareTrackService = &SoftwareTrackService{client: c}
	c.ContractService = &ContractService{client: c}
	c.CollectorsService = &CollectorsService{client: c}
	c.CountService = &CountService{client: c}
	c.SyslogService = &SyslogService{client: c}

	return c, nil
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// Float64 is a helper routine that allocates a new Float64 value
// to store v and returns a pointer to it.
func Float64(v float64) *float64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

func (c *Client) makeRequestToWriter(ctx context.Context, req *http.Request, w io.Writer) error {
	req.Header.Add("x-api-key", c.APIKey)
	rc := req.WithContext(ctx)
	res, err := c.HTTPClient.Do(rc)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if _, err := io.Copy(w, res.Body); err != nil {
		return err
	}
	return nil
}

// makeRequest provides a single function to add common items to the request.
func (c *Client) makeRequest(ctx context.Context, req *http.Request, v interface{}) error {
	req.Header.Add("x-api-key", c.APIKey)
	if !c.lim.Allow() {
		c.lim.Wait(ctx)
	}
	rc := req.WithContext(ctx)
	res, err := c.HTTPClient.Do(rc)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var ciscobcsErr error
		switch res.StatusCode {
		case 400:
			ciscobcsErr = ErrBadRequest
		case 401:
			ciscobcsErr = ErrUnauthorized
		case 403:
			ciscobcsErr = ErrForbidden
		case 500:
			ciscobcsErr = ErrInternalError
		default:
			ciscobcsErr = ErrUnknown
		}
		return ciscobcsErr
	}
	if res.StatusCode == http.StatusCreated {
		return nil
	}
	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}
	return nil
}
